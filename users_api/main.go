package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID         string  `json:"id"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MaidenName string  `sjon:"maidenName"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Age        int32   `json:"age"`
	Weight     float64 `json:"weight"`
	Adress     Adress  `json:"address"`
}

type Adress struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
}

type Coordinates struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func listUsersFromStructuredJson(c *gin.Context) {
	josnFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}
	defer josnFile.Close()
	byteValue, _ := ioutil.ReadAll(josnFile)
	var users Users

	json.Unmarshal(byteValue, &users)

	c.IndentedJSON(http.StatusOK, users)
}

func listUsers(c *gin.Context) {
	// Open the jsonfile
	jsonFile, err := os.Open("users.json")

	// Display error when opening json file
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened the json file")

	defer jsonFile.Close()
	// Read to byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	//Unamrashall the bytearray contan contents of json files into the users struct
	json.Unmarshal([]byte(byteValue), &result)

	c.IndentedJSON(http.StatusOK, result["users"])
}

func main() {

	router := gin.Default()
	router.GET("users", listUsers)
	router.GET("known_users", listUsersFromStructuredJson)

	router.Run("localhost:8090")
}
