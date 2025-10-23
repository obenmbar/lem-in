package main

import (
	"fmt"
	lemino "lemino/functions"
	"log"
	"os"
	"time"
)

func main() {
	sta := time.Now()
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
   lemino.Parser(Donee)
  elapsed := time.Since(sta)
  fmt.Println(elapsed)
   

}
