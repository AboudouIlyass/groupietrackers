package config

type Data struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Categorie   string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rate    `json:"rating"`
}
type Rate struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}
