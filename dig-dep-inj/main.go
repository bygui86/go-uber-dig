package main

import (
	"github.com/bygui86/go-uber-dig/dig-dep-inj/repository"
	"github.com/bygui86/go-uber-dig/dig-dep-inj/service"
	"go.uber.org/dig"
)

var digContainer *dig.Container

var svc service.ClientService

func main() {
	// prepare all components for dependency injection
	provideComponents()

	// instantiate all components with dependency injection
	err := initComponents()
	if err != nil {
		panic(err)
	}

	// perform actions
	svc.DoAction()
}

func initComponents() error {
	return digContainer.Invoke(
		func(s service.ClientService) {
			svc = s
		},
	)
}

func provideComponents() {
	digContainer = dig.New()

	digContainer.Provide(
		func() repository.ClientRepository {
			return repository.NewClientRepositoryImpl()
		},
	)

	digContainer.Provide(service.NewClientServiceImpl)
}
