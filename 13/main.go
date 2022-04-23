//Поменять местами два числа без создания временной переменной
package main

import "fmt"

func main() {
	var (
		a = [2]int{2, 4}
	)
	a[0], a[1] = a[1], a[0]

	fmt.Println(a)
}
