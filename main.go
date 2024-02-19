package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ctrl, err := ProvideStudentController("api.user", "1-1-15@A")
	if err != nil {
		log.Fatal(err)
	}

	// Routes
	e.GET("/students", ctrl.GetAllStudents)
	e.GET("/students/:id", ctrl.GetStudent)
	e.POST("/students", ctrl.CreateStudent)
	e.PUT("/students/:id", ctrl.ChangeStudent)
	e.DELETE("/students/:id", ctrl.DeleteStudent)
	//e.PATCH("/students/:id", control.PatchStudent)

	// Starting server
	e.Logger.Fatal(e.Start(":8080"))
}
