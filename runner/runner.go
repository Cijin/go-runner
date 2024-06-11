package runner

import (
	"os"
	"time"
)

type Runner struct {
	interrupt chan os.Signal

	timeout <-chan time.Time

	tasks []func(int)
}
