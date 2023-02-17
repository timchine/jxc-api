package dto

import "github.com/timchine/jxc/model"

type ReqAddCargoKind struct {
	CargoKind  model.CargoKind   `json:"cargo_kind"`
	CargoAttrs []model.CargoAttr `json:"cargo_attrs"`
}
