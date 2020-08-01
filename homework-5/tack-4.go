// 4  * Напишите утилиту для копирования файлов, используя пакет flag.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	oldP := flag.String("oldpath", "oldnameF.txt", "Какой файл копируем, указываем полный путь до фаайла")
	newP := flag.String("newpath", "newnameF.txt", "Новое имя файла")

	flag.Parse()

	fmt.Println("oldpath:", *oldP)
	fmt.Println("newpath:", *newP)

//увы, но ругается на стринг
	err := os.Rename(oldP, newP)
	if err != nil {
		log.Fatal(err)
	}

}
