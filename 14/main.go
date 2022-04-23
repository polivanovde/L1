//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}
package main

import "fmt"

func main() {
	checkType(11)
	checkType("со словом")
	checkType(true)
	c := make(chan int)
	checkType(c)
}

func checkType(i interface{}) {
	//просто определяем тип и сохраняем d gthtvtyye.
	typeI := fmt.Sprintf("%T", i)
	fmt.Println(typeI)

	//или используем свитч и делаем определенные действия с определенным типом
	switch v := i.(type) {
	case int:
		fmt.Println("Это число, его можно например умножить:", v*2)
	case string:
		fmt.Println("Это строка, ее сконкатенируем " + v)
	case bool:
		fmt.Println("Это boolean и он", v)
	default:
		fmt.Printf("Этот тип в свитче не определен, это тип - %T!\n", v)
	}
}
