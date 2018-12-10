package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var nums []int

	for s.Scan() {
		var num int
		fmt.Sscanf(s.Text(), "%d", &num)
		nums = append(nums, num)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print(Sum(nums...))
}

func Sum(a ...int) (t int) {
	for _, i := range a {
		t += i
	}
	return
}
