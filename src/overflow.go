package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getBitImplCount(value int) int {
	s := int64(value)
	bitImpl := strconv.FormatInt(s, 2)
	bitCount := 0
	for _, char := range bitImpl {
		bit := int(char - '0')
		if bit == 1 {
			bitCount++
		}
	}
	return bitCount
}

func overflow() {

	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	result := 0
	for scanner.Scan() {
		X, _ := strconv.Atoi(scanner.Text())

		for i := 1; i <= X; i++ {
			if getBitImplCount(X/i) <= getBitImplCount(i) {
				result++
			}
		}

		fmt.Println(result)
		result = 0
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
