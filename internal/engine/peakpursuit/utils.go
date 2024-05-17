package peakpursuit

import (
	"time"
)

func (s *SlotMachine) SetVisualDelay(delay time.Duration) {
	s.GameConfig.VisualDelay = delay
}
