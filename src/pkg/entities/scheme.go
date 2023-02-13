package entities

type Scheme struct {
	Name   string  `json:"Name"`
	Tables []Table `json:"Tables"`
}
