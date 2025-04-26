package models

import "gorm.io/gorm"

type Flight struct {
	gorm.Model
	FlightNumber string
	Departure    string
	Arrival      string
	Status       string
}
