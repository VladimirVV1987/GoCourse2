// 2 Перепишите программу-конвейер, ограничив количество передаваемых для обработки значений и обеспечив корректное завершение всех горутин.

//увы, сколько я не боролся с правильным завершением, так и не получилось.

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	wg.Add(1)
	go func() {
		for x := 0; x <= 10; x++ { //ограничим 10-ю
			naturals <- x
		}

	}()
	defer wg.Done()

	// возведение в квадрат
	wg.Add(1)
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	defer wg.Done()

	// печать
	for {
		fmt.Println(<-squares)
	}

	wg.Wait()
}
