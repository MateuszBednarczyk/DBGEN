package entities

type Table struct {
	primaryKey Column
	columns    []Column
}

