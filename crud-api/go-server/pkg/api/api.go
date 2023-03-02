package api

import "github.com/choisangh/crud-api/go-server/pkg/db"

type APIs struct {
	db *db.DBHandler
}

func NewAPI(handler *db.DBHandler) *APIs {
	a := APIs{db: handler}

	return &a
}
