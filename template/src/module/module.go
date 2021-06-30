package module

import (
	"github.com/ProjectAthenaa/sonic-core/sonic"
)

type Task struct {
}



var _ sonic.Module = Task{}
var _ sonic.Module = (*Task)(nil)
