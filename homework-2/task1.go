//1. Написать функцию, которая определяет, четное ли число.

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Введите число, а мы определим четное ли оно или нет: ")
	var numberS int
	fmt.Scan(&numberS)

	if numberS%2 == 0 {
		fmt.Println(numberS, "четное")
	} else {
		fmt.Println(numberS, "нечетное")
	}


}
