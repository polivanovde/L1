//Реализовать собственную функцию sleep

package main

import (
    "fmt"
    "strings"
)

func checkRepeat(st string) bool {
    st = strings.ToLower(st) //приводим строку к нижнему регистру

    for _,elem := range st { //перебираем каждый символ
        if strings.Count(st, string(elem)) > 1 { //проверяем наличие символа в строке
            return false //если символов больше одного то возращаем отрицательный результат
        }
    }
    return true
}

func main() {
    var st string
    fmt.Scan(&st)

    fmt.Println(checkRepeat(st))
}
