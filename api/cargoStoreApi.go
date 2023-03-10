package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICargoStoreApi interface {
	SearchCargoStore() gin.HandlerFunc
	SetCargoStoreWarning() gin.HandlerFunc
}

type cargoStoreApi struct {
	*gorm.DB
}

func NewCargoStoreApi(db *gorm.DB) ICargoStoreApi {
	return &cargoStoreApi{DB: db}
}

// @Summary		查询库存
// @Description	通过种类 或 种类类型 或货品名称 编码查询库存
// @Param			ReqSearchCargoStore	body		dto.ReqSearchCargoStore		true	"查询条件"
// @Response		200		{object}	[]model.CargoStore	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_store/_search [post]
func (c *cargoStoreApi) SearchCargoStore() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// @Summary		设置库存告警
// @Description	设置库存告警
// @Param			cargoStore	body		model.CargoStore		true	"货品ID"
// @Response		200		{object}	Response	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_store/warning [put]
func (c *cargoStoreApi) SetCargoStoreWarning() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
