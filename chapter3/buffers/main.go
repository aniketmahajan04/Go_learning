package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(intToString(12345))
}

func intToString(values int) string {
	var buf bytes.Buffer
	s := strconv.Itoa(values)
	n := len(s)

	pre := n % 3
	if pre == 0 && n > 3 {
		pre = 3
	}
	buf.WriteString(s[:pre])
	for i := pre; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
