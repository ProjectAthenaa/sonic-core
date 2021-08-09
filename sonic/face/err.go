package face

import "errors"

var (
	ErrFirstTaskCommandError = errors.New("first command error")
	ErrFirstTaskDataError = errors.New("first data error")
	ErrTaskIsRunning      = errors.New("task is running")
	ErrTaskNotInit        = errors.New("task not init")
	ErrTaskIsNotRunning   = errors.New("task is not running")
	ErrTaskPauseTimeout = errors.New("task pause has timed out")
)
