package main

import "fmt"
import "math/big"

func main() {
	f1, f2, f3 := 10.5, 10.9, 10.2
	floatSum := f1 + f2 + f3
	fmt.Println("floatsum = ", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(10.5)
	b2.SetFloat64(10.9)
	b3.SetFloat64(10.2)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Println("big sum = ", &bigSum)
}
