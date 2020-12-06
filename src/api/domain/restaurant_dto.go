package domain

type Restaurant struct {
	Id            string       `json:"id"`
	Distance      int          `json:"distance"`
	Title         string       `json:"title"`
	Icon          string       `json:"icon"`
	Vicinity      string       `json:"vicinity"`
	AverageRating float64      `json:"average_rating"`
	Position      []float64    `json:"position"`
	OpeningHours  OpeningHours `json:"openingHours"`
}

type OpeningHours struct {
	Text   string `json:"text"`
	Label  string `json:"label"`
	IsOpen bool   `json:"isOpen"`
}
