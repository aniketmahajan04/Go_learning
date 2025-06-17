package main

import "fmt"

/*
If the result of an arithmetic operation, whether signe d or unsigned,has more bits than can be
represented in the result type, it is said to overflow. The high-order bits that do not fit are
silently discarded. If the original number is a signed type, the result could be negative if the
leftmost bit is a 1, as in the int8 example here
*/
func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127 -128 1
	fmt.Print("hello")

	// position of bits 7 6 5 4 3 2 1 0
	var x uint8 = 1<<1 | 1<<5 // turning 0 to 1 at postions 1 and 5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	/*
	   	x : 00100010
	    y : 00000110
	    &^  --------
	        00100000 ← only light 5 is ON in `x` but NOT in `y`
	*/

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			/*
				i = 0: x & (1<<0) = 00100010 & 00000001 → 0 → Not printed

				i = 1: 00100010 & 00000010 → 00000010 → YES → Print 1

				...

				i = 5: 00100010 & 00100000 → 00100000 → YES → Print 5
			*/
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) //"bronze", "silver", "gold"
	}
	var apples int32 = 1
	var oranges int16 = 2
	var compote int = apples + oranges // Invalid operation: mismatched types int32 and int16
}
