package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Задание 2. Интерфейс io.Reader")

	file, err := os.Open("note.txt")
	if err != nil {
		fmt.Println("Файл отсутствует", err)
	}
	defer file.Close()

	data := make([]byte, 1)
	for {
		n, err := file.Read(data)
		if err != nil { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}
}
