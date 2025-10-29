package main

import (
	"fmt"
	"os"

	"tired/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run . 'fileName.txt'")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	farm, err := functions.ValidateFormat(string(data))
	if err != nil {
		fmt.Println("ERROR: invalid data format,", err)
		return
	}

	paths, assigned := functions.EdmondKarp(&farm)
	if len(paths) == 0 {
		fmt.Println("ERROR: not a single path goes from start room to end room")
		return
	}

	functions.MooveAnts(paths, farm.Antnumber, string(data), assigned)
}
