package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func scan_example() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			// handle error
			fmt.Println(s)
		} else {
			fmt.Println(i * 2)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
