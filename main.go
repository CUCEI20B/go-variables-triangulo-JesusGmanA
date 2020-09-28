package main

import "fmt"

func main() {
	var base, altura float64
	fmt.Scan(&base)
	fmt.Scan(&altura)
	area := base * altura / 2
	fmt.Println(area)
}
