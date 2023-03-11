package model

// 计量
type Measure struct {
	Model
	MeasureID int     `json:"measure_id"`
	CargoID   int     `json:"cargo_id"`                //关联货品
	IsBase    bool    `json:"is_base"`                 //是否为基础计量单位
	Unit      string  `json:"unit"`                    //单位
	Calc      float64 `json:"calc"`                    //换算比例
	Status    int     `json:"status" gorm:"default:1"` //状态 1 正常 8 删除
}

// 货物 分类
type CargoKind struct {
	Model
	CkID   int    `json:"ck_id" gorm:"primaryKey;autoIncrement"`
	CkCode string `json:"ck_code"`                 //货品编码
	CkName string `json:"ck_name"`                 //货品名称
	Intro  string `json:"intro"`                   //货品简介
	Type   int    `json:"type"`                    //1物料 2制品
	Status int    `json:"status" gorm:"default:1"` //状态 1 正常 8 删除
}

// 货物属性
type CargoAttr struct {
	Model
	CaID      int    `json:"ca_id" gorm:"primaryKey;autoIncrement"` //属性ID
	CkID      int    `json:"ck_id"`                                 //关联货品种类
	AttrName  string `json:"attr_name"`                             //属性名称
	Type      int    `json:"type"`                                  //1 选择 2 文本
	AttrValue string `json:"attr_value"`                            //属性值 ｜ 符号分割
	Status    int    `json:"status" gorm:"default:1"`               //状态 1 正常 8 删除
}

// 货物
type Cargo struct {
	Model
	CargoID       int    `json:"cargo_id" gorm:"primaryKey;autoIncrement"` //货品ID
	CkID          int    `json:"ck_id"`                                    //货品种类ID
	CargoName     string `json:"cargo_name"`                               //货品名称
	CargoCode     string `json:"cargo_code"`                               //货品编码
	ThumbnailName string `json:"thumbnail_name"`                           //缩略图
	ImageName     string `json:"image_name"`                               //大图
	Status        int    `json:"status" gorm:"default:1"`                  //状态 1 正常 8 删除
}

type CargoAttrValue struct {
	Model
	CavID     int    `json:"cav_id"`                  //ID
	CaID      int    `json:"ca_id"`                   //货物属性ID
	CargoID   int    `json:"cargo_id"`                //货物ID
	AttrName  string `json:"attr_name"`               //属性名称
	AttrValue string `json:"value"`                   //属性值
	Status    int    `json:"status" gorm:"default:1"` //状态 1 正常 8 删除
}

type CargoProcess struct {
	Model
	ProcessID int     `json:"process_id" gorm:"primaryKey;autoIncrement"`
	CargoID   int     `json:"cargo_id"`   // 关联货品 cargo
	Order     int     `json:"order"`      //生产顺序
	IsExact   bool    `json:"is_exact"`   //是否精准
	MaxUse    float64 `json:"max_use"`    //当用量是精准时， 取最大值， 否则为区间
	MinUse    float64 `json:"min_use"`    //当用量是精准时， 取最大值， 否则为区间
	MeasureID int     `json:"measure_id"` //计量单位 关联 measure
}
