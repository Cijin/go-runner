package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("timeout recieved")
	ErrInterrupt = errors.New("got interrupt")
)

type Runner struct {
	interrupt chan os.Signal

	timer *time.Timer

	timeout time.Duration

	complete chan error

	tasks []func(int)
}

func New(d time.Duration) *Runner {
	timer := time.NewTimer(d)

	return &Runner{
		// os interrupts are non blocking, any uncaught
		// signals are thrown away
		interrupt: make(chan os.Signal, 1),
		timer:     timer,
		timeout:   d,
		complete:  make(chan error),
		tasks:     []func(int){},
	}
}

func (r *Runner) Add(funcs ...func(int)) {
	r.tasks = append(r.tasks, funcs...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	r.timer.Reset(r.timeout)

	select {
	case <-r.timer.C:
		return ErrTimeout

	case err := <-r.complete:
		return err
	}
}

func (r *Runner) run() error {
	for i, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(i)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true

	default:
		return false
	}
}
