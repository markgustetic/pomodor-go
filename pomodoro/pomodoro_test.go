package pomodoro

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestEmptyCount(t *testing.T) {
	pomodoro := Pomodoro{}

  assert.Equal(t, pomodoro.PomodoroCount(), 0)
}
