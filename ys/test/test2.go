package main

import "fmt"

func main() {
	maps := make(map[string]string)
	maps["boue"] = "sfffasf"

	value, esist := maps["boue"]

	if esist {
		fmt.Print(value)
	}

}
