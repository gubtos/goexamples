package variadic

import "fmt"

func receivingArray(values []int) {
	for i := range values {
		fmt.Println(i)
	}
}

func receivingVariadic(values ...int) {
	for i := range values {
		fmt.Println(i)
	}
}

func variadicComparison() {
	myArray := []int{1, 2, 3}

	receivingArray(myArray)

	receivingVariadic(1, 2, 3)
	receivingVariadic(myArray...)
}
