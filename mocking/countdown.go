package countdown

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfiguraleSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfiguraleSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper, countdownStart int, finalWord string) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}
