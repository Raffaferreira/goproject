package jsonpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type BirdList struct {
	Birds []Bird `json:"birds"`
}

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
	Dimensions  Dimensions
}

func Main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)

	json.Unmarshal([]byte(birdJson), &bird)

	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
}

// Convert an string to an object Struct
func JsonArray() {
	file2, err := ioutil.ReadFile("./files/birds.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	//birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`

	var birds []Bird
	json.Unmarshal([]byte(file2), &birds)

	fmt.Printf("Birds : %+v", birds)
}

func NestedJsonObject() {
	//birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	file3, err := ioutil.ReadFile("./files/birdnested.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var bird Bird
	json.Unmarshal([]byte(file3), &bird)

	fmt.Println(bird)
	fmt.Println(bird.Species)
	fmt.Println(bird.Description)
	fmt.Println(bird.Dimensions.Height)
	fmt.Println(bird.Dimensions.Width)
}

func UnstructuredDataJson() {
	content, err := ioutil.ReadFile("./files/config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload map[string]interface{}

	err = json.Unmarshal(content, &payload)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's print the unmarshalled data!
	log.Printf("origin: %s\n", payload["origin"])
	log.Printf("user: %s\n", payload["user"])
	log.Printf("status: %t\n", payload["active"])
}
