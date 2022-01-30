package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Задание 1. Работа с файлами")

	var (
		notes   string
		timeTxt string
		num     int64
	)

	t := time.Now()

	file, _ := os.Create("note.txt")
	defer file.Close()

	for {
		num++
		fmt.Scan(&notes)
		if notes == "exit" {
			break
		} else {
			timeTxt = t.String()
			timeTxt = t.Format("2006-01-02 15:04:05")
			file.WriteString(strconv.FormatInt(num, 10) + ". " + timeTxt + " " + notes + "\n")
		}
	}
}
