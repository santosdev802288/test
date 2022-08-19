//go:build integration
// +build integration

package it

import (
	"context"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"siigo.com/kubgo/it/containers"
	"siigo.com/kubgo/src/domain/kubgo"
	"siigo.com/kubgo/src/infrastructure/finder"
	"siigo.com/kubgo/src/infrastructure/repository"
	"sync"
	"testing"
)

func TestSaveKubgoRepository(t *testing.T) {
	t.Parallel()

	// Arrange
	collection, container := containers.BuildMongoInfrastructure(t)
	kubgoRepository := repository.NewKubgoRepository(context.Background(), collection, logrus.New())
	kubgoFinder := finder.NewKubgoFinder(context.Background(), collection, logrus.New())
	defer container.Terminate(context.Background())

	// Act  --------
	totalKubgos := 50
	wg := sync.WaitGroup{}
	wg.Add(totalKubgos)

	for i := 0; i < totalKubgos; i++ {
		go func() {
			defer wg.Done()
			kubgoDocument := kubgo.NewKubgo(cqrs.NewUUID())

			// save kubgos
			kubgoError := <-kubgoRepository.Save(kubgoDocument)
			if kubgoError != nil {
				t.Fatalf("Failed to save in mongo database: %v", kubgoError)
			}
		}()
	}

	wg.Wait()

	// Assert --------

	// Find all
	records := <-kubgoFinder.GetAll()

	assert.Nil(t, records.Error)
	assert.Equal(t, len(records.Kubgos), totalKubgos)

}
