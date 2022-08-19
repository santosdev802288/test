//go:build integration
// +build integration

package it

import (
	"context"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"siigo.com/kubgo/it/containers"
	"siigo.com/kubgo/src/domain/kubgo"
	"siigo.com/kubgo/src/infrastructure/finder"
	"siigo.com/kubgo/src/infrastructure/repository"
	"testing"
)

func TestGetAllKubgosFinder(t *testing.T) {

	t.Parallel()

	// Arrange  --------
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())

	finderKubgo := finder.NewKubgoFinder(context.Background(), collection, logrus.New())

	// Act  --------
	response := <-finderKubgo.GetAll()

	// Assert --------
	assert.Nil(t, response.Error)
	assert.Equal(t, len(response.Kubgos), 0)
}

func TestGetAllKubgosWithElementsFinder(t *testing.T) {

	t.Parallel()

	// Arrange  ----------
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())

	finderKubgo := finder.NewKubgoFinder(context.Background(), collection, logrus.New())
	repository := repository.NewKubgoRepository(
		context.Background(), collection, logrus.New(),
	)

	id := cqrs.NewUUID()
	kubgoDocument := kubgo.NewKubgo(id)
	kubgoError := <-repository.Save(kubgoDocument)
	if kubgoError != nil {
		t.Fatalf("Failed to save in mongo database: %v", kubgoError)
	}

	// Act  --------
	response := <-finderKubgo.GetAll()

	// Assert --------
	assert.Nil(t, response.Error)
	assert.Equal(t, len(response.Kubgos), 1)
	assert.Equal(t, response.Kubgos[0].Id, id)
	assert.Equal(t, response.Kubgos[0].OccurredAt, kubgoDocument.OccurredAt)
}

func TestGetKubgoNotFoundFinder(t *testing.T) {

	t.Parallel()

	// Arrange  --------
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())

	finderKubgo := finder.NewKubgoFinder(context.Background(), collection, logrus.New())

	id, invalidIdError := uuid.FromString(cqrs.NewUUID())
	if invalidIdError != nil {
		t.Error(invalidIdError)
	}

	// Act  --------
	response := <-finderKubgo.Get(id)

	// Assert --------
	assert.NotNil(t, response.Error)
	assert.Equal(t, response.Error.Error(), "kubgo not found")
}

func TestGetKubgoFinder(t *testing.T) {

	t.Parallel()

	// Arrange
	collection, container := containers.BuildMongoInfrastructure(t)
	defer container.Terminate(context.Background())

	finderKubgo := finder.NewKubgoFinder(context.Background(), collection, logrus.New())
	repository := repository.NewKubgoRepository(
		context.Background(), collection, logrus.New(),
	)

	id := cqrs.NewUUID()
	kubgoDocument := kubgo.NewKubgo(id)
	kubgoError := <-repository.Save(kubgoDocument)
	if kubgoError != nil {
		t.Fatalf("Failed to save in mongo database: %v", kubgoError)
	}

	// Act  --------
	uid, invalidIdError := uuid.FromString(id)
	if invalidIdError != nil {
		t.Error(invalidIdError)
	}
	response := <-finderKubgo.Get(uid)

	// Assert --------
	assert.Nil(t, response.Error)
	assert.Equal(t, response.Kubgo.Id, id)
	assert.Equal(t, response.Kubgo.OccurredAt, kubgoDocument.OccurredAt)
}
