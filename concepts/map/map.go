package mapitem

import "fmt"

func MainMap() {
	var colorMap map[string]string
	fmt.Println(colorMap)

	var colorMap2 = map[string]string{"white": "#FFFFFF", "black": "#000000", "red": "#FF0000", "blue": "#0000FF", "green": "#00FF00"}
	fmt.Println(colorMap2)

	var colorMap3 = make(map[string]string)
	colorMap3["white"] = "#FFFFFF"
	colorMap3["black"] = "#000000"
	colorMap3["red"] = "#FF0000"
	colorMap3["blue"] = "#0000FF"
	colorMap3["green"] = "#00FF00"
	fmt.Println(colorMap3)
	fmt.Println(colorMap["white"])

	for key, value := range colorMap3 {
		fmt.Println("Hex value of ", key, "is", value)
	}

	//update value from map
	colorMap3["white"] = "#F0600D"
	fmt.Println(colorMap3["white"])

	//delete a key:value pair
	delete(colorMap3, "white")
	fmt.Print(colorMap3)

}
