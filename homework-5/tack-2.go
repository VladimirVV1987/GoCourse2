//2 Что бы вы изменили в программе чтения из файла? Приведите исправленный вариант, обоснуйте свое решение в комментарии.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil" // для не больших файлов, до 1 мб
	"log"
	"os"
)

func main() {
	file, err := os.Open("tack-3.csv")
	if err != nil {
		log.Fatal(err) //Мои правки
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err) //Мои правки
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err) //Мои правки
	}

	fmt.Println(string(bs))

	//Мои правки
	fmt.Println("==== Мои правки ====")

	// для не больших файлов, до 1 мб без лишних движений
	fileContent, err := ioutil.ReadFile("tack-3.csv")
	if err != nil {
		log.Fatal(err) //Мои правки
	}
	fmt.Println(string(fileContent))

	// для больших файлов

	fmt.Println("==== для больших файлов, построчный вывод ====")
	//код ниже не отработает, его надо ставить выше, но после его выполнения он выходит (exit status 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
