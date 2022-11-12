package controller

import (
	"fmt"
	"sync"
	"net/http"
	
	"idea/students/entity"

	"github.com/labstack/echo/v4"
)

var sequence func() string = func() (func() string) {

	var index int = 3
	return func() string {index++; return fmt.Sprintf("%06d", index);}
}()

var mtx sync.Mutex
var source map[string]*entity.Student = map[string]*entity.Student{

	"000001" : {"000001","alpha","one"},
	"000002" : {"000002","beta","two"},
	"000003" : {"000003","gamma","three"},
}

func GetAllStudents(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, source)
}
func GetStudent(ctx echo.Context) error {
	
	id := ctx.Param("id")
	
	mtx.Lock()
	defer mtx.Unlock()

	if student := source[id]; student != nil {
		return ctx.JSON(http.StatusOK, student)
	}
	return ctx.String(http.StatusNotFound, fmt.Sprintf("Could not find student record of id: %q\n", id))
}
func CreateStudent(ctx echo.Context) error {

	var student entity.Student
	
	if err := ctx.Bind(&student); err != nil {
		return err
	}
	
	var Id string
	for Id = sequence(); source[Id] != nil; {}
	
	student.Id = Id
	source[Id] = &student

	return ctx.JSON(http.StatusCreated, student)
}
func ChangeStudent(ctx echo.Context) error {

	var student entity.Student
	id := ctx.Param("id")
	
	if err := ctx.Bind(&student); err != nil {
		return err
	}
	if source[id] == nil {
		return ctx.String(http.StatusNotFound, fmt.Sprintf("Could not find student record of id: %q\n", id))
	}

	student.Id = id
	source[id] = &student

	return ctx.JSON(http.StatusOK, student)
}
func PatchStudent(ctx echo.Context) error {

	id := ctx.Param("id")
	
	if source[id] == nil {
		return ctx.String(http.StatusNotFound, fmt.Sprintf("Could not find student record of id: %q\n", id))
	}
	if firstname := ctx.QueryParam("firstname"); firstname != "" {
		source[id].FirstName = firstname
	}
	if lastname := ctx.QueryParam("lastname"); lastname != "" {
		source[id].LastName = lastname
	}
	return ctx.JSON(http.StatusOK, source[id])
}
func DeleteStudent(ctx echo.Context) error {

	id := ctx.Param("id")
	delete(source,id)
	return ctx.NoContent(http.StatusNoContent)
}