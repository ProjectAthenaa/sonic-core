package sonic

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/antibots/perimeterx"
)

//Is a wrapper for the internal NewClient method of PerimeterX
func NewPerimeterXClient(svcURL ...string) (perimeterx.PerimeterXClient, error) {
	return perimeterx.NewClient(svcURL...)
}
