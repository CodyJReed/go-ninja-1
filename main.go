package main

import "fmt"

var w int = 42

var z bool = true
type ownType int
var x ownType
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println("This is the value of y:", y)
}