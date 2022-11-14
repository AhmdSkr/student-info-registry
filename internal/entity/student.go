package entity

type StudentResponseForm struct {
	Id          int64 `json:"id" xml:"id"`
	StudentInfo StudentRequestForm
}

type StudentRequestForm struct {
	Firstname string `json:"firstname" xml:"firstname" form:"firstname" query:"firstname"`
	Lastname  string `json:"lastname" xml:"lastname" form:"lastname" query:"lastname"`
}
