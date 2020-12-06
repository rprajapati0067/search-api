package domain

type Response struct {
	Restaurant    *[3]Common `json:"restaurant"`
	Transport     *[3]Common `json:"transport"`
	PetrolStation *[3]Common `json:"petrol_station"`
}
