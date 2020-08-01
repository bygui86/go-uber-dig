package service

import (
	"fmt"
	"github.com/bygui86/go-uber-dig/dig-dep-inj/repository"
)

type ClientServiceImpl struct {
	clientRepository repository.ClientRepository
}

func (c *ClientServiceImpl) DoAction() {
	c.clientRepository.Do()
	fmt.Println("Service done!")
}

func NewClientServiceImpl(clientRepository repository.ClientRepository) ClientService {
	return &ClientServiceImpl{clientRepository}
}
