package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/markgustetic/pomodor-go/pomodoro"
)

func main() {
	p := pomodoro.Pomodoro{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Pomodoro Started")
		p.SetTimer()

		fmt.Printf("\nBreak Started\n")
		p.SetBreak()

		fmt.Print("\nPress enter to start next Pomodoro")
		reader.ReadString('\n')
	}
}
