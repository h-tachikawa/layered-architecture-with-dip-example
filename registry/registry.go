package registry

import (
	"layered-architecture-with-dip-example/adapter/api/server/handler"
	"layered-architecture-with-dip-example/domain/repository"
	"layered-architecture-with-dip-example/usecase"
)

var reg = &Registry{}

type Registry struct {
	// Repositories
	VirtualLiverRepository repository.VirtualLiverRepository

	//UseCases
	VirtualLiverUseCase usecase.VirtualLiverUseCase

	//Handlers
	VirtualLiverHandler handler.VirtualLiverHandler
}

func InitRegistry() {
}

func GetRegistry() *Registry {
	return reg
}
