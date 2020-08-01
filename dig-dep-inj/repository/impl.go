package repository

import "fmt"

type ClientRepositoryImpl struct{}

func (c *ClientRepositoryImpl) Do() {
	fmt.Println("Repository done!")
}

func NewClientRepositoryImpl() ClientRepository {
	return &ClientRepositoryImpl{}
}
