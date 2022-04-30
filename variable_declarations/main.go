package main

import "fmt"

func main() {
	var i int
	var U, V, W float64
	var k = 0
	var x, y float32 = -1, -2
	var (
		j       int
		u, v, s = 2.0, 3.0, "bar"
	)

	fmt.Printf("[var i int] => %v %T\n", i, i)
	fmt.Printf("[var U, V, W float64] => %v %T, %v %T, %v %T\n", U, U, V, V, W, W)
	fmt.Printf("[var k = 0] => %v %T\n", k, k)
	fmt.Printf("[var x, y float32 = -1, -2] => %v %T, %v %T\n", x, x, y, y)
	fmt.Printf("[j] => %v %T\n", j, j)
	fmt.Printf("[u, v, s = 2.0, 3.0, \"bar\"] => %v %T, %v %T, %v %T\n", u, u, v, v, s, s)

}
