package api

import (
	"crypto/md5"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/timchine/jxc/api/dto"
	"github.com/timchine/jxc/model"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type ICargoApi interface {
	AddCargoKind() gin.HandlerFunc
	GetCargoKind() gin.HandlerFunc
	UpdateCargoKind() gin.HandlerFunc
	DeleteCargoKind() gin.HandlerFunc
	SearchCargoKind() gin.HandlerFunc
	AddCargo() gin.HandlerFunc
	DeleteCargo() gin.HandlerFunc
	UpdateCargo() gin.HandlerFunc
	GetCargo() gin.HandlerFunc
	SearchCargo() gin.HandlerFunc
	UploadImage() gin.HandlerFunc
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
// @Param			货品种类	body		dto.CargoKindWithAttrs	true	"货品种类和属性"
// @Response		200		{object}	Response				"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind [post]
func (c *cargoApi) AddCargoKind() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req dto.CargoKindWithAttrs
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
// @Param			ck_id	path		int						true	"货品种类ID"
// @Response		200		{object}	dto.CargoKindWithAttrs	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind/{ck_id} [get]
func (c *cargoApi) GetCargoKind() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req  string
			id   int
			err  error
			res  Response
			data dto.CargoKindWithAttrs
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
// @Param			货品种类	body		dto.CargoKindWithAttrs	true	"货品种类和属性"
// @Response		200		{object}	Response				"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind [put]
func (c *cargoApi) UpdateCargoKind() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req dto.CargoKindWithAttrs
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
			if v.CaID == 0 {
				v.CkID = req.CargoKind.CkID
				err = tx.Create(&v).Error
				if err != nil {
					Log().Error(err.Error())
					tx.Rollback()
					res.Error(ctx, 500, "修改失败")
					return
				}
				break
			}
			err = tx.Where("ca_id=?", v.CaID).Updates(&v).Error
			if err != nil {
				Log().Error(err.Error())
				tx.Rollback()
				res.Error(ctx, 500, "修改失败")
				return
			}
		}

		tx.Commit()
		res.Success(ctx)
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
// @Param			search	query		string				false	"货品code 或者 货品名称"
// @Param			page	query		string				false	"页数"
// @Param			size	query		string				false	"每页条数"
// @Response		200		{object}	[]model.CargoKind	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo_kind/_search [get]
func (c *cargoApi) SearchCargoKind() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req  dto.ReqSearchCargoKind
			err  error
			data []model.CargoKind
			res  Response
		)

		err = ctx.BindQuery(&req)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "参数错误")
			return
		}
		if req.Page == 0 {
			req.Page = 1
		}
		if req.Size == 0 {
			req.Size = 10
		}
		tx := c.Model(&model.CargoKind{}).Where("status != 8")
		if req.Search != "" {
			search := fmt.Sprintf("%%%s%%", req.Search)
			tx = tx.Where("ck_code like ? or ck_name like ?", search, search)
		}
		err = tx.Count(&req.Total).Limit(req.Size).Offset((req.Page - 1) * req.Size).Find(&data).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "搜索失败")
			return
		}
		req.Data = data
		req.TotalPage = math.Ceil(float64(req.Total) / float64(req.Size))
		res.Success(ctx, req)
	}
}

// @Summary		新增货品
// @Description	此接口用于新增原材料或制品， 新增时包括 计量和属性值
// @Accept			json
// @Produce		json
// @Param			ReqAddCargo	body		dto.ReqAddCargo	true	"货品和属性 和计量单位"
// @Response		200			{object}	Response		"status 200 表示成功 否则提示msg内容"
// @Router			/cargo [post]
func (c *cargoApi) AddCargo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req dto.ReqAddCargo
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
		err = tx.Create(&req.Cargo).Error
		if err != nil {
			tx.Rollback()
			Log().Error(err.Error())
			res.Error(ctx, 500, "新增失败")
			return
		}
		// 新增 计量单位
		for k := range req.Measures {
			req.Measures[k].CargoID = req.Cargo.CargoID
		}
		err = tx.Create(&req.Measures).Error
		if err != nil {
			tx.Rollback()
			Log().Error(err.Error())
			res.Error(ctx, 500, "新增失败")
			return
		}
		//新增属性值
		for k := range req.CargoAttrValues {
			req.CargoAttrValues[k].CargoID = req.Cargo.CargoID
		}
		err = tx.Create(&req.CargoAttrValues).Error
		if err != nil {
			tx.Rollback()
			Log().Error(err.Error())
			res.Error(ctx, 500, "新增失败")
			return
		}
		// todo 修改图片状态
		tx.Commit()
		res.Success(ctx)
	}
}

