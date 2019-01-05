package main

import "fmt"

type Pomodoro struct{}

func (p *Pomodoro) Start() {
	fmt.Println("Pomodoro Started")
}
