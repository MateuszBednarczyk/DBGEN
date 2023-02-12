package entities

type ForeignKey struct {
	name               string
	source             Column
	sourceTableName    string
	relatedTo          Column
	relatedToTableName string
}
