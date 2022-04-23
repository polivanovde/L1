//Разработать программу, которая переворачивает подаваемую
//на вход строку (например: «главрыба — абырвалг»). Символы могут быть unicode
package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)                                    //переводим символы в срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { //обозначаем начало и конец среза и уменьшаем к центру
		runes[i], runes[j] = runes[j], runes[i] //обоюдно меняем местами сиволы
	}
	return string(runes)
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(Reverse(input))
}