// @Summary		删除货品
// @Description	删除货品
// @Param			cargo_id	path		int			true	"货品ID"
// @Response		200			{object}	Response	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo/{cargo_id} [delete]
func (c *cargoApi) DeleteCargo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			cargoId int
			err     error
			res     Response
		)
		id := ctx.Param("cargo_id")
		cargoId, err = strconv.Atoi(id)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "参数错误")
			return
		}
		tx := c.Begin()
		err = tx.Model(&model.Cargo{}).Where("cargo_id=?", cargoId).Update("status", 8).Error
		if err != nil {
			tx.Rollback()
			Log().Error(err.Error())
			res.Error(ctx, 500, "删除失败")
			return
		}
		// 删除 计量单位
		err = tx.Model(&model.Measure{}).Where("cargo_id=?", cargoId).Update("status", 8).Error
		if err != nil {
			tx.Rollback()
			Log().Error(err.Error())
			res.Error(ctx, 500, "删除失败")
			return
		}
		//删除属性值
		err = tx.Model(&model.CargoAttrValue{}).Where("cargo_id=?", cargoId).Update("status", 8).Error
		if err != nil {
			tx.Rollback()
			Log().Error(err.Error())
			res.Error(ctx, 500, "删除失败")
			return
		}
		// todo 修改图片状态
		tx.Commit()
		res.Success(ctx)
	}
}

// @Summary		修改货品
// @Description	修改货品
// @Accept			json
// @Produce		json
// @Param			ReqAddCargo	body		dto.ReqAddCargo	true	"货品和属性 和计量单位"
// @Response		200			{object}	Response		"status 200 表示成功 否则提示msg内容"
// @Router			/cargo [put]
func (c *cargoApi) UpdateCargo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req dto.ReqAddCargo
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
		// todo 如果图片更改 标记原有图片为未使用状态
		err = tx.Where("cargo_id=?", req.Cargo.CargoID).Updates(&req.Cargo).Error
		if err != nil {
			tx.Rollback()
			Log().Error(err.Error())
			res.Error(ctx, 500, "修改失败")
			return
		}
		// 修改 计量单位
		for _, v := range req.Measures {
			if v.MeasureID == 0 {
				v.CargoID = req.Cargo.CargoID
				err = c.Create(&v).Error
				if err != nil {
					tx.Rollback()
					Log().Error(err.Error())
					res.Error(ctx, 500, "修改失败")
					return
				}
				break
			}
			err = c.Where("measure_id=?", v.MeasureID).Updates(&v).Error
			if err != nil {
				tx.Rollback()
				Log().Error(err.Error())
				res.Error(ctx, 500, "修改失败")
				return
			}
		}
		//修改属性值
		for _, v := range req.CargoAttrValues {
			if v.CavID == 0 {
				v.CargoID = req.Cargo.CargoID
				err = c.Create(&v).Error
				if err != nil {
					tx.Rollback()
					Log().Error(err.Error())
					res.Error(ctx, 500, "修改失败")
					return
				}
				break
			}
			err = c.Where("cav_id=?", v.CavID).Updates(&v).Error
			if err != nil {
				tx.Rollback()
				Log().Error(err.Error())
				res.Error(ctx, 500, "修改失败")
				return
			}
		}
		// todo 修改图片状态
		tx.Commit()
		res.Success(ctx)
	}
}

