package main

import (
	"dino/dinowebportal"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	Webserver string `json:"webserver"` // uppercase because Webserver field will exported
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration) // create pointer of configuration type
	json.NewDecoder(file).Decode(config)
	fmt.Println("Starting web server on address", config.Webserver)
	dinowebportal.RunWebPortal(config.Webserver)
}
