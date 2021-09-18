package model

type PhoneNumber interface{}

type CreatePhoneNumberInput struct {
	PhoneNumber `json:"number"`
}

type CreatePhoneNumberOutput struct {
	Message string `json:"message"`
}
