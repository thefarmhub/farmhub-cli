package model

// Project represents a single project with minimal fields.
type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"mode"`
}
