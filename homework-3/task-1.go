// 1 Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.

package main

import (

    "fmt"

)




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


func main() {

//легковушка
  PCar := Auto{
    Vender: "Audi",
    Model: "TT",
    YearManufacture: 2005,
    TrunkVolume: 40,
    BodyVolume: 700,
    EngineRunning: false,
    WindowsOpen: true,
    FullTrunkVolume: 30,
  }


//грузовик
  Truke := Auto{
    Vender: "КамАЗ",
    Model: "65201",
    YearManufacture: 2014,
    TrunkVolume: 45,
    BodyVolume: 95000,
    EngineRunning: false,
    WindowsOpen: false,
    FullTrunkVolume: 100,
  }


fmt.Println(PCar.Model)
fmt.Printf("Производитель: %v\nЗапущен ли двигатель: %v\n", Truke.Vender, Truk.EngineRunning)


fmt.Println("Все ", PCar)

}
