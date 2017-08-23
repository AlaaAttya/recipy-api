package main

import (
	"fmt"
	"log"

	"github.com/alaaattya/recipy-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	}

	config := &bongo.Config{
		ConnectionString: viper.GetString("local.database.host"),
		Database:         viper.GetString("local.database.dbName"),
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	myPerson := &models.Person{
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
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
