package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Задание 4. Пакет ioutil")

	file, err := os.Open("note.txt")
	if err != nil {
		fmt.Println("Файл отсутствует", err)
	}
	defer file.Close()

	a, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))
}
