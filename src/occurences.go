package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func occurences() {

	digits := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	digit := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		for _, char := range s {
			digit = int(char - '0')
			digits[digit]++
		}

		for i := 0; i < 10; i++ {
			fmt.Println(i, digits[i])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
