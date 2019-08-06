package main
/*

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p

*/
import "fmt"

func main() {

	i := 0
	//Inifite Loop
	for 
	{
		if i > 9 {
			fmt.Printf("i: %d, type: %T\n", i, i);
			break
		}
		i++
	}

	// Curly Braces Should Follow on the same line for loop
	for i :=0 ; i < 20; i++ { 
		fmt.Printf("%d ", i) 
	}
	
	fmt.Printf("\ni: %d \n", i)

	for  ; i < 20; i++ { 
		fmt.Printf("%d ", i) 
	}
	
	fmt.Printf("\ni: %d \n", i)
}
