//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var (
		wg           = &sync.WaitGroup{}
		workersCount int
		res          = make(chan interface{})
	)

	fmt.Print("Workers count: ")
	fmt.Scan(&workersCount)
	if workersCount > 0 {
		//использую контекст для остановки горутин т.к контекст может остановить
		//как и главную так и отдельные горутины куда он был передан
		ctx, cancelFunc := context.WithCancel(context.Background())
		termChan := make(chan os.Signal, 1)

		signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT)
		go cancelHandler(termChan, cancelFunc)

		wg.Add(workersCount + 1)
		sender := Sender(ctx, wg)
		for i := 0; i < workersCount; i++ {
			go Pool(i, sender, res, ctx, wg)
		}

		go func() {
			defer close(res)
			for r := range res {
				fmt.Printf("Recieved mess: %v\n", r)
			}
		}()
	} else {
		fmt.Println("'Workers count' is not number or less than 0")
	}

	wg.Wait()
}

func Sender(ctx context.Context, wg *sync.WaitGroup) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer wg.Done()
		defer close(out)
		var i int
		for {
			//ожидаем завершения и прерываем отправку в канал данных
			select {
			case <-ctx.Done():
				return
			default:
			}
			i++
			fmt.Printf("Send mess: %v\n", i)
			time.Sleep(time.Second / 10)
			out <- i
		}
	}()

	return out
}

func Pool(id int, in <-chan interface{}, res chan<- interface{}, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range in {
		//ожидаем завершения и закрываем оставшиеся воркеры
		select {
		case <-ctx.Done():
			fmt.Printf("Worker-reader closed: %d\n", id)
			return
		default:
		}
		fmt.Printf("Worker-reader id: %d\n", id)
		res <- data
	}
}

//хендлер прерывания программы который ожидает из канала соответствующий сигнал
func cancelHandler(termChan chan os.Signal, cancelFunc func()) {
	<-termChan
	fmt.Println("*********************************\nShutdown signal received\n*********************************")
	cancelFunc()
	fmt.Println("Workers exiting")
}
