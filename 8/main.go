//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		n int64
		k int
	)
	for {
		fmt.Print("Enter max 19 digit number: ")
		fmt.Scanln(&n)
		fmt.Print("Enter max 10 digit number: ")
		fmt.Scanln(&k)
		new := n ^ (1 << (k - 1))
		fmt.Printf("Bin    : %s\n", strconv.FormatInt(int64(n), 2))
		fmt.Printf("New    : %d\n", new)
		fmt.Printf("Bin New: %s\n", strconv.FormatInt(int64(new), 2))
	}
}
