package shape

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/antibots"
	"google.golang.org/grpc"
)

func NewClient(svcURL ...string) (ShapeClient, error) {
	var address = "shape.antibots.svc.cluster.local:3000"
	if antibots.Debug == "1" {
		if len(svcURL) == 1 {
			address = svcURL[0]
		} else if len(svcURL) == 0 {
			return nil, antibots.DebugModeParameter
		} else {
			return nil, antibots.ExactlyOneArgumentError
		}
	} else {
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return NewShapeClient(conn), nil
}
