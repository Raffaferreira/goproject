package jsonpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type BirdList struct {
	BirdList []Bird `json:"birds_array`
}

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string `json: "species"`
	Description string `json: "description"`
	Dimensions  Dimensions
}

func Main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)

	json.Unmarshal([]byte(birdJson), &bird)

	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
}

func JsonArray() {
	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)

	fmt.Printf("Birds : %+v", birds)
}

func JsonArrayFile() {
	file, err := ioutil.ReadFile("birds.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var birds []Bird
	json.Unmarshal([]byte(file), &birds)

	for i := 0; i < len(birds); i++ {
		fmt.Println("Description : ", birds[i].Description)
		fmt.Println("Species : ", birds[i].Species)
	}
}

func JsonArrayFile2() {
	file, err := ioutil.ReadFile("birds.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	birds := BirdList{}

	json.Unmarshal([]byte(file), &birds)

	for i := 0; i < len(birds.BirdList); i++ {
		fmt.Println("Description : ", birds.BirdList[i].Description)
		fmt.Println("Species : ", birds.BirdList[i].Species)
	}
}
