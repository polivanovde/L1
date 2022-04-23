//Реализовать конкурентную запись данных в map
package main

import (
	"fmt"
	"sync"
	//     "time"
	"strconv"
)

var (
	k int
	m = make(map[int]string)
)

func service(wg *sync.WaitGroup, value string) {
	m[k] = value
	k++
	wg.Done()
}

func main() {
	fmt.Println("main() started")
	var (
		wg  sync.WaitGroup
		val string
	)

	for i := 0; i < 1000; i++ { //тысяча рутин записывают тысячу значений
		wg.Add(1)
		val = strconv.Itoa(i)
		go service(&wg, val)
	}

	wg.Wait()
	fmt.Println(m, len(m)) // проверям что длина мапы равна 1000 - ничего потеряно
}
