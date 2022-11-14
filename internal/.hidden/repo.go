func (receiver *StudentRepository) ReadStudentByFullName(firstname string, lastname string) ([]entity.Student, error) {

	rows, err := receiver.Source.Query(`SELECT * FROM students WHERE FirstName = ? and LastName = ?`, firstname, lastname)
	if err != nil {
		return nil, err
	}

	var student entity.Student
	var students []entity.Student
	for rows.Next() {

		if err := rows.Scan(student.Id, student.FirstName, student.LastName); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}