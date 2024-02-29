package student

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

const (
	driver_name             string = "mysql"
	default_user            string = "student-info-server"
	default_pass            string = "password"
	default_connection_type string = "tcp"
	default_address         string = "127.0.0.1:3306"
	default_db_name         string = "student_info"
)

const (
	query_create string = `
	INSERT INTO 
		students
			(FirstName,MiddleName,LastName,BirthDate,Gender,Phone,Address,Country)
	VALUES 
		(?,?,?,?,?,?,?,?)`
	query_read string = `
	SELECT 
		ID, FirstName, MiddleName, LastName, BirthDate, Gender, Phone, Address, Country
	FROM 
		students
	WHERE 
		ID = ?`
	query_read_all string = `
	SELECT 
		ID, FirstName, MiddleName, LastName, BirthDate, Gender, Phone, Address, Country 
	FROM
		students`
	query_update string = `
	UPDATE 
		students
	SET 
		FirstName = ?, MiddleName = ?, LastName = ?, BirthDate = ?, Gender = ?, Phone = ?, Address = ?, Country = ?
	WHERE 
		ID = ?`
	query_delete string = `DELETE FROM students WHERE ID = ?`
)

var source *sql.DB

func init() {

	const (
		schema_path string = "./internal/student/rsc/schema.sql"
		init_path   string = "./internal/student/rsc/init.sql"
	)

	var (
		cfg *mysql.Config
		dsn string
		err error
	)

	cfg = &mysql.Config{
		Net:    default_connection_type,
		User:   default_user,
		Passwd: "One-Two-Three-Four5",
		Addr:   "localhost:3306",
		DBName: "student_info",
	}
	dsn = cfg.FormatDSN()

	if source, err = sql.Open(driver_name, dsn); err != nil {
		log.Fatal(err)
	}

	if err = source.Ping(); err != nil {
		log.Fatal(err)
	}

	if err = run_sql_script(source, schema_path); err != nil {
		// Would rather warn on failure
		log.Fatal(err)
	}

	if err = run_sql_script(source, init_path); err != nil {
		// Would rather warn on failure
		log.Fatal(err)
	}
}

func Create(student Information) (int64, error) {

	const query string = query_create

	var (
		result sql.Result // query result
		err    error      // generic error
	)

	if result, err = source.Exec(
		query,
		student.Firstname, student.Middlename, student.Lastname,
		student.BirthDate, student.Gender,
		student.Phone, student.Address, student.Country,
	); err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func ReadAll() ([]Entity, error) {

	const (
		buffer_size      int    = 500
		buffer_max_index int    = buffer_size - 1
		query            string = query_read_all
	)

	var (
		rows   *sql.Rows = nil                         // query resultant rows
		err    error     = nil                         // generic error
		entity Entity    = Entity{}                    // binding student entity
		index  int                                     // row index modulo size of buffer
		list   []Entity                                // resultant student entity list
		buffer []Entity  = make([]Entity, buffer_size) // student entity buffer, for less append() calls
	)

	if rows, err = source.Query(query); err != nil {
		return nil, err
	}
	defer rows.Close()

	for index = 0; rows.Next(); {

		rows.Scan(
			&entity.Id,
			&entity.Firstname, &entity.Middlename, &entity.Lastname,
			&entity.BirthDate, &entity.Gender,
			&entity.Phone, &entity.Address, &entity.Country,
		)

		buffer[index] = entity
		index++

		if index == buffer_max_index {
			list = append(list, buffer...)
			index = 0
		}
	}
	if index != 0 {
		list = append(list, buffer[:index]...)
	}

	return list, nil
}

func Read(id int64) (*Entity, error) {

	const query string = query_read

	var (
		entity Entity   = Entity{} // binding student entity
		row    *sql.Row = nil      // query resultant row
	)

	row = source.QueryRow(query, id)

	if err := row.Scan(
		&entity.Id,
		&entity.Firstname, &entity.Middlename, &entity.Lastname,
		&entity.BirthDate, &entity.Gender,
		&entity.Phone, &entity.Address, &entity.Country,
	); err != nil {
		return nil, err
	}

	return &entity, nil
}

func Update(id int64, info Information) error {

	const query string = query_update

	var (
		result sql.Result = nil // query result
		err    error      = nil // generic error
	)

	if result, err = source.Exec(
		query,
		info.Firstname, info.Middlename, info.Lastname,
		info.BirthDate, info.Gender,
		info.Phone, info.Address, info.Country,
		id); err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
}

func Delete(id int64) error {

	const query string = query_delete

	var (
		result sql.Result = nil // query result
		err    error      = nil // generic error
	)

	if result, err = source.Exec(query, id); err != nil {
		return err
	}

	if effected, err := result.RowsAffected(); err != nil {
		return err
	} else if effected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
