package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		to convert an integer to a string, one option is to use fmt.Sprintf;
		another is to use the function strconv.Itoa(intger to ASCII)
	*/
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2)) // 1111011

	s := fmt.Sprintf("x=%b", x) // Sprintf should be called
	fmt.Println(s)

	a, err := strconv.Atoi("123")
	if err != nil {
		return
	}
	fmt.Printf(
		"a=%d\n",
		a,
	) /*
		a=123% (%) is shown by vscode to tell the end of output line its not fault of my code
		\n adding newline this can fix that
	*/
	b, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		return
	}

	fmt.Printf(
		"b=%d\n",
		b,
	)
}
