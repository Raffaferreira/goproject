package jsonpack

import (
	"encoding/json"
	"fmt"
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
	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)

	fmt.Printf("Birds : %+v", birds)
}
