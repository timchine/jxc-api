package log

import (
	"fmt"
	logger2 "gorm.io/gorm/logger"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ModeDev  = "dev"
	ModeProd = "prod"
)

func NewLogger(mode string, level zapcore.Level) error {
	var (
		logConfig zap.Config
		err       error
		//file      *os.File
	)
	switch mode {
	case "", ModeDev:
		logConfig = zap.NewDevelopmentConfig()
		logConfig.Level = zap.NewAtomicLevelAt(level)
		logConfig.DisableCaller = true
		logConfig.Sampling = &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
			//Hook: func(message zapcore.Entry, deci zapcore.SamplingDecision) {
			//	if file == nil {
			//		file, err = os.OpenFile(time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0766)
			//		if err != nil {
			//			fmt.Println(err)
			//			return
			//		}
			//	}
			//	if file.Name() != time.Now().Format("2006-01-02")+".log" {
			//		file.Close()
			//		file, err = os.OpenFile(time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0766)
			//		if err != nil {
			//			fmt.Println(err)
			//			return
			//		}
			//	}
			//	file.WriteString(message.Message + "\n")
			//},
		}
		logConfig.EncoderConfig.EncodeTime = timeEncoder
		logConfig.EncoderConfig.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {}
		// logConfig.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	case ModeProd:
		logConfig = zap.NewProductionConfig()
		logConfig.Level = zap.NewAtomicLevelAt(level)
	default:
		panic("unknown run mode it mast dev or prod")
	}
	logger, err = logConfig.Build()
	return err
}

var (
	logger *zap.Logger
)

func Logger() *zap.Logger {
	return logger
}

func timeEncoder(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	s := fmt.Sprintf("\x1b[0;33m%s\x1b[0m", time.Format("[2006-01-02 15:04:05]"))
	encoder.AppendString(s)
}

type ormLog struct {
	*zap.Logger
}

func NewOrmLog() logger2.Writer {
	return &ormLog{
		logger,
	}
}

func (o *ormLog) Printf(s string, i ...interface{}) {
	fmt.Println(123)
	o.Logger.Info(fmt.Sprintf(s, i...))
}
