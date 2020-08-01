


// 1 Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.

package main

// import "fmt"
//
// func main() {
// 	var pCar map[string]string
//
//   short := map[string]string{
//     "vender": "Audi",
//     "model":  "TT",
//     "yearManufacture":  "2005",
//     "trunkVolume":  "45",
//     "bodyVolume":  "700",
//     "engineRunning":  "true",
//     "windowsOpen":  "false",
//     "fullTrunkVolume":  "30",
//
//
//
//   }
//
// profile := make(map[string]string, 10)



import (
    "encoding/json"
    "fmt"
    "log"
)


func main() {


type Auto struct {
    Vender string
    Model string
    YearManufacture int
    TrunkVolume int
    BodyVolume int
    EngineRunning string
    WindowsOpen string
    FullTrunkVolume int

}
//Грузовик
truke := "{\"Vender\": \"КамАЗ\",\"Model\": \"65201\",\"YearManufacture\": 2014,\"TrunkVolume\": 45,\"BodyVolume\": 95000,\"EngineRunning\": \"false\",\"WindowsOpen\": \"false\",\"FullTrunkVolume\": 100}"


//легковушка
//pCar := "{\"Vender\": \"Audi\",\"Model\": \"TT\",\"YearManufacture\": 2005,\"TrunkVolume\": 45,\"BodyVolume\": 700,\"EngineRunning\": \"true\",\"WindowsOpen\": \"false\",\"FullTrunkVolume\": 30}"



//Грузовик
    bT := []byte(truke)
    mT := Auto{}

    err := json.Unmarshal(bT, &mT)
      if err != nil {
          log.Println(err)
          return
      }

      fmt.Println(mT)



    //   //легковушка
    // bP := []byte(pCar)
    // mP := Auto{}
    //
    // err := json.Unmarshal(bP, &mP) //Ругается  = no new variables on left side of :=
    //         if err != nil {
    //             log.Println(err)
    //             return
    //         }
    //
    //         fmt.Println(mP)


}
