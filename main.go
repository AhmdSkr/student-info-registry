package main

import (

	"idea/students/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Routes
	e.GET("/students", controller.GetAllStudents)
	e.GET("/students/:id", controller.GetStudent)
	e.POST("/students", controller.CreateStudent)
	e.PUT("/students/:id", controller.ChangeStudent)
	e.PATCH("/students/:id", controller.PatchStudent)
	e.DELETE("/students/:id", controller.DeleteStudent)

	// Starting server
	e.Logger.Fatal(e.Start(":8080"))
}