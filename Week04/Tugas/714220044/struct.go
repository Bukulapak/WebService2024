package _714220044

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Position string             `bson:"position"`
	Salary   float64         `bson:"salary"`
}

type Honor struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string            `bson:"name"`
	Amount float64            `bson:"amount"`
}

type Salary struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	EmployeeID primitive.ObjectID `bson:"employee_id"`
	Name	   Employee			`bson:"name"`
	Month      int                `bson:"month"`
	Year       int                `bson:"year"`
}

type Payment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	RecipientID primitive.ObjectID `bson:"recipient_id"`
	Date        string             `bson:"date"`
	Amount      float64            `bson:"amount"`
}
