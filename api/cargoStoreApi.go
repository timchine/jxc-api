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
