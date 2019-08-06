package main

import "fmt"

func main() {

	array := [...]string{"hello", "world"}
	//	array = append(array, "!") // cannot compile

	slice := []string{"hello", "world"}
	slice = append(slice, "!")

	fmt.Printf("array is: %T , slice is %T\n", array, slice)

	// %v : value of the variable/string/array
	fmt.Printf("i: %v \n", slice)
	fmt.Printf("i: %v \n", slice[1:])
	
/*
In simple words _ is an ignored value. Go doesn't permit unused local variables and throws a compilation error if you try to do so. So you trick compiler saying "please ignore this value" by placing a _

For example:

// Here range of os.Args returns index, value
for index, arg := range os.Args[1:] {
    fmt.Println(arg)
}

Here Go throws compilation error because you are not using index (un used local variable) at all in the loop. Now you should make index to _ to say compiler to ignore it.

for _, arg := range os.Args[1:] {
    fmt.Println(arg)
}


*/


/*
 The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index. 
*/
	for _, v := range slice {
		fmt.Printf("%s, ", v)
	}

	for _, v := range array {
		fmt.Printf("%s, ", v)
	}
}