// @Summary		获取货品详情
// @Description	获取货品详情
// @Param			cargo_id	path		int				true	"货品ID"
// @Response		200			{object}	dto.ReqAddCargo	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo/{cargo_id} [get]
func (c *cargoApi) GetCargo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			cargoId int
			err     error
			res     Response
			data    dto.ReqAddCargo
		)
		id := ctx.Param("cargo_id")
		cargoId, err = strconv.Atoi(id)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "参数错误")
			return
		}
		err = c.Model(&model.Cargo{}).Where("cargo_id=?", cargoId).First(&data.Cargo).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "查询失败")
			return
		}
		// 删除 计量单位
		err = c.Model(&model.Measure{}).Where("cargo_id=?", cargoId).Find(&data.Measures).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "查询失败")
			return
		}
		//删除属性值
		err = c.Model(&model.CargoAttrValue{}).Where("cargo_id=?", cargoId).Find(&data.CargoAttrValues).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "查询失败")
			return
		}
		res.Success(ctx, data)
	}
}

// @Summary		搜索货品
// @Description	通过货品code 或者 货品名称搜索 货品 (可以是原材料 和制品)
// @Param			search	query		string				false	"货品code 或者 货品名称"
// @Param			type	query		string				true	"类型 1:原材料 2:半成品 3:成品"
// @Param			page	query		string				false	"页数"
// @Param			size	query		string				false	"每页条数"
// @Response		200		{object}	[]model.CargoKind	"status 200 表示成功 否则提示msg内容"
// @Router			/cargo/_search [get]
func (c *cargoApi) SearchCargo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req  dto.ReqSearchCargo
			err  error
			res  Response
			data []model.Cargo
		)
		err = ctx.BindQuery(&req)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "参数错误")
			return
		}
		if req.Page == 0 {
			req.Page = 1
		}
		if req.Size == 0 {
			req.Size = 10
		}
		tx := c.Table("cargo as c").Select("c.cargo_id, c.ck_id, c.cargo_name, c.cargo_code, c.status").
			Joins("left join cargo_kind as ck on ck.ck_id = c.ck_id").Where("c.status != 8 and ck.type = ?", req.Type)
		if req.Search != "" {
			search := fmt.Sprintf("%%%s%%", req.Search)
			tx = tx.Where("c.cargo_code like ? or c.cargo_name like ?", search, search)
		}
		err = tx.Count(&req.Total).Limit(req.Size).Offset((req.Page - 1) * req.Size).Find(&data).Error
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "搜索失败")
			return
		}
		req.Data = data
		req.TotalPage = math.Ceil(float64(req.Total) / float64(req.Size))
		res.Success(ctx, req)
	}
}

// @Summary		上传图片
// @Description	上传图片
// @Param			search	formData		file				true	"文件"
// @Response		200		{object}	Response	"status 200 表示成功 否则提示msg内容"
// @Router			/image [post]
func (c *cargoApi) UploadImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			res   Response
			image model.Image
		)
		file, err := ctx.FormFile("file")
		if err != nil || file.Size < 200 {
			Log().Error(err.Error())
			res.Error(ctx, 500, "文件为空")
			return
		}
		fr, err := file.Open()
		var (
			b = make([]byte, 200)
		)
		fr.Read(b[:100])
		fr.ReadAt(b[100:], file.Size-100)
		fr.Close()
		imageHash := fmt.Sprintf("%x", md5.Sum(b))
		// 查询图片
		c.Where("image_hash=?", imageHash).First(&image)
		if image.ImageID != 0 {
			res.Success(ctx, image)
			return
		}
		fr, err = file.Open()
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "文件为空")
			return
		}
		img, err := imaging.Decode(fr)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "文件为空")
			return
		}
		img = imaging.Fill(img, 100, 100, imaging.Center, imaging.Lanczos)
		fr.Close()
		key := uuid.New().String()
		err = imaging.Save(img, "static/upload/"+key)
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "保存文件失败")
			return
		}
		image.ImageHash = imageHash
		image.ThumbnailName = key
		err = ctx.SaveUploadedFile(file, "static/upload/"+uuid.New().String())
		if err != nil {
			Log().Error(err.Error())
			res.Error(ctx, 500, "保存文件失败")
			return
		}
		image.ImageName = key
		err = c.Create(&image).Error
		if err != nil {
			res.Error(ctx, 500, "保存文件失败")
			return
		}
		res.Success(ctx, image)
	}
}
