package _714220030

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) *mongo.Database {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
    if err != nil {
        fmt.Printf("MongoConnect: %v\n", err)
        return nil
    }
    return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) interface{} {
    insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
    if err != nil {
        fmt.Printf("InsertOneDoc: %v\n", err)
        return nil
    }
    return insertResult.InsertedID
}

func InsertTicket(ticket Tiket) interface{} {
    ticket.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
    return InsertOneDoc("jul2024", "tickets", ticket)
}

func GetAllTickets() []Tiket {
    var tickets []Tiket
    collection := MongoConnect("jul2024").Collection("tickets")
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        fmt.Println("GetAllTickets :", err)
        return nil
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var ticket Tiket
        if err := cursor.Decode(&ticket); err != nil {
            fmt.Println("Error decoding ticket:", err)
            continue
        }
        tickets = append(tickets, ticket)
    }
    if err := cursor.Err(); err != nil {
        fmt.Println("Error with cursor:", err)
    }
    return tickets
}

func GetTicketByID(id primitive.ObjectID) (Tiket, error) {
    var ticket Tiket
    collection := MongoConnect("jul2024").Collection("tickets")
    filter := bson.M{"_id": id}
    err := collection.FindOne(context.Background(), filter).Decode(&ticket)
    if err != nil {
        fmt.Printf("Error retrieving ticket: %v\n", err)
        return Tiket{}, err
    }
    return ticket, nil
}

func InsertKomentar(komentar Komentar) (interface{}, error) {
    
    if komentar.Datetime == 0 {
        komentar.Datetime = primitive.NewDateTimeFromTime(time.Now())
    }

    insertResult, err := MongoConnect("jul2024").Collection("komentar").InsertOne(context.TODO(), komentar)
    if err != nil {
        fmt.Printf("InsertKomentar: %v\n", err)
        return nil, err
    }
    return insertResult.InsertedID, nil
}

func GetKomentarByID(komentarID primitive.ObjectID) (Komentar, error) {
    var komentar Komentar

    filter := bson.M{"_id": komentarID}
    err := MongoConnect("jul2024").Collection("komentar").FindOne(context.Background(), filter).Decode(&komentar)
    if err != nil {
        fmt.Printf("GetKomentarByID: %v\n", err)
        return Komentar{}, err
    }
    return komentar, nil
}

func InsertHelpdesk(helpdesk Helpdesk) (interface{}, error) {
	insertResult, err := MongoConnect("jul2024").Collection("helpdesk").InsertOne(context.TODO(), helpdesk)
	if err != nil {
		fmt.Printf("InsertHelpdesk: %v\n", err)
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func GetHelpdeskByID(id primitive.ObjectID) (Helpdesk, error) {
	var helpdesk Helpdesk
	filter := bson.M{"_id": id}
	err := MongoConnect("jul2024").Collection("helpdesk").FindOne(context.Background(), filter).Decode(&helpdesk)
	if err != nil {
		fmt.Printf("GetHelpdeskByID: %v\n", err)
		return Helpdesk{}, err
	}
	return helpdesk, nil
}
