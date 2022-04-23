//К каким негативным последствиям может привести данный фрагмент кода,
//и как это исправить? Приведите корректный пример реализации
package main

import (
	"fmt"
	"strconv"
)

//var justString string //переменная доступна во всем пакете main но используется только в одной функции, лучше объявить ее там

func someFunc() {
	var justString string
	v := createHugeString(1 << 10) //функция не определена

	//justString = v[:100]
	//получим ошибку slice bounds out of range [:100] with length 4, т.к оператор среза s[i:j]
	//j должен быть меньше либо равен cap: j <= cap(s)

	//заменяем j до конца слайса если его len больше нужного (100)
	if len(v) < 100 {
		justString = v[:]
	} else {
		justString = v[:100]
	}

	fmt.Println(justString)
}

func createHugeString(i int) string {
	return strconv.Itoa(i)
	//либо строку бинарного представления
	//return	fmt.Printf("%b", i)
}

func main() {
	someFunc()
}
