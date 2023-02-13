package entities

type Database struct {
	Name   string `json:"Name"`
	Scheme Scheme `json:"Scheme"`
}
