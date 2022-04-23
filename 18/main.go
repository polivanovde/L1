//Реализовать структуру-счетчик, которая будет инкрементироваться
//в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика
package main

import (
	"fmt"
	"sync"
)

type Queue struct { //определяем структуру счетчика
	count int
}

func (q *Queue) push(i int) { //метод получает значение которое прибавляет к счетчику
	q.count = q.count + i
}

func main() {
	var (
		wg    sync.WaitGroup
		n     int = 200
		queue     = new(Queue)
	)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() { //запускаем все операции увеличения счетчика в горутинах
			queue.push(i)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(queue.count) //значение суммы счетчика будет разным т.к запущено N горутин
}
