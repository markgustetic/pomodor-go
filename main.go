package main

import (
	"fmt"

	"github.com/markgustetic/pomodor-go/pomodoro"
)

func main() {
	fmt.Println("Started")

	p := pomodoro.Pomodoro{}
	p.Start()

	fmt.Println("Ended")
}
