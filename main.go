package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/x/bsonx"
)

var contractAddress = "0xa30cf1135be5af62e412f22bd01069e2ceba8706"

func main() {
	s := NewStorage()
	c := NewColler(s)
	contactDetails := GetContactByContractAddress(c, contractAddress)
	fmt.Println(contactDetails)
	s.DB.Collection("contact").InsertOne(context.Background(), bsonx.MDoc{
		"contractAddress": bsonx.String(contractAddress),
		"email":           bsonx.String(contactDetails.Email),
		"twitter":         bsonx.String(contactDetails.Twitter),
		"linkedin":        bsonx.String(contactDetails.Linkedin),
		"discord":         bsonx.String(contactDetails.Discord),
		"opensea":         bsonx.String(contactDetails.Opensea),
	})
}
