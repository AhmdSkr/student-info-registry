package data

import (
	"fmt"

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
	db.AutoMigrate(&StudentModel{})
	return &StudentRepository{Source: db}, nil
}

func (receiver *StudentRepository) CreateStudent(student *StudentModel) error {
	return receiver.Source.Create(&student).Error
}
func (receiver *StudentRepository) ReadAllStudents(students *[]StudentModel) error {
	return receiver.Source.Find(students).Error
}
func (receiver *StudentRepository) ReadStudentById(student *StudentModel, id int64) error {
	student.Id = id
	return receiver.Source.First(student).Error
}
func (receiver *StudentRepository) UpdateStudentById(student *StudentModel, id int64) error {
	student.Id = id
	if err := receiver.ReadStudentById(&StudentModel{}, id); err != nil {
		return err
	}
	return receiver.Source.Save(student).Error
}
func (receiver *StudentRepository) DeleteStudentById(id int64) error {
	if err := receiver.ReadStudentById(&StudentModel{}, id); err != nil {
		return err
	}
	return receiver.Source.Delete(&StudentModel{}, id).Error
}
