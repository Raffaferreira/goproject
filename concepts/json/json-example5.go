package jsonpack

import (
	"encoding/json"
	"fmt"
)

func Main5() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))
}
