package finder

import (
	"context"
	"testing"

	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"siigo.com/kubgo/it/containers"
	"siigo.com/kubgo/src/domain/kubgo"
	"siigo.com/kubgo/src/infrastructure/repository"
)

func TestGetAll(t *testing.T) {
	//Arrange
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())

	finderKubgo := NewKubgoFinder(context.Background(), collection, logrus.New())
	repositoryKubgo := repository.NewKubgoRepository(context.Background(), collection, logrus.New())

	kubgoDocument := kubgo.NewKubgo(cqrs.NewUUID())
	kubgoError := <-repositoryKubgo.Save(kubgoDocument)
	if kubgoError != nil {
		t.Fatalf("Failed to save in mongo database: %v", kubgoError)
	}

	// Act  --------
	response := <-finderKubgo.GetAll()

	// Assert --------
	assert.Nil(t, response.Error)
	assert.Equal(t, len(response.Kubgos), 1)
}

func TestGetUnique(t *testing.T) {
	//Arrange
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())
	finderKubgo := NewKubgoFinder(context.Background(), collection, logrus.New())

	id, invalidIdError := uuid.FromString(cqrs.NewUUID())
	if invalidIdError != nil {
		t.Error(invalidIdError)
	}

	// Act  --------
	response := <-finderKubgo.Get(id)

	// Assert --------
	assert.NotNil(t, response.Error) //Kubgo NOt Found
	assert.NotNil(t, response.Kubgo)
}
