package main

import "github.com/choisangh/gin_project/pkg/router"

func main() {
	r := router.Router()

	r.Run(":8080")

}
