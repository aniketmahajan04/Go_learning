package main

import "fmt"

// var x int = 12

func main() {
	// var x int = 12
	// p := &x
	// fmt.Println("Value of x before incre: ", x)
	// incre(x)
	// fmt.Println("Value of x after incre: ", x)
	// increPointer(&x)
	// fmt.Println("Value of x after increPointer: ", x)
	// fmt.Println("Value of x pointer: ", *p)
	// *p = 20
	// fmt.Println("Value of x after pointer change: ", x)

    // Tuple assignment meant assigning multi variable at once
	x, y := 10, 20
	fmt.Println("Before swap: x = ", x, ", y = ", y)
	x, y = y, x // Swapping values
	fmt.Println("After swap: x = ", x, ", y = ", y)

	// This will work when a function have multiple return one is responce and second is error
	// res, err = response
	// if err != nil { panic(err)

	//Assingabality
	medals := []string{"Gold", "Silver", "Bronze"}
	//or
	medals[0] = "Gold"
	medals[1] = "Silver"
	medals[2] = "Bronze"
}

func incre(x int) int {
	x++
	return x
}

func increPointer(x *int) *int {
    *x++
	return x
}

