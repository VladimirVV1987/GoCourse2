// 2 Инициализировать несколько экземпляров структур. Применить к ним различные действия. Вывести значения свойств экземпляров в консоль.
package main

import "fmt"

func main() {

  type Auto struct {
      Vender string
      Model string
      YearManufacture int
      TrunkVolume int
      BodyVolume int
      EngineRunning bool
      WindowsOpen bool
      FullTrunkVolume int

  }

var PCar = Auto{
    Vender: "BMW",
    YearManufacture: 2008,
    TrunkVolume: 27,
}

var resY = PCar.Vender

fmt.Println("Первый вывод")
fmt.Println("Год выпуска: ", PCar.YearManufacture)
fmt.Println("Производитель: ", resY)


// еще структура


type Truke struct {
    AutoCar  Auto
    RadiusDisc int
}

// Так будет выглядеть объявление
var newTruke = Truke{AutoCar: Auto{Vender: "Volvo", FullTrunkVolume: 12}, RadiusDisc: 43}

fmt.Println("\nПосле объявления")
fmt.Println("Производитель: ", newTruke.AutoCar.Vender)
fmt.Println("Радиус дисков: ", newTruke.RadiusDisc)

fmt.Println("Все: ", newTruke.AutoCar)






}
