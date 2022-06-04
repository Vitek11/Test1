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

	var x, y int32
	var z1, z2 float64

	x = 25
	y = 42

	z1 = math.Sqrt(float64(x))
	z2 = math.Sqrt(float64(y))

	fmt.Println(z1, z2)
}
