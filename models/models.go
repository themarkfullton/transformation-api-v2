package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Clinic : Struct for trans-friendly clinics that will be stored in the database
type Clinic struct {
	ID primitive.ObjectID
	Name string
	State string
	Street string
	City string
	Zip string
	Phone string
	Description string
	Website string
	Image string
}

// Resource: Struct for trans-friendly resources (used by fashion, health, history, identity, etc)
type Resource struct {
	ID primitive.ObjectID
	Name string
	Source string
	Target string
	Description string
	Link string
}

