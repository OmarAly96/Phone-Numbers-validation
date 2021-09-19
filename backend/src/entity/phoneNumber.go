package entity

type PhoneNumber struct {
	Id      int    `json:"id"`
	Country string `json:"country"`
	State   bool   `json:"state"`
	Code    string `json:"code"`
	Number  string `json:"number"`
}
