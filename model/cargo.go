package model

// 计量
type Measure struct {
	Model
	MeasureID int    `json:"measure_id"`
	CkID      int    `json:"ck_id"`
	IsBase    bool   `json:"is_base"`
	Unit      string `json:"unit"`
	Calc      string `json:"calc"`
	Status    int    `json:"status"`
}

// 货物 分类
type CargoKind struct {
	Model
	CkID   int    `json:"ck_id"`
	CkCode string `json:"ck_code"`
	CkName string `json:"ck_name"`
	Intro  string `json:"intro"`
	Type   int    `json:"type"` //原材料 半成品 成品
	Status int    `json:"status"`
}

// 货物属性
type CargoAttr struct {
	Model
	CaID      int    `json:"ca_id"`
	CkID      int    `json:"ck_id"`
	AttrName  string `json:"attr_name"`
	Type      int    `json:"type"` //1 选择 2 文本
	AttrValue string `json:"attr_value"`
	Status    int    `json:"status"`
}

// 货物
type Cargo struct {
	Model
	CargoID   int    `json:"cargo_id"`
	CkID      int    `json:"ck_id"`
	CargoName string `json:"cargo_name"`
	CargoCode string `json:"cargo_code"`
	Status    int    `json:"status"`
}

type CargoAttrValue struct {
	Model
	CavID     int    `json:"cav_id"`
	CaID      int    `json:"ca_id"`
	CargoID   int    `json:"cargo_id"`
	AttrName  string `json:"attr_name"`
	AttrValue string `json:"value"`
	Status    int    `json:"status"`
}
