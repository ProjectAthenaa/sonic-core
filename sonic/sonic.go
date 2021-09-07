package sonic

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/antibots/perimeterx"
	"github.com/ProjectAthenaa/sonic-core/sonic/antibots/shape"
	"github.com/ProjectAthenaa/sonic-core/sonic/antibots/ticket"
)

// NewPerimeterXClient Is a wrapper for the internal NewClient method of PerimeterX
func NewPerimeterXClient(svcURL ...string) (perimeterx.PerimeterXClient, error) {
	return perimeterx.NewClient(svcURL...)
}

// NewShapeClient Is a wrapper for the internal NewClient method of Shape
func NewShapeClient(svcURL ...string) (shape.ShapeClient, error) {
	return shape.NewClient(svcURL...)
}

// NewTicketClient Is a wrapper for the internal NewClient method of Shape
func NewTicketClient(svcURL ...string) (ticket.TicketClient, error) {
	return ticket.NewClient(svcURL...)
}

type Site uint

const (
	NewBalance Site = iota + 1
	Hibbet
)
