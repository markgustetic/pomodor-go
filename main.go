package main

import (
	"fmt"
)

func main() {
	fmt.Println("Started")

	p := Pomodoro{}
	p.Start()

	fmt.Println("Ended")
}
