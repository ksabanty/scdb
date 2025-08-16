package models

type Card struct {
	Player       string `json:"player"`
	Team         string `json:"team"`
	Manufacturer string `json:"manufacturer"`
	Collection   string `json:"collection"`
	Year         string `json:"year"`
	Value        string `json:"value"`
	Rookie       bool   `json:"rookie"`
	Numbered     bool   `json:"numbered"`
	Auto         bool   `json:"auto"`
}
