package antibots

import (
	"errors"
	"os"
)

var (
	ExactlyOneArgumentError = errors.New("This function accepts 0 or 1 arguments")
	DebugModeParameter      = errors.New("Debug modes need custom address")
	Debug                   = os.Getenv("DEBUG")
)
