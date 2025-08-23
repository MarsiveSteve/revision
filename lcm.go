// //Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

//     The hash equation is computed as follows:

// (ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

//     If the resulting character is unprintable add 33 to it.

// package main

// import "fmt"

// func main() {
// 	str1 := "hella there!"
// 	str := "there! hello"
// 	for i := range str1 {
// 		fmt.Print(string(str[i]))
// 	}
// 	fmt.Println()
// }

package main

import "fmt"

func main() {
	
	fmt.Println(Lcm(2, 7))
	fmt.Println(Lcm(0, 4))

	
}

func Lcm(a, b int) int {
product := a * b

	for b != 0 {
		a, b = b, a%b
	}

	return product / a
}
