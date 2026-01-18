package main

import "fmt"

func main() {
	addResult := Add(10, 5)
	fmt.Printf("10 + 5 = %v\n", addResult)

	subtractResult := Subtract(10, 5)
	fmt.Printf("10 - 5 = %v\n", subtractResult)

	multiplyResult := Multiply(10, 5)
	fmt.Printf("10 * 5 = %v\n", multiplyResult)

	divideResult := Divide(10, 5)
	fmt.Printf("10 / 5 = %v\n", divideResult)

	fmt.Println("\n小数を含む計算:")
	fmt.Printf("2.5 + 1.5 = %v\n", Add(2.5, 1.5))
	fmt.Printf("5.5 - 2.5 = %v\n", Subtract(5.5, 2.5))
	fmt.Printf("3.5 * 2 = %v\n", Multiply(3.5, 2))
	fmt.Printf("7.5 / 2.5 = %v\n", Divide(7.5, 2.5))
}
