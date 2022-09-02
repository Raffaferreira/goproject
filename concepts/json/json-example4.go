package jsonpack

import (
	"encoding/json"
	"fmt"
)

// If you don’t know the structure of your JSON properties beforehand, you cannot use structs to unmarshal your data.
func Main4() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{})

	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
}
