//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
package main

import "fmt"

func main() {
	var ( //обозначаем два буферизированых канала определеной длины равной массиву
		a   = make(chan int, 8)
		b   = make(chan int, 8)
		arr = [8]int{2, 4, 6, 8, 10, 12, 14, 16}
	)
	//запускаем горутины на чтение массива и запись в первый канала
	//и на чтение из канала, возведение в квадрат и отправку во второй канал
	go send(a, arr)
	go square(a, b)

	for val := range b {
		fmt.Println(val) //читаем
	}

	fmt.Println("finished")
}

func send(a chan int, arr [8]int) {
	defer close(a) //закрываем канал после окончания записи
	for _, i := range arr {
		a <- i
	}
}

func square(c, b chan int) {
	defer close(b) //закрываем канал после окончания записи, чтению не помешает
	for val := range c {
		b <- val * val
	}
}
