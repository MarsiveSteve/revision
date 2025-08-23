package piscine

import "fmt"

func HashCode(dec string) string {
	length := len(dec)
	result := make([]byte, len(dec))

	for i, ch := range dec {
		hashed := (int(ch) + length) % 127
		if hashed < 33 {
			hashed += 33
		}
		result[i] = byte(hashed)
	}

	return string(result)
}

func main() {
	fmt.Println(HashCode("A"))           // Output: B
	fmt.Println(HashCode("AB"))          // Output: CD
	fmt.Println(HashCode("BAC"))         // Output: EDF
	fmt.Println(HashCode("Hello World")) // Output: Spwwz+bz}wo
}
