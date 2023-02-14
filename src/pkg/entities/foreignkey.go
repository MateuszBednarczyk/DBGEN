package entities

type ForeignKey struct {
	Name      string `json:"Name"`
	Source    Column `json:"Source"`
	RelatedTo Column `json:"RelatedTo"`
}
