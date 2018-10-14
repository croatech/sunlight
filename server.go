package main

import (
		"github.com/labstack/echo"
	_ "github.com/jinzhu/gorm/dialects/postgres"
		"feed/config/database"
)

func main() {
	database.Initialize()

	e := echo.New()

	e.Logger.Fatal(e.Start(":1323"))
}