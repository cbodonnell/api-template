package models

// User struct
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Discoverable bool   `json:"discoverable"`
	IRI          string `json:"url"`
}
