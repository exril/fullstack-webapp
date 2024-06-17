package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

// func structers() {
	type User struct {
		FirstName string `json:"firstname" bson:"firstname"`
		LastName  string `json:"lastname" bson:"lastname"`
		Email     string `json:"email" bson:"email"`
		Password  string `json:"password" bson:"password"`
	}

	type Todo struct {
		ID        primitive.ObjectID `json:"_id,omitempty"`
		Completed bool               `json:"completed"`
		Body      string             `json:"body"`
	}

// }