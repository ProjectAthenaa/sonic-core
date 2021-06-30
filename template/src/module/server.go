package module

import module "github.com/ProjectAthenaa/sonic/protos"

type Server struct {
	module.UnimplementedModuleServer
}

func (s Server) Task(server module.Module_TaskServer) error {
	panic("implement me")
}

