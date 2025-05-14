## 07-Arithmetic Operators

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Modulus:", result)

	const p float64 = 22 / 7.0
	fmt.Println("Constant p:", p)

	// Overflow example
	var maxInt int64 = 9223372036854775807 // Maximum value for int64
	fmt.Println("Max Int:", maxInt)

	maxInt = maxInt + 1
	fmt.Println("Overflowed Max Int:", maxInt)

	// Overflow example with uint
	var maxUint uint64 = 18446744073709551615 // Maximum value for uint64
	maxUint = maxUint + 1
	fmt.Println("Overflowed Max Uint:", maxUint)

	// Underflow example
	var smallFloat float64 = -1.7976931348623157e+308 // Minimum value for float64
	fmt.Println("Small Float:", smallFloat)
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("Underflowed Small Float:", smallFloat)

}
```