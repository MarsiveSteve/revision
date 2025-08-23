package main

import (
	"fmt"
	"os"
)

func main() {
	argz := os.Args
	count := 0
	for _, element := range argz {
		if element == "01" || element == "galaxy" || element == "galaxy 01" {
			count++
		}
	}
	if count > 0 {
		fmt.Println("Alert!!!")
	}
}
