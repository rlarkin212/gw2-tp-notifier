package models

type Item struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   int64  `json:"id"`
	Icon string `json:"icon"`
	Sale Sale
}
