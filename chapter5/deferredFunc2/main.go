package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
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

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, ut prefer error from Copy, if any.

	// converting this in defer function to close() open file
	// if closeErr := f.Close(); err == nil {
	// 	err = closeErr
	// }

	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}
