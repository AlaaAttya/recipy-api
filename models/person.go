package models

import "github.com/go-bongo/bongo"

// Person model
type Person struct {
	bongo.DocumentBase `bson:",inline"`
	FirstName          string
	LastName           string
	Gender             string
}
