package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName, attr"`
	LastName  string   `xml:"lastName, attr"`
}

func IndexHandler(c *gin.Context) {
	c.XML(200, Person{FirstName: "Ridho", LastName: "Achmad"})
}

// func IndexHandler(c *gin.Context) {
// 	name := c.Params.ByName("name")
// 	c.JSON(200, gin.H{
// 		"message": "Test Gin FW " + name,
// 	})
// }

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	// router.Run(":8080")
	router.Run()
}
