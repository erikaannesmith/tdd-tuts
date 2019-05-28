package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"
const write = "write"
const sleep = "sleep"

type Sleeper interface { Sleep() }

type CountdownOperationsSpy struct { Calls []string }

type ConfigurableSleeper struct {
	duration time.Duration
	sleep 	 func(time.Duration)
}

type SpyTime struct { durationSlept time.Duration }

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}