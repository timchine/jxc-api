package api

import (
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/api/dto"
	"gorm.io/gorm"
)

type ICargoApi interface {
	AddCargoKind() gin.HandlerFunc
}

type cargoApi struct {
	*gorm.DB
}

func NewCargoApi(db *gorm.DB) ICargoApi {
	return &cargoApi{DB: db}
}

func (c *cargoApi) AddCargoKind() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req dto.ReqAddCargoKind
			err error
			res Response
		)
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "参数错误")
			return
		}

	}
}
