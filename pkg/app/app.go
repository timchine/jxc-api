package app

import (
	"context"
	"github.com/timchine/jxc/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
)

type app struct {
	mode           string
	name           string
	daemon         map[string]DaemonFunc
	stage          map[string]StageFunc
	orderFuncNames []string
	cleanFunc      []CleanFunc
	StageChan      chan bool
}

type CleanFunc func() error
type StageFunc func(ctx context.Context) (CleanFunc, error)
type DaemonFunc func(ctx context.Context) error

func NewApp(name string, mode string, logLevel zapcore.Level) (*app, error) {
	s := &app{
		mode:      mode,
		name:      name,
		daemon:    make(map[string]DaemonFunc),
		stage:     make(map[string]StageFunc),
		StageChan: make(chan bool, 1),
	}
	err := log.NewLogger(mode, logLevel)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (a *app) Run() error {
	if err := a.runStage(); err != nil {
		log.Logger().Error(err.Error())
		return err
	}
	if len(a.daemon) != 0 {
		if err := a.runDaemon(); err != nil {
			log.Logger().Error(err.Error())
			return err
		}
	}
	<-a.StageChan
	return nil
}

func (a *app) Close() error {
	for _, f := range a.cleanFunc {
		err := f()
		if err != nil {
			log.Logger().Error(err.Error())
			return err
		}
	}
	return nil
}

func (a *app) AddStageFunc(stageFunc StageFunc) *app {
	name := getFuncName(stageFunc)
	a.stage[name] = stageFunc
	log.Logger().Info("set stage func ...", zap.String("name", name))
	return a
}

func (a *app) AddDaemonFunc(daemonFunc DaemonFunc) *app {
	name := getFuncName(daemonFunc)
	a.daemon[name] = daemonFunc
	log.Logger().Info("set daemon func", zap.String("name", name))
	return a
}

func (a *app) SignalDaemon(signalChan chan os.Signal) DaemonFunc {
	return func(ctx context.Context) error {
		<-signalChan
		if len(a.StageChan) == 1 {
			<-a.StageChan
		}
		a.StageChan <- true
		return nil
	}
}

func (a *app) runStage() error {
	var (
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer cancel()
	for name, fc := range a.stage {
		go func(fc StageFunc) {
			defer func() {
				if r := recover(); r != nil {
					log.Logger().Panic("init stage panic:", zap.Stack("stage"))
				}
			}()
			clean, err := fc(ctx)
			if err != nil {
				log.Logger().Error("setup stage error"+err.Error(), zap.String("name", name))
				cancel()
			} else {
				delete(a.stage, name)
				log.Logger().Info("setup stage success", zap.String("name", name))
			}

			if clean != nil {
				a.cleanFunc = append(a.cleanFunc, clean)
			}
			if 0 == len(a.stage) {
				a.StageChan <- true
			}
		}(fc)
	}

	return nil
}

func (a *app) runDaemon() error {
	var (
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer cancel()
	for name, fc := range a.daemon {
		go func(fc DaemonFunc) {
			defer func() {
				if r := recover(); r != nil {
					log.Logger().Panic("init daemon panic:", zap.Stack("stage"))
				}
			}()
			err := fc(ctx)
			if err != nil {
				log.Logger().Error("setup daemon error"+err.Error(), zap.String("name", name))
				cancel()
			} else {
				log.Logger().Info("setup daemon success", zap.String("name", name))
			}
		}(fc)
	}
	var (
		signalChan = make(chan os.Signal, 1)
	)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	_ = a.SignalDaemon(signalChan)(ctx)
	return nil
}

func getFuncName(f interface{}) string {
	fv := reflect.ValueOf(f)

	if fv.Kind() != reflect.Func {
		return ""
	}
	return runtime.FuncForPC(fv.Pointer()).Name()
}
