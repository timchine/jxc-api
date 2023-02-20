package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/timchine/jxc/docs"
	"github.com/timchine/jxc/pkg/app"
	"github.com/timchine/jxc/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func routers(r *gin.RouterGroup, db *gorm.DB) {
	docs.SwaggerInfo.BasePath = "/api/jxc"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	customerRouter(r, db)
	cargoRouter(r, db)
}

func Run(db *gorm.DB) app.DaemonFunc {
	return func(ctx context.Context) error {
		router := gin.New()
		router.Use(LogHandler, Recovery)
		routers(router.Group("/api/jxc/"), db)
		return router.Run(viper.GetString("server.addr"))
	}
}

func LogHandler(ctx *gin.Context) {
	var (
		start = time.Now()
		path  = ctx.Request.URL.Path
		query = ctx.Request.URL.RawQuery
	)
	ctx.Next()
	cost := time.Since(start)
	log.Logger().Info(path,
		zap.String("method", ctx.Request.Method),
		zap.Int("status", ctx.Writer.Status()),
		zap.String("query", query),
		zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
		zap.Duration("cost", cost),
	)
}

func Recovery(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Logger().Error("[Recovery from panic]",
				zap.Any("error", err),
			)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	ctx.Next()
}
