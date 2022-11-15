package repository

import (
	"fmt"
	"idea/students/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type StudentRepository struct {
	Source *gorm.DB
}

func connect(username, password, connectType, address, databaseName string) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%v:%v@%v(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, connectType, address, databaseName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	db.AutoMigrate(&model.StudentModel{})
	return &StudentRepository{Source: db}, nil
}

func (receiver *StudentRepository) CreateStudent(student *model.StudentModel) error {
	return receiver.Source.Create(&student).Error
}
func (receiver *StudentRepository) ReadAllStudents(students *[]model.StudentModel) error {
	return receiver.Source.Find(students).Error
}
func (receiver *StudentRepository) ReadStudentById(student *model.StudentModel, id int64) error {
	student.Id = id
	return receiver.Source.First(student).Error
}
func (receiver *StudentRepository) UpdateStudentById(student *model.StudentModel, id int64) error {
	student.Id = id
	if err := receiver.ReadStudentById(&model.StudentModel{}, id); err != nil {
		return err
	}
	return receiver.Source.Save(student).Error
}
func (receiver *StudentRepository) DeleteStudentById(id int64) error {
	if err := receiver.ReadStudentById(&model.StudentModel{}, id); err != nil {
		return err
	}
	return receiver.Source.Delete(&model.StudentModel{}, id).Error
}
