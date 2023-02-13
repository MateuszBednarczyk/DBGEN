package entities

type ForeignKey struct {
	Name               string `json:"Name"`
	Source             Column `json:"Source"`
	SourceTableName    string `json:"SourceTableName"`
	RelatedTo          Column `json:"RelatedTo"`
	RelatedToTableName string `json:"RelatedToTableName"`
}
