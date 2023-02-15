package router

import (
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/api"
	"gorm.io/gorm"
)

func customerRouter(r *gin.RouterGroup, db *gorm.DB) {
	a := api.NewChannelApi(db)
	r.POST("customer", a.AddChannel())
	r.PUT("customer", a.EditChannel())
}
