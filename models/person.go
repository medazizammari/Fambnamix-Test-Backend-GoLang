package models

type Person struct {
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	DateOfBirth string `json:"date_of_birth"`
	Address string `json:"address"`
	Age int `json:"age"`
	PhotoId string `json:"photo_id"`
}