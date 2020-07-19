// Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу. Используйте тип данных float64 и функции из пакета math.

package main

import (
  "fmt"
  "math"
)

func main(){
  var katet1 float64 = 4
  var katet2 float64 = 7
  var square float64 = katet1 * katet2 * 0.5
  var hypotenuse float64 = math.Sqrt(math.Pow(katet1, 2) + math.Pow(katet2, 2))

  var perimeter float64 = katet1 + katet2 + hypotenuse

square = (math.Ceil(square*100)/100)
perimeter = (math.Ceil(perimeter*100)/100)
hypotenuse = (math.Ceil(hypotenuse*100)/100)



fmt.Println("Площадь: ", square)
fmt.Println("Периметр: ", perimeter)
fmt.Println("Гипотенуза: ", hypotenuse)

}
