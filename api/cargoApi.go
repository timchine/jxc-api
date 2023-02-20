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

// @Summary 新增货品种类
// @Description 新增货品种类，新增种类时同时新增种类相关规格属性
// @Accept  json
// @Produce  json
// @Param 货品种类 body dto.ReqAddCargoKind true "货品种类和属性"
// @Reponse 200 {object} Response "status 200 表示成功 否则提示msg内容"
// @Router /cargo_kind [post]
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
