package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IChannelApi interface {
	AddChannel() gin.HandlerFunc
	EditChannel() gin.HandlerFunc
}

type channelApi struct {
	*gorm.DB
}

func NewChannelApi(db *gorm.DB) IChannelApi {
	return &channelApi{
		db,
	}
}
