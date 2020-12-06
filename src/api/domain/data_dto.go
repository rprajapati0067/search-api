package domain

type Data struct {
	Result Result `json:"results"`
}

type Result struct {
	Next  string       `json:"next"`
	Items []Restaurant `json:"items"`
}
