package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Комбинации круглых скобок")
	fmt.Println("--------------------------")
	fmt.Println("Введите количество пар скобок:")

	var parentheses int
	fmt.Scan(&parentheses)
	parentheses *= 2

	answer := ""
	bin := ""
	count := 0

	for a := parentheses; a != 0; a-- {
		bin += "1"
	}

	binary, _ := strconv.ParseInt(bin, 2, 0)

	for len(bin) == parentheses {
		b := bin
	chek:
		for len(b) != 0 {

			if b[:1] == "1" {
				count++
			} else if b[:1] == "0" {
				count--
			}
			if count < 0 {
				break chek
			}
			b = b[1:]
		}

		if count == 0 {
			answer += " " + bin
		} else {
			count = 0
		}
		binary--
		bin = strconv.FormatInt(binary, 2)
	}

	bin = strings.Trim(answer, " ")
	answer = ""

	for len(bin) != 0 {
		if bin[:1] == "1" {
			answer += "("
		} else if bin[:1] == "0" {
			answer += ")"
		} else if bin[:1] == " " {
			answer += "\",\""
		}
		bin = bin[1:]
	}
	answer = "[\"" + answer + "\"]"
	fmt.Println(answer)
}
