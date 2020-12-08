package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Event stores information about a hackathon event
type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	URL         string    `json:"url"`
}

// Response formats for AMP expected response
type Response struct {
	Items []Event `json:"items"`
}

func main() {

	jsonFile, err := os.Open("all_hackathons.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Response
	json.Unmarshal([]byte(byteValue), &result.Items)
	log.Println("hola ke ase")
	r := gin.Default()

	r.GET("/all", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(200, result)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
