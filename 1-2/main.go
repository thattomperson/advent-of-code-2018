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

	log.Print(Rept(nums...))
}

func Rept(a ...int) (t int) {
	m := make(map[int]bool)
	m[0] = true
	for {
		for _, i := range a {
			t += i
			if _, ok := m[t]; ok {
				return
			}
			m[t] = true
		}
	}
}
