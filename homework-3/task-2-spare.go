// 2 Инициализировать несколько экземпляров структур. Применить к ним различные действия. Вывести значения свойств экземпляров в консоль.



package main

import "fmt"

func main() {

type position struct {
    X int
    Y int
}

var ThridPoint = position{
    Y: 6,
    X: 3,
}

var resY = ThridPoint.Y

fmt.Println(ThridPoint.X)
fmt.Println(resY)


// еще структура


type Circle struct {
    Point  position
    Radius int
}

// Так будет выглядеть объявление
var FirstCircle = Circle{Point: position{18, 21}, Radius: 43}


fmt.Println(FirstCircle.Point.Y)






}
