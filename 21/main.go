//Реализовать паттерн «адаптер» на любом примере

package main
import (
    "fmt"
)

//объявляем интерфейс с которым должна работать адаптируемая система
type OldInterface interface {
    InsertToDatabase(Data interface{}) (bool, error)
}

//структура которую адаптируем - содержит необходимую информацию и по которой строится основная функция
type AddCustomInfoToMysql struct {
    DbName string
}

//основная функция - инсерт (просто печать для примера)
func (pA *AddCustomInfoToMysql) InsertToDatabase(info interface{}) (bool, error) {
    switch info.(type) {
    case string:
        fmt.Println("add ", info.(string), " to ", pA.DbName, " successful!")
    }
    return true, nil
}

//новый интерфейс объявляет функцию с которой будем работать - новый интерфейс ожидает реализации своей функции,
//отличной от базового (адаптируемого)
//(на входе - строку которую распечатает функция базового интерфейса)
type NewInterface interface {
    SaveData(Data interface{}) (bool, error)
}

//адаптер реализует целевой интерфейс
type Adapter struct {
    OldInterface
}

//реализуем новую функцию - новая функция выполняет (и возвращает соответствующие данные)
//базовый функционал - инсерт в БД
func (pA *Adapter) SaveData(Data interface{}) (bool, error) {
    fmt.Println("In Adapter")
    return pA.InsertToDatabase(Data)
}

func main() {
    var inNew NewInterface //объявляем новый интерфейс
    inNew = &Adapter{OldInterface: &AddCustomInfoToMysql{DbName: "mysql"}} //добавляем в новый интерфейс адаптер,
                                                                           //который содержит адаптируемую структуру с
                                                                           //нужными элементами - имя ДБ и функционал инсерта

    inNew.SaveData("hello world") //вызываем новую функцию, которая выполнит инсерт (печать) из базового интерфейса
}