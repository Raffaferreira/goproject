package jsonpack

import (
	"encoding/json"
	"fmt"
)

// The keys need to be strings, the values can be
// any serializable value
func Main7() {
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total Birds": 2,
	}

	data, _ := json.Marshal(birdData)
	fmt.Println(string(data))
}
