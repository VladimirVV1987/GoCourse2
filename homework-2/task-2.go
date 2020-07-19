//2. Написать функцию, которая определяет, делится ли число без остатка на 3.

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Введите число, а мы определим четное ли оно или нет: ")
	var numberS int
	fmt.Scan(&numberS)

	if numberS%3 == 0 {
		fmt.Println(numberS, "можно разделить на 3")
	} else {
		fmt.Println(numberS, "при деление есть остаток")
	}


}
