//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	a string
	b string
}

type Action struct {
	Human //анонимное поле структ. Human <- тут условно мы и наследуем методы и поля
	c     int
}

func main() {
	//задана сразу
	// person := Action{
	//     c: 70,
	//     Human: Human{
	//         a: "White male",
	//         b: "133",
	//     },
	// }

	//или пустая
	person := new(Action)
	person.a = "Black male"
	person.b = "by day"
	person.c = 70

	person.HumanType()
	fmt.Printf("can walk %d km %s\n", person.c, person.b)
}

func (h *Human) HumanType() {
	fmt.Printf("Human is %s ", h.a)
}
