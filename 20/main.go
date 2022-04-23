//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow»
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//функция разворота из пред задания но по начальному срезу
func Reverse(m []string) []string {
	fmt.Println(m)
	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}
	return m
}

func main() {
	scan := bufio.NewScanner(os.Stdin) //сканируем ввод вместе с пробелами
	scan.Scan()
	input := scan.Text()
	words := strings.Split(input, " ") //делаем из ввода срез со строками

	fmt.Println(strings.Join(Reverse(words), " ")) //соединяем в строку и печатаем
}
