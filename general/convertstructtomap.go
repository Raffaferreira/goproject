package general

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Name  string
	Score int
}

func Main6() {
	ms := MyStruct{Name: "Rafael", Score: 19}

	var mapItem map[string]interface{}
	data, _ := json.Marshal(ms)
	json.Unmarshal(data, &mapItem)

	fmt.Println(mapItem["Name"])
	fmt.Println(mapItem["Score"])
	fmt.Println(mapItem)
}
