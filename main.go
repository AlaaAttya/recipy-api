package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
)

func main() {
	config := &bongo.Config{
		ConnectionString: "192.168.99.100:37017",
		Database:         "bongotest",
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	myPerson := &Person{
		FirstName: "Testy",
		LastName:  "McGee",
		Gender:    "male",
	}
	errs := connection.Collection("people").Save(myPerson)

	if vErr, ok := errs.(*bongo.ValidationError); ok {
		fmt.Println("Validation errors are:", vErr.Errors)
	} else {
		fmt.Println("Got a real error:")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
