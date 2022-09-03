package jsonpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func Main3() {
}

func JsonArrayFile2() {
	file, err := ioutil.ReadFile("./files/birds.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	birds := BirdList{}

	json.Unmarshal([]byte(file), &birds)

	for i := 0; i < len(birds.Birds); i++ {
		fmt.Println("Description : ", birds.Birds[i].Description)
		fmt.Println("Species : ", birds.Birds[i].Species)
	}
}
