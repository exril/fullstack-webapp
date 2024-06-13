package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//func todoStruct() {
	type Todo struct {
		ID        primitive.ObjectID `json:"_id,omitempty"`
		Completed bool               `json:"completed"`
		Body      string             `json:"body"`
	}
//	fmt.Println("Structure Loaded")
//}
