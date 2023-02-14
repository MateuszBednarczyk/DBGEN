package entities

type Database struct {
	Name        string       `json:"Name"`
	Tables      []Table      `json:"Tables"`
	ForeignKeys []ForeignKey `json:"ForeignKeys"`
}
