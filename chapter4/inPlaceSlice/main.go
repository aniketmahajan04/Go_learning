package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noneEmpty(data)) // ["one", "three"]
	fmt.Printf("%q\n", data)            // [ "one, "three", "three"]
	// s := []int{5, 6, 7, 8, 9}
	// fmt.Println(remove(s, 2))  // [5 6 8 9]
	// fmt.Println(remove2(s, 2)) // [5 6 8 9]

	// fmt.Println(reverse(&s)) // [5 6 8 9]
	slices := []string{"a", "a", "b", "c", "c", "c", "d", "d"}
	fmt.Println(removeAdjacent(slices))
	utf8Data := []byte("Hello,   \t  \nWorld!\n\n")
	fmt.Printf("%q\n", squashSpaces(utf8Data))
	data2 := []byte("Hello, 世界")
	fmt.Println(modifiedReverese(data2))
}

func modifiedReverese(slice []byte) []byte {
	// Step: 1 Decode runes and record their byte ranges
	var runes [][]byte
	for i := 0; i < len(slice); {
		_, size := utf8.DecodeRune(slice[i:])
		runes = append(runes, slice[i:i+size])
		i += size
	}
	// Step: 2 Copy reversed runed back to the original slice
	pos := 0
	for i := len(runes) - 1; i >= 0; i-- {
		copy(slice[pos:], runes[i])
		pos += len(runes[i])
	}
	return slice
}

func noneEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:]) // shift a[i+1:] left one index
	return slice[:len(slice)-1]  // truncate slice
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] // replace a[i] with the last element
	return slice[:len(slice)-1]    // truncate slice
}

func reverse(slice *[]int) []int {
	n := len(*slice)
	for i := 0; i < n/2; i++ {
		(*slice)[i], (*slice)[n-i-1] = (*slice)[n-i-1], (*slice)[i]
	}
	return *slice
}

func removeAdjacent(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}
	write := 1 // start from 1 because the first element is always kept
	for read := 1; read < len(slice); read++ {
		if slice[read] != slice[read-1] {
			slice[write] = slice[read]
			write++
		}
	}
	return slice[:write]
}

func squashSpaces(data []byte) []byte {
	write := 0
	read := 0
	inSpaceRune := false

	for read < len(data) {
		r, size := utf8.DecodeRune(data[read:])
		if unicode.IsSpace(r) {
			if !inSpaceRune {
				data[write] = ' '
				write++
				inSpaceRune = true
			}
		} else {
			n := utf8.EncodeRune(data[write:], r)
			write += n
			inSpaceRune = false
		}
		read += size
	}
	return data[:write]
}
