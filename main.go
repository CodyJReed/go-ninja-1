package main

import "fmt"

var w int = 42
var y string = "James Bond"
var z bool = true
type ownType int
var x ownType

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}