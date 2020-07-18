package main

import (
	"fmt"
	"math"
)

const usd float64 = 70.13

func main(){
	var rub float64
	fmt.Println("Курс usd: ", usd, " руб.")
	fmt.Println("Введите сумму в рублях: ")
	fmt.Scan(&rub)
	var result float64 = rub / usd
	result = math.Ceil(result*100)/100

	fmt.Println("Вы получите: ", result, " usd")

}
