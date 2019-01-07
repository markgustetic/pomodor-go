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

	timer := time.NewTimer(pomodoroTime)

	<-timer.C

	pomodoroCount++

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

	timer := time.NewTimer(breakTime)

	<-timer.C

	fmt.Println("Break Ended")
}
