package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		
		return
	}
	Str := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]
	result := ""

	for i := range Str {
		ch := Str[i]
		if string(ch) == old {
			result += new 
			continue
		}
		result += string(ch)
	}
	fmt.Println(result)
}
