//Удалить i-ый элемент из слайса

package main

import (
	"fmt"
)

func main() {
	var (
		a []int = []int{5, 6, 7, 8, 9, 8, 2}
		i int   = 2
	)

	a = append(a[:i], a[i+1:]...) //удаление это добавление в слайс из начального слайса ДО i и после i

	fmt.Println(a)
}
