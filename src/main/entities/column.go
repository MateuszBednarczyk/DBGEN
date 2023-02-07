package entities

type ColumnType uint

const (
	ColumnInteger ColumnType = iota
	ColumnText    ColumnType = iota
	ColumnUUID    ColumnType = iota
)

type Column struct {
	name       string
	size       uint
	columnType ColumnType
}
