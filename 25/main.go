//Реализовать собственную функцию sleep

package main

import (
	"fmt"
    "time"
)

//используем функция After пакета time которая запускает таймер и возвращает канал Time
//выполнение является блокирующим и поэтому рутина будет ожидать
func SleepSeconds(x int) {
    timer := time.NewTimer(time.Second * time.Duration(x))
    <-timer.C

    //или
    //<-time.After(time.Second * time.Duration(x))
}

func main() {
    var n int

    fmt.Println("start main")
    SleepSeconds(5)
    fmt.Printf("finish after %d seconds\n", n)
}
