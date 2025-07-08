package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int) // using make function to create a map
	ages["alice"] = 30           // adding a key value pair to the map
	ages["bob"] = 25

	fmt.Println(ages)
	// another way
	ages2 := map[string]int{
		"alice":  30,
		"aniket": 21,
	}
	ages2["bob"] = ages2["bob"] + 1
	ages2["bob"] += 1 // shorthand
	ages2["bob"]++
	// _ = &ages2["bob"] // compile error
	for name, age := range ages2 {
		fmt.Printf("%s:\t%d,\n", name, age)
	}
	var names []string
	for name := range ages2 {
		names = append(names, name)
	}
	sort.Strings(names) // sorting the names
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages2[name])
	}
	fmt.Println(ages2)
	ages2["aniket"] = ages2["aniket"] + 1 // manually incrementing a value

	fmt.Println(ages2["alice"]) // 30
	delete(ages2, "alice")      // deleting a key from the map
	fmt.Println(ages2)

	var aniket map[string]int
	fmt.Println(aniket == nil)    // true, aniket is nil
	fmt.Println(len(aniket) == 0) // true, length is 0
	// aniket["aniket"] = 21 // this will panic because aniket is nil

	// age, ok := ages["bob"]
	// if !ok {
	/* "bob" is not a key in this map; age == 0 */
	// }
	// this is assgining a value and checking the value
	// if age, ok := ages["bob"]; !ok { /* ... */}
	xMap := map[string]int{"a": 1}
	yMap := map[string]int{"a": 1}
	fmt.Println(equal(xMap, yMap))                                     // true
	fmt.Println(equal(map[string]int{"a": 1}, map[string]int{"b": 1})) // false
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
