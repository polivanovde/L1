//Написать программу, которая конкурентно рассчитает
//значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	for fn := range square(arr) {
		fmt.Println(fn)
	}
}

func square(arr [5]int) <-chan int {
	c := make(chan int, 5)

	go func() { //генерируем данные в анонимной горутине чтобы программе не нужно было ожидать вычисления всех квадратов
		defer close(c)
		for _, i := range arr {
			c <- i * i
		}
	}()

	return c
}
