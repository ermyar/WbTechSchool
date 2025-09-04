package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 40)
	a.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
	b := new(big.Int)
	b.SetString("1000000000000000000000000000000000000000000000", 10)
	c := big.NewInt(7)
	res := new(big.Int)
	fmt.Println(b, "-", c, "=", res.Sub(b, c))

	fmt.Println(a, "+", b, "=", res.Add(a, b))

	fmt.Println(res.Mul(a, b))

	fmt.Println(b, "/", a, "=", res.Div(b, a))
	fmt.Println(a, "/", b, "=", res.Div(a, b))
}
