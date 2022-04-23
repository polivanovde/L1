//Разработать программу, которая будет последовательно отправлять
//значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться
package main

import (
	"context"
	"fmt"
	"time"
)

const time_in_seconds = 5

func main() {
	//используем контекст для прерывания
	ctx, cancelFunc := context.WithCancel(context.Background())

	//функция таймера и прерывания ожидает в отдельной горутине
	go cancelTimer(cancelFunc)

	//читаем данные из канала последовательно без горутин (конкуренции)
	for r := range generate(ctx) {
		fmt.Println("Response ", r)
	}
}

func generate(ctx context.Context) <-chan int {
	var (
		i int
		c = make(chan int)
	)

	go func() {
		defer close(c)
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			time.Sleep(time.Second / 10)
			c <- i
			i++
		}
	}()

	return c
}
func cancelTimer(cancelFunc func()) {
	timer := time.NewTimer(time.Second * time_in_seconds)
	defer timer.Stop()

	<-timer.C
	fmt.Println("Finish it")
	cancelFunc()
}
