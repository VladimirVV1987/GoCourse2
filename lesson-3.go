package main

import "fmt"

func main() {
	var years int = 5
	var deposit int
	var procent int
  var summProcent int

	fmt.Printf("Введите сумму депозита: ")
	fmt.Scan(&deposit)
	fmt.Printf("Введите % в год: ")
	fmt.Scan(&procent)
	var depositStart int = deposit

	for i := 1; i <= years; i++ {
		deposit += deposit / 100 * procent
		summProcent = deposit - depositStart
		fmt.Println("За ", int(i), "год Вы получите: ", deposit, "руб., сумма процентов = ", summProcent, "руб.")
	}

}

// num0Elem, err:= fmt.Scanln(&deposit, &procent)
// if err != nill{
//   fmt.Println(err)
//
// }
// fmt.Println(num0Elem,deposit,procent)
