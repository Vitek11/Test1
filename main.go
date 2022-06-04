package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello , my dear friend!")
	var s, s1, s2 string
	s1 = "How"
	s2 = " are you?"
	s = s1 + s2
	fmt.Println(s)

	var x int32 = 36
	var y int32 = 11
	var z1, z2, z3, z4 int32
	z1 = x + y
	z2 = x * y
	z3 = x / y
	z4 = x % y

	fmt.Println("First number:", x)
	fmt.Println("Second number:", y)
	fmt.Println(x, "+", y, "=", z1)
	fmt.Println(x, "*", y, "=", z2)
	fmt.Println(x, "/", y, "=", z3, " (", z4, ")")

	var z5 float32
	z5 = float32(x) / float32(y)
	fmt.Println(x, "/", y, "=", z5)

	var z6, z7 float64
	z6 = math.Sqrt(float64(x))
	z7 = math.Sqrt(float64(y))

	fmt.Println("sqrt (", x, ") = ", z6)
	fmt.Println("sqrt (", y, ") = ", z7)

}
