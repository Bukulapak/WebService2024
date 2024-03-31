package _714220030

import (
    "testing"
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/stretchr/testify/assert"
    
)

func TestInsertTicket(t *testing.T) {
    ticket := Tiket{
        NomorTiket: "123456",
        Prioritas: "High",
        Kategori: "Bug",
        Deskripsi: "Something is not working",
        Status: "Open",
        Datetime: primitive.NewDateTimeFromTime(time.Now()),
        Penyelia: "Supervisor",
        AssignedTo: "Developer",
        Komentar: []Komentar{
            Komentar{
                Pengirim: "User",
                Isi: "This is a comment",
                Datetime: primitive.NewDateTimeFromTime(time.Now()),
            },
        },
    }
    
    insertedID := InsertTicket(ticket)
    assert.NotNil(t, insertedID, "InsertTicket should return an ID")
}

func TestGetAllTickets(t *testing.T) {
    tickets := GetAllTickets()
    assert.NotEmpty(t, tickets, "GetAllTickets should return non-empty list of tickets")
}

func TestGetTicketByID(t *testing.T) {
    ticket := Tiket{
        NomorTiket: "123456",
        Prioritas: "High",
        Kategori: "Bug",
        Deskripsi: "Something is not working",
        Status: "Open",
        Datetime: primitive.NewDateTimeFromTime(time.Now()),
        Penyelia: "Supervisor",
        AssignedTo: "Developer",
        Komentar: []Komentar{
            Komentar{
                Pengirim: "User",
                Isi: "This is a comment",
                Datetime: primitive.NewDateTimeFromTime(time.Now()),
            },
        },
    }
    
    insertedID := InsertTicket(ticket)
    assert.NotNil(t, insertedID, "InsertTicket should return an ID")
    
    insertedIDPrimitive, _ := insertedID.(primitive.ObjectID)
    retrievedTicket, err := GetTicketByID(insertedIDPrimitive)
    assert.NoError(t, err, "GetTicketByID should not return an error")
    assert.Equal(t, ticket.NomorTiket, retrievedTicket.NomorTiket, "Retrieved ticket should have the same ticket number")
}

func TestInsertKomentarAndGetKomentarByID(t *testing.T) {
    
    komentar := Komentar{
        Pengirim: "User",
        Isi: "This is a comment",
        Datetime: primitive.NewDateTimeFromTime(time.Now()),
    }

    
    insertedID, err := InsertKomentar(komentar)
    assert.NoError(t, err, "InsertKomentar should not return an error")
    assert.NotNil(t, insertedID, "InsertKomentar should return an ID")

    
    insertedIDPrimitive, _ := insertedID.(primitive.ObjectID)
    retrievedKomentar, err := GetKomentarByID(insertedIDPrimitive)
    assert.NoError(t, err, "GetKomentarByID should not return an error")
    assert.Equal(t, komentar.Pengirim, retrievedKomentar.Pengirim, "Retrieved komentar should have the same sender")
    assert.Equal(t, komentar.Isi, retrievedKomentar.Isi, "Retrieved komentar should have the same content")
    assert.Equal(t, komentar.Datetime, retrievedKomentar.Datetime, "Retrieved komentar should have the same datetime")
}

func TestInsertHelpdeskAndGetHelpdeskByID(t *testing.T) {
	
	helpdesk := Helpdesk{
		Nama:       "Helpdesk 1",
		Departemen: "IT",
		Ext:        "1234",
		Email:      "helpdesk1@example.com",
	}

	
	insertedID, err := InsertHelpdesk(helpdesk)
	assert.NoError(t, err, "InsertHelpdesk should not return an error")
	assert.NotNil(t, insertedID, "InsertHelpdesk should return an ID")

	
	insertedIDPrimitive, _ := insertedID.(primitive.ObjectID)
	retrievedHelpdesk, err := GetHelpdeskByID(insertedIDPrimitive)
	assert.NoError(t, err, "GetHelpdeskByID should not return an error")
	assert.Equal(t, helpdesk.Nama, retrievedHelpdesk.Nama, "Retrieved helpdesk should have the same name")
	assert.Equal(t, helpdesk.Departemen, retrievedHelpdesk.Departemen, "Retrieved helpdesk should have the same department")
	assert.Equal(t, helpdesk.Ext, retrievedHelpdesk.Ext, "Retrieved helpdesk should have the same extension")
	assert.Equal(t, helpdesk.Email, retrievedHelpdesk.Email, "Retrieved helpdesk should have the same email")
}