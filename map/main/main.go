package main

import "fmt"

func main() {
	colors := map[string]string{
		"black": "#FFFFF",
		"white": "#00000",
	}

	colors["red"] = "#ff0000"
	fmt.Println(colors)

	fmt.Println(colors["red"], colors["Reddd"])
	delete(colors, "red")
	fmt.Println(colors["red"], "Nil")

}
