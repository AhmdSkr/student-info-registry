package entity

type Student struct {

	Id			string	`json:"id"			xml:"id"`
	FirstName	string	`json:"firstname"	xml:"firstname"	form:"firstname"	query:"firstname"`
	LastName	string	`json:"lastname"	xml:"lastname"	form:"lastname"		query:"lastname"`
}
