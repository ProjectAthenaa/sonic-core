package face

import "errors"

var (
	ErrFirstTaskCommandError = errors.New("first command error")
	ErrFirstTaskDataError    = errors.New("first data error")
	ErrTaskIsRunning         = errors.New("task is running")
	ErrTaskNotInit           = errors.New("task not init")
	ErrTaskIsNotRunning      = errors.New("task is not running")
	ErrTaskPauseTimeout      = errors.New("task pause has timed out")
	ErrTaskIsPaused          = errors.New("task is already paused")
	ErrTaskIsNotPaused       = errors.New("task is not paused")
	ErrTaskHasNoData         = errors.New("task has no data")
	ErrCtxDeadlineExceeded= errors.New("context deadline exceeded")
)
