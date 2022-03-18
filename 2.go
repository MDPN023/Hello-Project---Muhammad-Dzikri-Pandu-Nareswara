package main

import "fmt"

func main() {
	var x int
	var a bool
	fmt.Print("Silahkan masukkan bilangan yang ingin dimasukkan : ")
	fmt.Scan(&x)
	if x%2 == 1 {
		a = true
		fmt.Print(a)
	} else {
		fmt.Print(a)
	}
}
