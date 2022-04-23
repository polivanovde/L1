//Найти уникальные элементы двух неупорядоченных множеств
package main

import "fmt"

func main() {
	var (
		a    = []int{1, 2, 3, 4, 5, 6}
		b    = []int{7, 8, 4, 3, 6, 9, 1, 12}
		dubl bool
	)

	for _, i := range b { //перебираем один из слайсов
		dubl = false
		for _, j := range a { //проверяем наличие элемента в другом слайсе
			if i == j {
				dubl = true
			}
		}
		if dubl == false {
			a = append(a, i) //добавляем в первый слайс элементы второго если они отсутствуют
		}
	}

	//еще вариант сразу объединить в новый слайс
	//sl := append(a, b...)
	//и выполнить удаление дублей, но без вложенного перебора
	//r := removeDuplicates(sl)
	//fmt.Println(r)

	fmt.Println(a)
}

//func removeDuplicates(nums []int) []int {
//	for i, j := 0, 1; i < len(nums); {
//		if j < len(nums) && nums[j] == nums[i] && i != j {
//			nums = append(nums[0:j], nums[j+1:]...)
//		} else if j < len(nums) {
//			j++
//		}
//		if j == len(nums) {
//			i++
//			j = i + 1
//		}
//	}
//
//	return nums
//}
