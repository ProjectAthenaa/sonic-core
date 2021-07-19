package face

import "errors"

var ErrFirstTaskCommandError = errors.New("first command error")
var ErrFirstTaskDataError = errors.New("first data error")
var ErrTaskIsRunning = errors.New("task is running")
var ErrTaskNotInit = errors.New("task not init")
var ErrTaskIsNotRunning = errors.New("task is not running")
