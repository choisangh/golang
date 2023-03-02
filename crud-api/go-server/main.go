package main

import (
	"fmt"

	"github.com/choisangh/crud-api/go-server/pkg/api"
	"github.com/choisangh/crud-api/go-server/pkg/db"
	"github.com/choisangh/crud-api/go-server/pkg/router"
)

func main() {
	dbHandler, err := db.NewAndConnectGorm("user:password@tcp(localhost:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error")
	}
	apis := api.NewAPI(dbHandler)
	r := router.Router(apis)

	r.Run(":8080")
}
