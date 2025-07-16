package main

import (
	"fmt"
	"log"
	"time"
)

/*
	defer keyword is always used with connecting, disconnecting or
	lock and unlock to ensure that resources are released in all cases
*/

func double(x int) int {
	return x + x
}

func deferredDouble(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func main() {
	fmt.Println(double(3))
	fmt.Println(deferredDouble(3))
	fmt.Println(triple(3))
}

// func ReadFile(filename string) ([]byte, error) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	return ReadAll(f)
// }

func bigSlowOperation() {
	//	The defer statement can also be used to pair ‘‘on entry’’ and ‘‘on exit’’ actions
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate show operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
