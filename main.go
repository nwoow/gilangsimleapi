package main

import (
	"DT/Router"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	Router.Router()
}
