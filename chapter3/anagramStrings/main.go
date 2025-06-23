package main

import "fmt"

func main() {
	fmt.Println(isAnagram("silent", "listen"))
}

func isAnagram(str1, str2 string) bool {
	string1 := sorted(str1)
	string2 := sorted(str2)
	return string1 == string2
}

func sorted(str1 string) string {
	runes := []rune(str1)
	n := len(runes)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}
