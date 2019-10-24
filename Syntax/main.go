package main

import (
	"fmt"
	"math"
)

func main() {
	hello()
}

//Hello
func hello() {
	fmt.Println("Hello, 世界")
	//valuesOld()
	//valuesNew()
	//conditional()
	//constants()
	//loops()
	//switches()
	//arrays()
	//slices()
	//maps()
	//ranges()
	//functions()
	//functionMultipleReturn()
	//variadics()
	//structs()
	interfaces()
}

//variable syntax
func valuesOld() {
	var x int = 5
	var y int = 7
	var sum int = x + y
	fmt.Println(sum)
}

func valuesNew() {
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)
}

func conditional() {
	x := 5
	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {

	} else {

	}
}

func constants() {
	const pi float64 = 3.1415926
	fmt.Println(pi)

	const c int = 1000

	fmt.Println(c)

	//cant't reassign constant value
	//c = 50

	var d int = 100
	fmt.Println(d)
}
func loops() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j <= 5; j++ {

	}
	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}
	for {
		fmt.Println("If it were not for the break it would keep printing forever")
		break
	}
}

func switches() {
	i := 4

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	j := 5

	switch j {
	case 1, 2:
		fmt.Println("one or two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Not one, two, or three")
	}

	switch {
	case j == 5:
		fmt.Println(j, "is equal to 5")
	case j < 5:
		fmt.Println(j, "is less than 5")
	case j > 5:
		fmt.Println(j, "is greater than 5")
	}
}

func arrays() {
	var intArr [5]int
	fmt.Println(intArr)

	var boolArr [10]bool
	fmt.Println(boolArr)

	var strArr [20]string
	fmt.Println(strArr)

	intArr[0] = 5
	intArr[1] = 10

	fmt.Println(intArr)
	fmt.Println(intArr[0])

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	fmt.Println(len(a))
	fmt.Println(len(boolArr))
}

func slices() {
	//make is slice-keyword
	s := make([]int, 3)

	//Are similar to arrays when assigning value
	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println(len(s))

	//append function is unique to slices.
	s = append(s, 4)
	fmt.Println(s)
	s = append(s, 5, 6, 7)
	fmt.Println()

	// slice syntax: prints element 1 and 2
	fmt.Println(s[1:3])
	// slice syntax: prints all up to 2nd element
	fmt.Println(s[:3])

	//concise slice definition
	t := []int{100, 200, 300}
	fmt.Println(t)

	//x := s
	//fmt.Println(x)
	//x[0] = 500
	//fmt.Println(x)
	//fmt.Println(s)

	// Use copy to prevent from changing both x and s
	x := make([]int, len(s))
	copy(x, s)

	x[0] = 500
	fmt.Println(x)
	fmt.Println(s)
}

func maps() {
	// Maps are similar to Python dictionary (i.e. hashtable, etc.)

	//specify key/value pair
	m := make(map[string]int)
	m["a"] = 0
	m["b"] = 1

	fmt.Println(m)
	//print the value of a specific key
	fmt.Println(m["a"])

	//len() function is overloaded to with maps:
	fmt.Println(len(m))

	//remove key/pair from map. (which can be done with delete keyword).
	delete(m, "a")
	fmt.Println(m)

	m2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(m2)
}

func ranges() {
	//range is similar to range in Python.
	strArr := []string{"a", "b", "c", "d"}
	// first var: index, second var: character
	for i, c := range strArr {
		fmt.Println("Index:", i)
		fmt.Println("Character", c)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println("key:", k, "val:", v)
	}
}

func func1(a int, b int) int {
	return a + b
}
func func2(a, b, c int) int {
	return a + b + c
}

func functions() {
	ans := func1(1, 1)
	fmt.Println(ans)

	ans2 := func2(1, 2, 3)
	fmt.Println(ans2)
}

func funcMultiple(a, b int) (int, int) {
	return a + b, a - b
}

func functionMultipleReturn() {
	addRes, subRes := funcMultiple(1, 1)
	fmt.Println("addRes: ", addRes, "subRes: ", subRes)
}

func variadicmult(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

func variadics() {
	fmt.Println(variadicmult(1, 2, 3, 4))
}

type employee struct {
	firstName string
	lastName  string
	id        int
}

func structs() {
	emp := employee{"Homer", "Simpson", 1}
	//when reassigning object (which is a struct here) use pointer
	empPtr := &emp
	fmt.Println(empPtr.firstName)
}

func methods() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())

	rPtr := &r
	fmt.Println("perimeter. ", rPtr.perimeter())

}

type geometry interface {
	area() float64
	perimeter() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func measures(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
func interfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measures(r)
	measures(c)
}
