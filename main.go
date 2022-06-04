package main

import "fmt"

func main() {
	fmt.Println("Hello , my dear friend!")
	var s, s1, s2 string
	s1 = "How"
	s2 = " are you?"
	s = s1 + s2
	fmt.Println(s)

	var x, y int32
	var z1 float32
	x = 10
	y = 3
	z1 = float32(x) / float32(y)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println(x, "/", y, "=", z1)

}
