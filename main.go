package main

import (
	"vac/driver"
	"vac/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	driver.InitDB()
	e := routes.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Start("127.0.0.1:8000")

}
