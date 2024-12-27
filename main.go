package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}
func Sub(a int, b int) int {
	return a - b
}
func DoMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	x := DoMath(82364, 65347, Add)
	fmt.Println(x)
	y := DoMath(82364, 65347, Sub)
	fmt.Println(y)

	fmt.Printf("%T\n", Add)
	fmt.Printf("%T\n", Sub)
	fmt.Printf("%T\n", DoMath)
}
