package main

import (
	"fmt"
	"idea/students/data"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StudentController struct {
	Repo *data.StudentRepository
}

func ProvideStudentController(username, password string) (*StudentController, error) {
	repo, err := data.ProvideStudentRepository(username, password)
	if err != nil {
		return nil, err
	}
	return &StudentController{Repo: repo}, nil
}

func (receiver *StudentController) CreateStudent(ctx echo.Context) error {

	var student data.StudentRequestForm
	if err := ctx.Bind(&student); err != nil {
		return err
	}

	studentModel := data.StudentModel{StudentRequestForm: student}
	if err := receiver.Repo.CreateStudent(&studentModel); err != nil {
		return errorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, studentModel)
}

func (receiver *StudentController) GetAllStudents(ctx echo.Context) error {

	var students []data.StudentModel
	if err := receiver.Repo.ReadAllStudents(&students); err != nil {
		return errorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, students)
}
func (receiver *StudentController) GetStudent(ctx echo.Context) error {

	id, err := readId(ctx)
	if err != nil {
		return errorHandler(ctx, err)
	}

	var student data.StudentModel
	err = receiver.Repo.ReadStudentById(&student, id)
	if err != nil {
		return errorHandler(ctx, err)
	}
	return ctx.JSON(http.StatusOK, student)
}

func (receiver *StudentController) ChangeStudent(ctx echo.Context) error {

	id, err := readId(ctx)
	if err != nil {
		return errorHandler(ctx, err)
	}

	var student data.StudentRequestForm
	if err := ctx.Bind(&student); err != nil {
		return err
	}

	studentModel := data.StudentModel{StudentRequestForm: student}
	if err := receiver.Repo.UpdateStudentById(&studentModel, id); err != nil {
		return errorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, studentModel)
}

func (receiver *StudentController) DeleteStudent(ctx echo.Context) error {

	id, err := readId(ctx)
	if err != nil {
		return errorHandler(ctx, err)
	}

	if err := receiver.Repo.DeleteStudentById(id); err != nil {
		return errorHandler(ctx, err)
	}

	return ctx.NoContent(http.StatusOK)
}

// Reads an id of type int64 from Path Parameter
func readId(ctx echo.Context) (int64, error) {

	param := ctx.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func errorHandler(ctx echo.Context, err error) error {

	switch Err := err.(type) {

	case *strconv.NumError:
		switch Err.Unwrap().Error() {

		case strconv.ErrSyntax.Error():
			return ctx.String(http.StatusBadRequest, fmt.Sprintf("Invalid Parameter Syntax for Id! Received Id: %v\n", Err.Num))
		case strconv.ErrRange.Error():
			return ctx.String(http.StatusBadRequest, fmt.Sprintf("Id out of range! Received Id: %v\n", Err.Num))
		default:
			return Err
		}

	default:
		switch err {

		case data.ErrRecordNotFound:
			return ctx.String(http.StatusBadRequest, "No record found.\n")
		default:
			return err
		}

	}
}
