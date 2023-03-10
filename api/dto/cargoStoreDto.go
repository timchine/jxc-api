package dto

type ReqSearchCargoStore struct {
	CargoKindType int    `json:"cargo_kind_type"` //1 物料 2制品
	CargoKindID   int    `json:"cargo_kind_id"`   //种类ID
	Search        string `json:"search"`          // 查询条件 货品名称或 货品code
}
