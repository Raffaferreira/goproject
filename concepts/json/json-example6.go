package jsonpack

import (
	"encoding/json"
	"fmt"
)

// we can set the "omitempty" property as part of the JSON tag
func Main6() {
	pigeon := &Bird{
		Species: "Pigeon",
	}

	data, _ := json.Marshal(pigeon)

	fmt.Println(data)
}
