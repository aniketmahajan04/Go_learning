package main

import (
	"os"
)

/*
	defer keyword is always used with connecting, disconnecting or
	lock and unlock to ensure that resources are released in all cases
*/

func main() {

}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadAll(f)
}
