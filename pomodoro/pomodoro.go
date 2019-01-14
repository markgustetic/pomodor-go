package pomodoro

import (
	"time"
)

//Pomodoro is the struct used to call methods on the timer
type Pomodoro struct{}

//StatusChan holds the two channels used to see where the timer is currently at
type StatusChan struct {
	TickerChan       <-chan time.Time
	DoneChan         <-chan bool
	PomodoroDuration time.Duration
}

const (
	pomodoroTime   time.Duration = 1 * time.Minute
	shortBreakTime time.Duration = 1 * time.Minute
	longBreakTime  time.Duration = 2 * time.Minute
)

var pomodoroCount int

//SetTimer will start the pomodoro timer
func (p *Pomodoro) SetTimer() StatusChan {
	statusChan := runTicker(pomodoroTime)

	pomodoroCount++

	return statusChan
}

//SetBreak will start the break timer
func (p *Pomodoro) SetBreak() StatusChan {
	var breakTime time.Duration

	if pomodoroCount < 4 {
		breakTime = shortBreakTime
	} else {
		breakTime = longBreakTime
		pomodoroCount = 0
	}

	statusChan := runTicker(breakTime)

	return statusChan
}

func runTicker(pomodoroDuration time.Duration) StatusChan {
	tickerChan := time.NewTicker(time.Second).C
	doneChan := make(chan bool)

	go func() {
		time.Sleep(pomodoroDuration)
		doneChan <- true
	}()

	return StatusChan{TickerChan: tickerChan, DoneChan: doneChan, PomodoroDuration: pomodoroDuration}
}
