package empanadasdb

import "database/sql"

type EmpanadasDBPort struct {
	db *sql.DB
}

type IEmpanadasDBPort interface {
}
