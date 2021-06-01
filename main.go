package main

import (
	"net/http"

	"github.com/Mersock/golang-echo-postgres-restful-api/controller"
	"github.com/Mersock/golang-echo-postgres-restful-api/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//echo  instance
	e := echo.New()
	storage.NewDB()

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//router
	e.GET("/", hello)
	e.GET("/students", controller.GetStudents)

	//start server
	e.Logger.Fatal(e.Start((":3000")))
}

//handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}
