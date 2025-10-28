package main

import (
	"fmt"
	"strings"
)

func enterNumber(n *int) {
	fmt.Print("Enter number: ")
	fmt.Scan(n)
}

func main() {
	var n, flag int
	var ans string
	enterNumber(&n)
	for n > 12307 {
		fmt.Println("Your number is over 12307")
		fmt.Print("Are you sure? (Enter y/n): ")
		fmt.Scan(&ans)
		if ans == "n" {
			enterNumber(&n)
		} else if ans == "y" {
			break
		}

	}
	for n < 12307 {
		if n < 0 {
			n = n * (-1)
		} else if n%7 == 0 {
			n = n * 39
		} else if n%9 == 0 {
			n = n*13 + 1
			continue
		} else {
			n = (n + 2) * 3
		}

		if n%13 == 0 && n%9 == 0 {
			flag = 1
			break
		} else {
			n++
		}
	}
	if flag == 1 {
		fmt.Println("\nservice error")
	} else {
		fmt.Println(strings.Repeat("=", 60))
		fmt.Printf("The result number is %d!", n)
	}
}
