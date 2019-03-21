package registry

import (
	"layered-architecture-with-dip-example/domain/repository"
	"layered-architecture-with-dip-example/infrastructure/persistence/datastore"
	"layered-architecture-with-dip-example/interfaces/api/server/handler"
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
	// Repositories
	reg.VirtualLiverRepository = datastore.NewVirtualLiverRepository()

	//UseCases
	reg.VirtualLiverUseCase = usecase.NewVirtualLiverUseCase(reg.VirtualLiverRepository)

	//Handlers
	reg.VirtualLiverHandler = handler.NewVirtualLiverHandler(reg.VirtualLiverUseCase)
}

func GetRegistry() *Registry {
	return reg
}
