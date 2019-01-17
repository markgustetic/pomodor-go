package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/markgustetic/pomodor-go/pomodoro"
)

func main() {
	p := pomodoro.Pomodoro{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Pomodoro Started")
		statusChan := p.SetTimer()
		getTime(statusChan)

		resetLine(fmt.Sprintf("Pomodoro Count: %d\n", p.PomodoroCount()))

		fmt.Printf("Break Started\n")
		statusChan = p.SetBreak()
		getTime(statusChan)

		fmt.Print("\nPress enter to start next Pomodoro")
		reader.ReadString('\n')
	}
}

func getTime(statusChan pomodoro.StatusChan) {
	timeCount := statusChan.PomodoroDuration

	for {
		select {
		case <-statusChan.TickerChan:
			resetLine(timeCount)
			timeCount = timeCount - time.Second
		case <-statusChan.DoneChan:
			return
		}
	}
}

func resetLine(outputLine interface{}) {
	fmt.Printf("\033[2K\r%s", outputLine)
}
