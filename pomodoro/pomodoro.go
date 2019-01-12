package pomodoro

import (
	"fmt"
	"time"
)

//Pomodoro is the struct used to call methods on the timer
type Pomodoro struct{}

//StatusChan holds the two channels used to see where the timer is currently at
type StatusChan struct {
	tickerChan <-chan time.Time
	doneChan   <-chan bool
}

const (
	pomodoroTime   time.Duration = 1 * time.Minute
	shortBreakTime time.Duration = 1 * time.Minute
	longBreakTime  time.Duration = 1 * time.Minute
)

var pomodoroCount int

//SetTimer will start the pomodoro timer
func (p *Pomodoro) SetTimer() {
	fmt.Println("Pomodoro Started")

	printStatus(pomodoroTime)

	pomodoroCount++

	fmt.Println()
	fmt.Println("Pomodoro Finished")
}

//SetBreak will start the break timer
func (p *Pomodoro) SetBreak() {
	fmt.Println("Break Started")

	var breakTime time.Duration

	if pomodoroCount < 4 {
		breakTime = shortBreakTime
	} else {
		breakTime = longBreakTime
		pomodoroCount = 0
	}

	printStatus(breakTime)

	fmt.Println()
	fmt.Println("Break Ended")
}

func printStatus(pomodoroDuration time.Duration) {
	statusChan := runTicker(pomodoroDuration)

	timeCount := pomodoroDuration

	for {
		select {
		case <-statusChan.tickerChan:
			//fmt.Printf("\r%s", timeCount)
			fmt.Printf("\033[2K\r%s", timeCount)
			timeCount = timeCount - time.Second
		case <-statusChan.doneChan:
			return
		}
	}
}

func runTicker(pomodoroDuration time.Duration) StatusChan {
	tickerChan := time.NewTicker(time.Second).C
	doneChan := make(chan bool)

	go func() {
		time.Sleep(pomodoroDuration)
		doneChan <- true
	}()

	return StatusChan{tickerChan: tickerChan, doneChan: doneChan}
}
