package main

import (
	"fmt"
	"os"
)

// Рекомендация. Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные.
// А что мы так проверяем?

func main() {

	fmt.Println("Задание 3. Уровни доступа")

	file, _ := os.Create("note.txt")
	os.Chmod("note.txt", 0444)
	defer file.Close()

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
