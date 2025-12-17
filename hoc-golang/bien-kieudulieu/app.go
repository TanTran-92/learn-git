package main

import (
	"bufio"
	"fmt"
	"os"
)

var address = "Ho Chi Minh City"

func main() {
	// fmt.Println("hello tantran - hoc golang")
	// randomuser()
	// var fullName  = "Tran Phuoc Tan"
	// fullName = "Shin tran"
	// fmt.Println(fullName)
	// var age int
	// age = 30
	// fmt.Println(age)

	// phone := "0903779755"
	// fmt.Println(phone)

	// toan := 5
	// tiengviet := 5
	// tunhien := 5
	// fmt.Println(toan)
	// fmt.Println(tiengviet)
	// fmt.Println(tunhien)
	// toan, tiengviet, tunhien := 5, 6, 7
	// fmt.Println(tiengviet)
	// fmt.Println(toan)
	// fmt.Println(tunhien)

	// var hoten string
	// fmt.Print("Vui long nhap ho ten: ")
	// fmt.Scan(&hoten)
	// fmt.Println("Ho ten: ", hoten)
	var hoten string
	fmt.Print("Vui long nhap ho ten: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		hoten = scanner.Text()
	}
	fmt.Println("Ho ten: ", hoten)
}
