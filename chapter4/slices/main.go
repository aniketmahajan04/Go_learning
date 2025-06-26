package main

import "fmt"

func main() {
	months := [...]string{
		1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "Jully", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December",
	}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	// endlessSummer := summer[:5]
	// fmt.Println(endlessSummer)

	// fmt.Println(summer[:20]) // out of range
	//
	// for _, s := range summer {
	// 	for _, q := range Q2 {
	// 		if s == q {
	// 			fmt.Printf("%s appearce in both\n", s)
	// 		}
	// 	}
	// }
	a := [...]int{1, 2, 3, 4, 5}
	reverse(a[:]) // a[:] converting array a into slice of same data
	fmt.Println(a)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
