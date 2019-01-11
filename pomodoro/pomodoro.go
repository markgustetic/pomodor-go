package pomodoro

import (
	"fmt"
	"time"
)

//Pomodoro is the struct used to call methods on the timer
type Pomodoro struct{}

const (
	pomodoroTime   time.Duration = 1 * time.Minute
	shortBreakTime time.Duration = 1 * time.Minute
	longBreakTime  time.Duration = 1 * time.Minute
)

var pomodoroCount int

//SetTimer will start the pomodoro timer
func (p *Pomodoro) SetTimer() {
	fmt.Println("Pomodoro Started")

	printStatus()

	pomodoroCount++

	fmt.Println()
	fmt.Println("Pomodoro Finished")
}

func runLoop() chan time.Duration {
	statusChan := make(chan time.Duration)

	tickerChannel := time.NewTicker(time.Second).C
	doneChan := make(chan bool)

	go func() {
		time.Sleep(time.Second * 5)
		doneChan <- true
	}()

	go func() {
		timeCount := time.Minute * 25
		statusChan <- timeCount

		for {
			select {
			case <-tickerChannel:
				//fmt.Printf("\r%s", timeCount)
				timeCount = timeCount - time.Second
				statusChan <- timeCount
			case <-doneChan:
				return
			}
		}
	}()

	return statusChan
}

func printStatus() {
	statusChan := runLoop()

	for {
		select {
		case currentSecond := <-statusChan:
			fmt.Printf("\r%s", currentSecond)
		default:
			return
		}
	}
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
