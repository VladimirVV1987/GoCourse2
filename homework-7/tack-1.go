// 1 Уберите из первого примера (Фибоначчи и спиннер) функцию, вычисляющую числа Фибоначчи. Как теперь заставить спиннер вращаться в течение некоторого времени? 10 секунд?

package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)

	fmt.Println("Ждем 10 сек")
	time.Sleep(10 * time.Second)

	fmt.Println("Прошло 10 сек")
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
