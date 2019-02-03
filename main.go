package main

import "fmt"

func init() {
	fmt.Println("Init!")
}

func main() {
	bazz()
	fmt.Println("Hello World?")
}

func bazz() {
	fmt.Println("bazz")
}