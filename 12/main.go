//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество
package main

import "fmt"

func main() {
	var (
		a = []string{"cat", "cat", "dog", "cat", "tree"}
		b = removeDuplicates(a)
	)

	fmt.Println(b)
}
func removeDuplicates(strs []string) []string {
	for i, j := 0, 1; i < len(strs); {
		if j < len(strs) && strs[j] == strs[i] && i != j { //если есть совпадение - удаляем элемент из сайса
			strs = append(strs[0:j], strs[j+1:]...)
		} else if j < len(strs) { //сравниваем следующий элемент и оставляем проверяемый пока j не дойдет до конца
			j++
		}
		if j == len(strs) { //когда дошли до конца возвращаем к началу переменную которая обходит весь массив и проверяем следующий
			i++
			j = i + 1
		}
	}

	return strs
}
