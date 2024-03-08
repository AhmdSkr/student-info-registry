package student

import "github.com/AhmdSkr/student-info-registry/internal/data"

type Entity struct {
	Id int64
	Information
}

type Information struct {
	data.Person
	data.Contact
}
