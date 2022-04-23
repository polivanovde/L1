//Реализовать бинарный поиск встроенными методами языка
package main

import "fmt"

func binSearch(search int, items []int) bool {

	//определяем начало и конец поиска
	low := 0
	high := len(items) - 1

	for low <= high {
		median := (low + high) / 2 //делим поиск поплам (bin)

		//если значение среднего ключа меньше искомого то смещаем начальную границу
		//(ищем правее мед), иначе смещаем правую (ищем левее мед)
		if items[median] < search {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	//ключ искомого значения - low
	//проверяем что не вышли за пределы среза
	if low == len(items) || items[low] != search {
		return false
	}
	fmt.Println("Позиция:", low)

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, -42, 63, 70, 100, 101}
	fmt.Printf("It`s found? %v\n", binSearch(105, items))
}
