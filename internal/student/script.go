package student

import (
	"bufio"
	"database/sql"
	"errors"
	"io"
	"os"
	"strings"
)

func run_sql_script(dbSource *sql.DB, scriptPath string) error {

	const stat_delim byte = ';'

	var (
		script io.ReadCloser // script reader
		input  bufio.Reader  // buffered script reader
		stmt   string        // statement extraction string
		err    error         // generic error
	)

	if script, err = os.Open(scriptPath); err != nil {
		return err
	}
	defer script.Close()

	input = *bufio.NewReader(script)

	// Extracting, Reading, and Executing a single statement from script
	for err = nil; !errors.Is(err, io.EOF); {

		stmt, err = input.ReadString(stat_delim)

		if strings.Trim(stmt, "; \n\r\f") == "" {
			continue
		}

		if _, err = dbSource.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}
