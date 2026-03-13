package main

import (
	"fmt"


)

var (
	hey bool = false
)


func main() {
	numbers := []int{10, 20, 30, 40, 50}
	soma := 0

	for _, num := range numbers {
		soma += num 
	}

	fmt.Println("Total:", soma)


values := []float64{1.5, 2.5, 3.0}
var total float64 = 0
for _, value := range values {
	total += value
}
fmt.Println("Total dos valores:", total)


 }