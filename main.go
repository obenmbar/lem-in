package main

import (
	"fmt"
	"log"
	"os"

	lemino "lemino/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE/go run main.go exemple.txt")
		return
	}
	file := os.Args[1]

	Data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("error en readfile", err)
	}
	Donee := string(Data)
	fmt.Println(Donee)
	fmt.Println()
	lemino.Parser(Donee)
}
