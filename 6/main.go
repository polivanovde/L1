//Реализовать все возможные способы остановки выполнения горутины
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		ticker          = time.NewTicker(1 * time.Second)
		done            = make(chan bool)
		to              = make(chan bool)
		choise          int
		ctx, cancelFunc = context.WithCancel(context.Background())
		termChan        = make(chan os.Signal, 1)
	)
	defer ticker.Stop()
	fmt.Print("Choise from 0 to 6: ")
	fmt.Scan(&choise)

	//останавливаемая рутина
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-done:
				return
			case <-to:
				goto FINISH
			}
		}
	FINISH:
		return
	}()

	switch choise {
	case 0:
		fmt.Println("Wait")
		time.Sleep(4 * time.Second)
		fmt.Println("Exit by return")
		done <- true
	case 1:
		fmt.Println("Press Ctrl+C")
		signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT)
		go cancelHandler(termChan, cancelFunc)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Exit by signal")
				return
			}
		}
	case 2:
		fmt.Println("Wait")
		time.Sleep(3 * time.Second)
		fmt.Println("Exit by goto")
		to <- true
	}

	fmt.Println("All stopped")
}

func cancelHandler(termChan chan os.Signal, cancelFunc func()) {
	<-termChan
	fmt.Println("*********************************\nShutdown signal received\n*********************************")
	cancelFunc()
}
