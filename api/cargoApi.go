package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/api/dto"
	"github.com/timchine/jxc/model"
	"gorm.io/gorm"
	"strconv"
)

type ICargoApi interface {
	AddCargoKind() gin.HandlerFunc
	GetCargoKind() gin.HandlerFunc
	UpdateCargoKind() gin.HandlerFunc
	DeleteCargoKind() gin.HandlerFunc
	SearchCargoKind() gin.HandlerFunc
}

type cargoApi struct {
	*gorm.DB
}

func NewCargoApi(db *gorm.DB) ICargoApi {
	return &cargoApi{DB: db}
}

// @Summary		新增货品种类
// @Description	新增货品种类，新增种类时同时新增种类相关规格属性
// @Accept			json
// @Produce		json
// @Param			货品种类	body		dto.ReqAddCargoKind	true	"货品种类和属性"
// @Response		200		{object}	Response			"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind [post]
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
		tx := c.Begin()
		err = tx.Create(&req.CargoKind).Error
		if err != nil {
			Log().Error(err.Error())
			tx.Rollback()
			res.Error(ctx, 500, "新增失败")
			return
		}
		for k := range req.CargoAttrs {
			req.CargoAttrs[k].CkID = req.CargoKind.CkID
		}
		err = tx.Create(&req.CargoAttrs).Error
		if err != nil {
			Log().Error(err.Error())
			tx.Rollback()
			res.Error(ctx, 500, "新增失败")
			return
		}
		tx.Commit()
		res.Success(ctx, req)
	}
}

// @Summary		获取货品详情
// @Description	获取货品详情，获取货品详情和相关属性
// @Param			ck_id	path		int					true	"货品种类ID"
// @Response		200		{object}	dto.ReqAddCargoKind	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind/{ck_id} [get]
func (c *cargoApi) GetCargoKind() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req  string
			id   int
			err  error
			res  Response
			data dto.ReqAddCargoKind
		)
		req = ctx.Param("ck_id")
		id, err = strconv.Atoi(req)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "参数错误")
			return
		}
		err = c.Where("ck_id=?", id).First(&data.CargoKind).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "查询失败")
			return
		}
		err = c.Where("ck_id=?", id).Find(&data.CargoAttrs).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "查询失败")
			return
		}
		res.Success(ctx, data)
	}
}

// @Summary		修改货品种类
// @Description	修改货品种类
// @Accept			json
// @Produce		json
// @Param			货品种类	body		dto.ReqAddCargoKind	true	"货品种类和属性"
// @Response		200		{object}	Response			"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind [put]
func (c *cargoApi) UpdateCargoKind() gin.HandlerFunc {
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
		tx := c.Begin()
		err = tx.Where("ck_id=?", req.CargoKind.CkID).Updates(req.CargoKind).Error
		if err != nil {
			Log().Error(err.Error())
			tx.Rollback()
			res.Error(ctx, 500, "修改失败")
			return
		}
		for _, v := range req.CargoAttrs {
			err = tx.Where("ca_id=?", v.CaID).Updates(&v).Error
			if err != nil {
				Log().Error(err.Error())
				tx.Rollback()
				res.Error(ctx, 500, "修改失败")
				return
			}
		}

		tx.Commit()
		res.Success(ctx, req)
	}
}

// @Summary		删除货品
// @Description	删除货品
// @Param			ck_id	path		int			true	"货品种类ID"
// @Response		200		{object}	Response	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind/{ck_id} [delete]
func (c *cargoApi) DeleteCargoKind() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req string
			id  int
			err error
			res Response
		)
		req = ctx.Param("ck_id")
		id, err = strconv.Atoi(req)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "参数错误")
			return
		}
		err = c.Model(&model.CargoKind{}).Where("ck_id=?", id).Update("status", 8).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "删除失败")
			return
		}
		err = c.Model(&model.CargoAttr{}).Where("ck_id=?", id).Update("status", 8).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "删除失败")
			return
		}
		res.Success(ctx)
	}
}

// @Summary		搜索货品种类
// @Description	通过货品code 或者 货品名称搜索 货品种类
// @Param			search	query		string					false	"货品code 或者 货品名称"
// @Response		200		{object}	[]dto.ReqAddCargoKind	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind/_search [get]
func (c *cargoApi) SearchCargoKind() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			//	req  string
			//	err  error
			data []model.CargoKind
		)
		//req = ctx.Query("search")
		c.Find(&data)
		fmt.Println(data)
	}
}
