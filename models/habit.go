package models

// Habit struct
type Habit struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
