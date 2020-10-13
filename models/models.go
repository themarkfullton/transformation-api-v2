package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Clinic : Struct for trans-friendly clinics that will be stored in the database
type Clinic struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	State       string             `bson:"state,omitempty" json:"state,omitempty"`
	Street      string             `bson:"street,omitempty" json:"street,omitempty"`
	City        string             `bson:"city,omitempty" json:"city,omitempty"`
	Zip         string             `bson:"zip,omitempty" json:"zip,omitempty"`
	Phone       string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Description string             `bson:"services,omitempty" json:"services,omitempty"`
	Website     string             `bson:"website,omitempty" json:"website,omitempty"`
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
}

// Resource: Struct for trans-friendly resources (used by fashion, health, history, identity, etc)
type Resource struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Source      string             `bson:"source,omitempty" json:"source,omitempty"`
	Target      string             `bson:"target,omitempty" json:"target,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Link        string             `bson:"link,omitempty" json:"link,omitempty"`
}

// Announcement : Used to display information on the splash page
type Announcement struct {
	ID			primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title    	string             `bson:"title,omitempty" json:"title,omitempty"`
	Author  	string             `bson:"author,omitempty" json:"author,omitempty"`
	Content 	string             `bson:"content,omitempty" json:"content,omitempty"`
	DateAdded	string				`bson:"dateAdded,omitempty" json:"dateAdded,omitempty"`
}