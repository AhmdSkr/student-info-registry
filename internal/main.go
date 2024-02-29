package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AhmdSkr/student-info-registry/internal/student"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func init() {

	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/students", ReadAll)
	e.GET("/students/:id", Read)
	e.POST("/students", Create)
	e.PUT("/students/:id", Update)
	e.DELETE("/students/:id", Delete)
}

func main() {
	if err := e.Start("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}

func Create(ctx echo.Context) error {

	var (
		id   int64               // student entity id, generated from database on creation
		info student.Information // student information, read from request payload
		url  string              // newly created student resource locator
		err  error               // generic error
	)

	// binding student information to request paylaod
	if err := ctx.Bind(&info); err != nil {
		return err
	}

	// creating a new student entity in the database
	if id, err = student.Create(info); err != nil {
		return err
	}

	// generate redirect url to newly created resource
	url = ctx.Echo().URL(Read, id)

	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func ReadAll(ctx echo.Context) error {

	var (
		students []student.Entity // student entity list
		err      error            // generic error
	)

	// fetching student entities from database
	if students, err = student.ReadAll(); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, students)
}

func Read(ctx echo.Context) error {

	var (
		id     int64           // student entity id, read from path parameter
		entity *student.Entity // student entity to be read
		err    error           // generic error
	)

	param := ctx.Param("id")

	// parsing path parameter
	if id, err = strconv.ParseInt(param, 10, 64); err != nil {
		return err
	}

	// fetching student entity from database
	if entity, err = student.Read(id); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, entity)
}

func Update(ctx echo.Context) error {

	var (
		id   int64               // student entity id, read from path parameter
		info student.Information // student information data structure
		url  string              // newly updated student resource locator
		err  error               // generic error
	)

	param := ctx.Param("id")

	// parsing path parameter
	if id, err = strconv.ParseInt(param, 10, 64); err != nil {
		return err
	}

	// binding request payload
	if err := ctx.Bind(&info); err != nil {
		return err
	}

	// processing request
	if err := student.Update(id, info); err != nil {
		return err
	}

	// generate redirect url to newly updated resource
	url = ctx.Echo().URL(Read, id)

	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func Delete(ctx echo.Context) error {

	var (
		id    int64  // student entity id, read from path parameter
		param string // id query parameter string
		err   error  // generic error
	)

	param = ctx.Param("id")

	// parsing path parameter
	if id, err = strconv.ParseInt(param, 10, 64); err != nil {
		return err
	}

	// deleting student from database
	if err = student.Delete(id); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
