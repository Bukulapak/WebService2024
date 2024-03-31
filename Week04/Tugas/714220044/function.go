package _714220044

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func ConnectionMonggo(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := ConnectionMonggo(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertEmployee(name string, position string, salary float64) (insertedID interface{}) {
	var emp Employee
	emp.Name = name
	emp.Position = position
	emp.Salary = salary
	return InsertOneDoc("PenggajianHonor", "employees", emp)
}

func InsertHonor(name string, amount float64) (insertedID interface{}) {
	var hono Honor
	hono.Name = name
	hono.Amount = amount
	return InsertOneDoc("PenggajianHonor", "honor", hono)
}

func InsertSalary(name Employee, mounth int, year int) (insertedID interface{}) {
    var sala Salary
    sala.Name = name
    sala.Month = mounth
    sala.Year = year
    return InsertOneDoc("PenggajianHonor", "salary", sala)
}


func GetAllEmployees(db string, collection string) ([]Employee, error) {

	var employees []Employee

	cursor, err := ConnectionMonggo(db).Collection(collection).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllEmployees: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var emp Employee
		if err := cursor.Decode(&emp); err != nil {
			return nil, fmt.Errorf("GetAllEmployees: %v", err)
		}
		employees = append(employees, emp)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("GetAllEmployees: %v", err)
	}

	return employees, nil
}

func GetAllSalary(db string, collection string) ([]Salary, error) {
	var salaries []Salary

	cursor, err := ConnectionMonggo(db).Collection(collection).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetSalary: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var sal Salary
		if err := cursor.Decode(&sal); err != nil {
			return nil, fmt.Errorf("GetSalary: %v", err)
		}
		salaries = append(salaries, sal)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("GetSalary: %v", err)
	}

	return salaries, nil
}

func GetAllHonor(db string, collection string) ([]Honor, error) {
    var honors []Honor
    
    cursor, err := ConnectionMonggo(db).Collection(collection).Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, fmt.Errorf("GetAllHonor: %v", err)
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var hon Honor
        if err := cursor.Decode(&hon); err != nil {
            return nil, fmt.Errorf("GetAllHonor: %v", err)
        }
        honors = append(honors, hon)
    }

    if err := cursor.Err(); err != nil {
        return nil, fmt.Errorf("GetAllHonor: %v", err)
    }

    return honors, nil
}
