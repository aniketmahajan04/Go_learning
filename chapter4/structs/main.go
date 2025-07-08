package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID       int
	Name     string
	Address  AddressStruct // Address is a struct, not a string
	DoB      time.Time
	Position struct {
		Title          string
		Level          int
		PromotionDate  time.Time // embeded struct
		RequiredSkills struct {
			PrimaryLanguage   string
			SecondaryLanguage string
			DevOps            struct { // struct inside struct
				Docker     bool
				Kubernetes bool
				CICD       bool
				Ancible    bool
			}
		}
	}
	Salary    int
	ManagerId int
}

// Recursive struct`
type tree struct {
	value       int
	left, right *tree
}

type AddressStruct struct {
	vilage  string
	city    string
	state   string
	pincode int
}

type (
	Point struct{ X, Y int } // first string literal with element exported
	T     struct{ a, b int } // second form of struct literal
)

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

var dilbert Employee

type address struct {
	hostname string
	port     int
}

type Circle struct {
	Center Point
	Radius int
}

type wheel struct {
	Circle
	Spokes int
}

type circle struct {
	Point  // <- field without name
	Radius int
}

func main() {
	// dilbert.Salary -= 5000 // demoted, for writing too few lines of codes
	// position := &dilbert.Position
	// *position = "Senior " + *position // promoted, for outsourcing to Elbonia

	// var aniket Employee
	// fmt.Println("Enter Employee Name:")
	// fmt.Scanln(&aniket.Name)
	// fmt.Println("Enter Employee Address:")
	// fmt.Scanln(&aniket.Address)
	// fmt.Println(aniket)
	// var employeeOfTheMonth *Employee = &dilbert
	// employeeOfTheMonth.Position += " (Proactive team player) "
	// // or
	// (*employeeOfTheMonth).Position += " (Proactive team player) "
	//
	// fmt.Println(EmployeeOfTheMonth(dilbert.ManagerId).Position)
	// id := dilbert.ID
	// EmployeeOfTheMonth(id)
	p := Point{1, 2}
	q := Point{1, 2}
	fmt.Println(p == q)                   // true
	fmt.Println(p.X == q.X && p.Y == q.Y) // true
	// pp := &Point{1, 2}
	ppp := new(Point)
	*ppp = Point{1, 2}
	fmt.Println(p)
	fmt.Println(Scale(Point{1, 2}, 5))
	dilbert.Salary = 200000
	fmt.Println(Bonus(&dilbert, 14))
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	// var w wheel
	// w.Circle.Center.X = 8
	// w.Circle.Center.Y = 8
	// w.Circle.Radius = 5
	// w.Spokes = 20
	// fmt.Println(w)
	// fmt.Println(w.Circle)

	// var W wheel
	// W.X = 8
	// W.Y = 8
	// W.Radius = 5
	// W.Spokes = 20

	a := wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: Trailing comman is necessary
	}
	fmt.Printf("%#v\n", a)
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// func EmployeeOfTheMonth(id int) *Employee { /* ... */ }
