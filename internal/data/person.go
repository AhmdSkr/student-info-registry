package data

type Named struct {
	Firstname  string
	Middlename string
	Lastname   string
}

const (
	P_GENDER_MALE = iota
	P_GENDER_FEMALE
)

type Person struct {
	Named
	BirthDate string
	Gender    int8
}

type Contact struct {
	Phone   string
	Address string
	Country string
}
