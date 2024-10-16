package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration) // pass in a sleep func
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)

}

type Sleeper interface {
	Sleep()
}
type DefaultSleeper struct{}

// DefaultSleeper implements Sleep method
// hence, compatible with Sleep interface
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)

}
func Countdown(out io.Writer, sleeper Sleeper) {
	// // 1 - not tests friendly
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)

	// 2 - simulat failure
	// for i := countdownStart; i > 0; i-- {
	// 	sleeper.Sleep()
	// }
	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// }

	fmt.Fprint(out, finalWord)
}

func main() {
	// // Use "pointer to a struct"
	// as the Sleep() method requires a "pointer receiver" of *DefaultSleeper
	// s := &DefaultSleeper{}
	// // Using ConfigurableSleeper
	s := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, s)
}
