package main

import (
	"vac/driver"
	"vac/routes"
)

func main() {
	driver.InitDB()
	e:=routes.New()
	e.Start("127.0.0.1:8000")
	
}