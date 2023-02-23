package dto

import "github.com/timchine/jxc/model"

type CargoKindWithAttrs struct {
	CargoKind  model.CargoKind   `json:"cargo_kind"`
	CargoAttrs []model.CargoAttr `json:"cargo_attrs"`
}

type ReqSearchCargoKind struct {
	Search string `json:"search"`
	PageData
}

type ReqAddCargo struct {
	Measures        []model.Measure        `json:"measures"`
	Cargo           model.Cargo            `json:"cargo"`
	CargoAttrValues []model.CargoAttrValue `json:"cargo_attr_values"`
}

type ReqSearchCargo struct {
	Search string `json:"search"`
	Type   string `json:"type"`
	PageData
}
