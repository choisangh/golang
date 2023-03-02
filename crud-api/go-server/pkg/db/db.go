package db

import (
	"github.com/choisangh/crud-api/go-server/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	gDB *gorm.DB
}

func NewAndConnectGorm(dsn string) (*DBHandler, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.Board{})

	dbHandler := DBHandler{
		gDB: db,
	}

	return &dbHandler, err
}
