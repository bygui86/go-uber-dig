package main

import (
	"github.com/bygui86/go-uber-dig/no-dep-inj/repository"
	"github.com/bygui86/go-uber-dig/no-dep-inj/service"
)

func main() {
	// create repository
	repo := repository.NewClientRepositoryImpl()

	// create service with repository as parameter
	svc := service.NewClientServiceImpl(repo)

	// perform actions
	svc.DoAction()
}
