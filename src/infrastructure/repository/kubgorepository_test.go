package repository

import (
	"context"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"siigo.com/kubgo/it/containers"
	"siigo.com/kubgo/src/domain/kubgo"
	"testing"
)

func TestSaveSuccess(t *testing.T) {
	//Arrange
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())
	repositoryKubgo := NewKubgoRepository(context.Background(), collection, logrus.New())
	kubgoDocument := kubgo.NewKubgo(cqrs.NewUUID())

	// Act  --------
	response := <-repositoryKubgo.Save(kubgoDocument)

	// Assert --------
	assert.Nil(t, response)
}

func TestDeleteSuccess(t *testing.T) {
	//Arrange
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())
	repositoryKubgo := NewKubgoRepository(context.Background(), collection, logrus.New())
	kubgoDocument := kubgo.NewKubgo(cqrs.NewUUID())

	// Act  --------
	response := <-repositoryKubgo.Delete(kubgoDocument)

	// Assert --------
	assert.Nil(t, response)
}

func TestUpdateSuccess(t *testing.T) {
	//Arrange
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())
	repositoryKubgo := NewKubgoRepository(context.Background(), collection, logrus.New())
	kubgoDocument := kubgo.NewKubgo(cqrs.NewUUID())

	// Act  --------
	response := <-repositoryKubgo.Update(kubgoDocument)

	// Assert --------
	assert.Nil(t, response)
}
