package main

import "fmt"

func main() {
	fmt.Println("Hello , my dear friend!")
	var s, s1, s2 string
	s1 = "How"
	s2 = " are you?"
	s = s1 + s2
	fmt.Println(s)

	var x, y, z1, z2 int32
	x = 25
	y = 7
	z1 = x / y
	z2 = x % y

	fmt.Println(z1, z2)

}
