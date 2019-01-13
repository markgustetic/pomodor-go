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
	longBreakTime  time.Duration = 2 * time.Minute
)

var pomodoroCount int

//SetTimer will start the pomodoro timer
func (p *Pomodoro) SetTimer() {
	printStatus(pomodoroTime)

	pomodoroCount++
}

//SetBreak will start the break timer
func (p *Pomodoro) SetBreak() {
	var breakTime time.Duration

	if pomodoroCount < 4 {
		breakTime = shortBreakTime
	} else {
		breakTime = longBreakTime
		pomodoroCount = 0
	}

	printStatus(breakTime)
}

func printStatus(pomodoroDuration time.Duration) {
	statusChan := runTicker(pomodoroDuration)

	timeCount := pomodoroDuration

	for {
		select {
		case <-statusChan.tickerChan:
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
