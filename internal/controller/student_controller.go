package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"idea/students/internal/entity"
	"idea/students/internal/repository"

	"github.com/labstack/echo/v4"
)

type StudentController struct {
	Repo *repository.StudentRepository
}

func ProvideStudentController(username, password string) (*StudentController, error) {
	repo, err := repository.ProvideStudentRepository(username, password)
	if err != nil {
		return nil, err
	}
	return &StudentController{Repo: repo}, nil
}

func (receiver *StudentController) CreateStudent(ctx echo.Context) error {

	var student entity.StudentRequestForm
	if err := ctx.Bind(&student); err != nil {
		return err
	}

	id, err := receiver.Repo.CreateStudent(student)
	if err != nil {
		return err
	}

	response := entity.StudentResponseForm{
		Id:          id,
		StudentInfo: student,
	}
	return ctx.JSON(http.StatusCreated, response)
}

func (receiver *StudentController) GetAllStudents(ctx echo.Context) error {

	students, err := receiver.Repo.ReadAllStudents()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, students)
}
func (receiver *StudentController) GetStudent(ctx echo.Context) error {

	id, err := readId(ctx)
	if err != nil {
		return err
	}

	student, err := receiver.Repo.ReadStudentById(int64(id))
	if err != nil {
		//return ctx.String(http.StatusNotFound, fmt.Sprintf("Could not find student record of id: %q\n", id))
		return err
	}

	return ctx.JSON(http.StatusOK, student)
}

func (receiver *StudentController) ChangeStudent(ctx echo.Context) error {

	id, err := readId(ctx)
	if err != nil {
		return err
	}

	var student entity.StudentRequestForm
	if err := ctx.Bind(&student); err != nil {
		return err
	}

	if err := receiver.Repo.UpdateStudentById(id, student); err != nil {
		return err
	}

	response := entity.StudentResponseForm{
		Id:          id,
		StudentInfo: student,
	}
	return ctx.JSON(http.StatusOK, response)
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

// helper functions
// reads an id of type int64 from Path Parameter
func readId(ctx echo.Context) (int64, error) {

	param := ctx.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func errorHandler(ctx echo.Context, err error) error {

	switch err {
	case strconv.ErrSyntax:
		return ctx.String(http.StatusBadRequest, "Invalid Parameter Syntax.")
	case sql.ErrNoRows:
		return ctx.String(http.StatusBadRequest, "No record found.")
	default:

		return err
	}
}
