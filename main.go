package main

import "fmt"

func main() {
	// demoDefefLIFO()
	fmt.Println(factorial(5))
	fmt.Println(sum([]int{1,2,3,4}))
}

func demoDefefLIFO() {
		x:=10

		defer fmt.Println(x)
		x = 99
		defer fmt.Println("Defer #1")
		defer fmt.Println("Defer #2")
		defer fmt.Println("Defer #3")
		fmt.Println("Cuerpo de la funcion ejecutandose")
}

func factorial(number int) int{ 

	if number < 0 {
		panic("Numero debe ser mayor o igual a 0")
	}

	// caso base
	if number == 0 {
		return 1
	}

	return number  * factorial(number -1)

}

func sum(nums []int) int {
	
	// caso base : cuando llegue a este punto deberia terminar la recursividad para evitar que se desborde la pila
	if len(nums)  == 0{
		return 0
	}

	return nums[0] + sum(nums[1:])



}