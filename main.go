package main

import (
	"fmt"
)

type num struct {
	a int
	b int
}

func add(n1, n2 num) num {
	return num{n1.a + n2.a, n1.b + n2.b}

}
func main() {
	var n1 = num{3, 5}
	var n2 = num{2, 6}
	var n3 = add(n1, n2)
	fmt.Println("(", n3.a, ",", n3.b, ")")
}
