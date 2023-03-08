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
	storeRouter(r, db)
}

func Run(db *gorm.DB) app.DaemonFunc {
	return func(ctx context.Context) error {
		router := gin.New()
		router.Use(LogHandler, Recovery, Cors)
		router.Static("/static/upload", "./static/upload")
		router.Use(LogHandler, Recovery)
		routers(router.Group("/api/jxc/"), db)
		return router.Run(viper.GetString("server.addr"))
	}
}

func Cors(ctx *gin.Context) {
	method := ctx.Request.Method
	origin := ctx.Request.Header.Get("Origin")
	if origin != "" {
		ctx.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
	}
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	ctx.Next()
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
