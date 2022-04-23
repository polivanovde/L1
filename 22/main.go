//Разработать программу, которая перемножает, делит, складывает, вычитает
//две числовых переменных a,b, значение которых > 2^20

package main

import (
	"fmt"
	"math/big"
)

func main() {
    var (
	    a string
        b string
        sep string
        bigA = new(big.Int)
        bigB = new(big.Int)
    )

    //порядок ввода: ввод 1 числа, ввод операции, ввод 2 числа
    fmt.Scan(&a, &sep, &b)
    bigA.SetString(a, 10)
    bigB.SetString(b, 10)

    switch sep { //
        case "+":
            bigA.Add(bigA, bigB)
            fmt.Println(bigA)
        case "-":
            bigA.Sub(bigA, bigB)
            fmt.Println(bigA)
        case "/":
            zero := big.NewInt(0)
            if bigB.Cmp(zero) == 0 {
                fmt.Println("деление на 0 =(")
                return
            }
            bigZ := new(big.Int)
            bigZ.Div(bigA, bigB)
            fmt.Println(bigZ)
        case "*":
            bigZ := new(big.Int)
            bigZ.Mul(bigA, bigB)
            fmt.Println(bigZ)
        default:
            fmt.Println("неизвестная арифметическая операция")
    }
}
