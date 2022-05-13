package main

import (
	"fmt"
	"math/rand"
	"time"
)

type queue []int

func main() {
	q := queue{}

	go q.manufacturer()
	go q.manufacturer()
	go q.manufacturer()

	go q.user()
	go q.user()

	fmt.Scanln()
}

func (q *queue) user() {
	for len(*(q)) != 0 {
		*(q) = (*(q))[:len(*(q))-1]
		fmt.Print(*(q), "\n\n", "Items: ", len(*(q)), "\n\n")
		time.Sleep(150 * time.Millisecond)
	}
}

func (q *queue) manufacturer() {
	for len(*(q)) <= 100 {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		*(q) = append(*(q), r.Intn(200))
		fmt.Print(*(q), "\n\n", "Items: ", len(*(q)), "\n\n")
		time.Sleep(150 * time.Millisecond)
	}

	for {
		if len(*(q)) <= 80 {
			q.manufacturer()
		}
	}
}
