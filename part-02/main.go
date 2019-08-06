package main

import "fmt"

// Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax. 
const w  =  3

var x = 2

func main() {
	var y int64

//  Inside a function, the := short assignment statement can be used 
//  in place of a var declaration with implicit type.
	z := 1	
	z = 3
	x = 3
	//w = 3 // fail

	name := "John"

	store_name := fmt.Sprintf("%s", name)
	fmt.Printf("%v\n", store_name)

	fmt.Printf("w: %d, type: %T\n", w, w)
	fmt.Printf("x: %d, type: %T\n", x, x)
	fmt.Printf("y: %d, type: %T\n", y, y)
	fmt.Printf("z: %d, type: %T\n", z, z)
}
