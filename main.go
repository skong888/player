package main

import (
	"fmt"
	"log"
	"player/tool"
)

func main() {
	response, err := tool.Updater("input.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
