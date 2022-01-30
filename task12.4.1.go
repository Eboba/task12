package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Задание 4. Пакет ioutil")

	var (
		notes   string
		timeTxt string
		num     int64
	)
	var b bytes.Buffer
	file, _ := os.Create("note.txt")
	t := time.Now()
	defer file.Close()

	for {
		num++
		fmt.Scan(&notes)
		if notes == "exit" {
			break
		} else {
			timeTxt = t.String()
			timeTxt = t.Format("2006-01-02 15:04:05")
			b.WriteString(strconv.FormatInt(num, 10) + ". " + timeTxt + " " + notes + "\n")
		}
	}
	ioutil.WriteFile("note.txt", b.Bytes(), 0666)

	defer file.Close()
}
