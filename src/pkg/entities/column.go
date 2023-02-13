package entities

type ColumnType int

const (
	ColumnInteger ColumnType = iota
	ColumnText
	ColumnBoolean
	ColumnUUID
)

type ColumnQueryGenerator interface {
	GenerateColumnQuery()
}

type Column struct {
	Name       string     `json:"Name"`
	Size       int        `json:"Size"`
	ColumnType ColumnType `json:"ColumnType"`
}

func GenerateColumnQuery(column *Column) string {
	return "create column"
}
