package pomodoro

import (
	"fmt"
	"time"
)

//Pomodoro is the struct used to call methods on the timer
type Pomodoro struct{}

//StatusChan will allow the caller to check the current time of the Pomodoro and check if it's finished
type StatusChan struct {
	doneChan      chan bool
	currentSecond time.Duration
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

	runLoop()

	pomodoroCount++

	fmt.Println("Pomodoro Finished")
}

func runLoop() chan StatusChan {
	statusChan := make(chan StatusChan)

	tickerChannel := time.NewTicker(time.Second).C
	doneChan := make(chan bool)

	go func() {
		time.Sleep(time.Second * 5)
		doneChan <- true
	}()

	go func() {
		timeCount := time.Minute * 25

		for {
			select {
			case <-tickerChannel:
				fmt.Printf("\r%s", timeCount)
				timeCount = timeCount - time.Second
			case <-doneChan:
				return
			}
		}
	}()

	return statusChan
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

	timer := time.NewTimer(breakTime)

	<-timer.C

	fmt.Println("Break Ended")
}
