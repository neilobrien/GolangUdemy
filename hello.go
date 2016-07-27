package main

import "fmt"

func main() {
	fmt.Println("hello Println")
	fmt.Printf("hello, world")
	fmt.Print("Hellp Print\n")

	// var age float64 = 10.5
	// var isTrue bool = True
	random := true

	var myName string = ("Neil")
	fmt.Println(myName + " OBrien")

	fmt.Printf("%T \n", random)

	Array := [5]int{1, 2, 3, 4, 5}

	for _, value := range Array {
		fmt.Println(value)
	}

	fmt.Println(factorial(4))

	x := 0
	changeX(&x)
	fmt.Println("x =", x)
}

func changeX(x *int) {
	*x = 2
}

func factorial(num int) int {

	if num == 0 {
		return 1
	}
	return num * factorial(num-1)

}
