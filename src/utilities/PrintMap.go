package utilities

import "fmt"

func PrintMap(boxes map[int][]string) {
	for key, value := range boxes {
		fmt.Printf("Box %d: %v\n", key, value)
	}
}
