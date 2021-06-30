package module

import (
	"github.com/ProjectAthenaa/sonic"
)

type Task struct {
}

var _ sonic.Module = Task{}
var _ sonic.Module = (*Task)(nil)
