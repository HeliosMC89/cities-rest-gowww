package models

type City struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	CountryCode string `json:"country,omitempty"`
	District    string `json:"district,omitempty"`
	Population  int    `json:"population,omitempty"`
}
