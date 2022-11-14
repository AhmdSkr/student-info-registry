package repository

import (
	"database/sql"
	"idea/students/internal/entity"

	"github.com/go-sql-driver/mysql"
)

type StudentRepository struct {
	Source *sql.DB
}

func connect(username, password, connectType, address, databaseName string) (*sql.DB, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 username,
		Passwd:               password,
		Net:                  connectType,
		Addr:                 address,
		DBName:               databaseName,
		AllowNativePasswords: true,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
func ProvideStudentRepository(username, password string) (*StudentRepository, error) {

	db, err := connect(
		username,
		password,
		"tcp",
		"127.0.0.1:3306",
		"Students",
	)
	if err != nil {
		return nil, err
	}

	return &StudentRepository{Source: db}, nil
}

func (receiver *StudentRepository) CreateStudent(student entity.StudentRequestForm) (int64, error) {

	result, err := receiver.Source.Exec(` INSERT INTO students (FirstName,LastName) VALUES (?,?);`, student.Firstname, student.Lastname)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (receiver *StudentRepository) ReadAllStudents() ([]entity.StudentResponseForm, error) {

	rows, err := receiver.Source.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var student entity.StudentResponseForm
	var students []entity.StudentResponseForm
	for rows.Next() {
		rows.Scan(&student.Id, &student.StudentInfo.Firstname, &student.StudentInfo.Lastname)
		students = append(students, student)
	}

	return students, nil
}

func (receiver *StudentRepository) ReadStudentById(id int64) (*entity.StudentResponseForm, error) {

	var student entity.StudentResponseForm
	row := receiver.Source.QueryRow(`SELECT * FROM students WHERE ID = ?`, id)

	if err := row.Scan(&student.Id, &student.StudentInfo.Firstname, &student.StudentInfo.Lastname); err != nil {
		return nil, err
	}
	return &student, nil
}

func (receiver *StudentRepository) UpdateStudentById(id int64, studentInfo entity.StudentRequestForm) error {

	result, err := receiver.Source.Exec(`UPDATE students SET FirstName = ?, LastName = ? WHERE ID = ?`, studentInfo.Firstname, studentInfo.Lastname, id)
	if err != nil {
		return err
	}
	effected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if effected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (receiver *StudentRepository) DeleteStudentById(id int64) error {

	result, err := receiver.Source.Exec(`DELETE FROM students WHERE ID = ?`, id)
	if err != nil {
		return err
	}

	effected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if effected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
