package main

import (
	"fmt"
	"math"
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota
	FlagBroadcast          // support broadcast access capability
	FlagLoopback           // is a loopback interface
	FlagPointToPoint       // belongs to a point-to-point link
	FlagMulticast          // supports multicast access capability
)

const (
	_         = iota
	KB uint64 = 1 << (10 * iota) // 1 << (10 * 1) = 1024
	MB                           // 1 << (10 * 2)
	GB                           // 1 << (10 * 3)
	TB
	EB
	ZB
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	// var v Flags = FlagMulticast | FlagUp
	// fmt.Printf("%b %t\n", v, IsUp(v)) // 10001 true
	// TurnDown(&v)
	// fmt.Printf("%b %t\n", v, IsUp(v)) // 10000 false
	// SetBroadcast(&v)
	// fmt.Printf("%b %t\n", v, IsUp(v))   // 10010 false
	// fmt.Printf("%b %t\n", v, IsCast(v)) // 10010 true

	var Pi64 float64 = math.Pi
	var x float32 = float32(Pi64)
	var y float64 = Pi64
	// var z complex128 = complex128(Pi64) // cannot convert float into complex
	fmt.Println(x)
	fmt.Println(y)
}
