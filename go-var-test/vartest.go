package main

import "fmt"

var integer int = 1
var str string = "Toto"
var boolean bool = true

func main() {
	localinteger := 32
	str := "Tata"
	fmt.Printf("Integer : \t\t%v\n", integer)
	fmt.Printf("Local Integer : \t%v\n", localinteger)
	fmt.Printf("String : \t\t%v\n", str)
	fmt.Printf("Boolean : \t\t%v\n", boolean)
}
