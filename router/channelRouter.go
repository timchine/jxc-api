package router

import (
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/api"
	"gorm.io/gorm"
)

func channelRouter(r *gin.RouterGroup, db *gorm.DB) {
	a := api.NewChannelApi(db)
	r.POST("channel", a.AddChannel())
	r.PUT("channel", a.EditChannel())
}
