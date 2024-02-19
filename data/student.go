package data

type StudentModel struct {
	Id int64 `gorm:"primarykey" json:"id" xml:"id"`
	StudentRequestForm
}

type StudentRequestForm struct {
	Firstname string `json:"firstname" xml:"firstname" form:"firstname" query:"firstname"`
	Lastname  string `json:"lastname" xml:"lastname" form:"lastname" query:"lastname"`
}
