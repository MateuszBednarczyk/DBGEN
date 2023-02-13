package entities

type Table struct {
	Name        string       `json:"Name"`
	PrimaryKey  ColumnType   `json:"PrimaryKey"`
	Columns     []Column     `json:"Columns"`
	ForeignKeys []ForeignKey `json:"foreign_keys"`
}
