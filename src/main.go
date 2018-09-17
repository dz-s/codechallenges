package main

import (
	"fmt"
)

//Solution for BinaryGap task
func Solution(N int) int {
	var maxGap int = 0
	var currGap int = 0
	var insideGap = false
	if N > 0 {

		for i := N; i > 0; i /= 2 {
			if i%2 == 0 {
				currGap++
			} else {
				if insideGap {
					if currGap > maxGap {
						maxGap = currGap
					}
					currGap = 0
				} else {
					insideGap = true
				}

			}
		}
	}

	return maxGap

}

func main() {
	fmt.Printf("result for of 1041: %d\n", Solution(1041))
	fmt.Printf("result for of 32: %d\n", Solution(32))
	occurences()
}
