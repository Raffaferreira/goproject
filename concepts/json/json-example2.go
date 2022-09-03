package jsonpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func Main2() {
}

func JsonArrayFile() {
	file, err := ioutil.ReadFile("./files/birds.json")

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
