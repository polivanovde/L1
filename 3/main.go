//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов с использованием конкурентных вычислений.
package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	for fn := range square(arr) {
		fmt.Println(fn)
	}
}

func square(arr [5]int) <-chan int {
	c := make(chan int, 1)
	var sum int
	go func() { //генерируем данные в анонимной горутине чтобы программе не нужно было ожидать вычисления всех квадратов
		for _, i := range arr {
			sum += i * i
		}
		c <- sum
		close(c)
	}()

	return c
}
