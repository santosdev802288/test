package controller

import (
	"context"
	"testing"

	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	kubgov1 "siigo.com/kubgo/src/api/proto/kubgo/v1"
)

func TestNewController(t *testing.T) {
	// Arrange
	cqrsDispatcher := *new(cqrs.Dispatcher)

	//Act
	controller := NewController(cqrsDispatcher, new(logrus.Logger))

	//Assert
	assert.NotNil(t, controller)
}

func TestNewControllerAddKubgoDontHandler(t *testing.T) {
	// Arrange
	dispatcher := cqrs.NewInMemoryDispatcher(&cqrs.DomainEventsDispatcher{})
	ctx := context.Background()
	// map aggregate to grpc
	kubgo := &kubgov1.Kubgo{}
	kubgoRequest := &kubgov1.AddKubgoRequest{}
	kubgoRequest.Kubgo = kubgo

	// Act
	ctl := NewController(dispatcher, logrus.New())
	var response, error = ctl.AddKubgo(ctx, kubgoRequest)

	//Assert
	assert.Nil(t, response)
	assert.NotNil(t, error)
}

func TestNewControllerLoadKubgoDontHandler(t *testing.T) {
	// Arrange
	dispatcher := cqrs.NewInMemoryDispatcher(&cqrs.DomainEventsDispatcher{})
	ctx := context.Background()

	kubgoRequest := &kubgov1.GetKubgoRequest{}
	kubgoRequest.Id = uuid.NewUUID()

	// Act
	ctl := NewController(dispatcher, logrus.New())
	var response, error = ctl.GetKubgo(ctx, kubgoRequest)

	//Assert
	assert.Nil(t, response)
	assert.NotNil(t, error)
}

func TestNewControllerLoadIdError(t *testing.T) {
	// Arrange
	dispatcher := cqrs.NewInMemoryDispatcher(&cqrs.DomainEventsDispatcher{})
	ctx := context.Background()

	kubgoRequest := &kubgov1.GetKubgoRequest{}

	// Act
	ctl := NewController(dispatcher, logrus.New())
	var response, error = ctl.GetKubgo(ctx, kubgoRequest)

	//Assert
	assert.Nil(t, response)
	assert.NotNil(t, error)
}

func TestNewControllerUpdateKubgo(t *testing.T) {
	// Arrange
	dispatcher := cqrs.NewInMemoryDispatcher(&cqrs.DomainEventsDispatcher{})
	ctx := context.Background()
	// map aggregate to grpc
	kubgo := &kubgov1.Kubgo{}
	kubgoRequest := &kubgov1.UpdateKubgoRequest{}
	kubgoRequest.Kubgo = kubgo

	// Act
	ctl := NewController(dispatcher, logrus.New())
	var response, error = ctl.UpdateKubgo(ctx, kubgoRequest)

	//Assert
	assert.Nil(t, response)
	assert.NotNil(t, error)
}

func TestNewControllerDeleteKubgo(t *testing.T) {
	// Arrange
	dispatcher := cqrs.NewInMemoryDispatcher(&cqrs.DomainEventsDispatcher{})
	ctx := context.Background()

	kubgoRequest := &kubgov1.DeleteKubgoRequest{}
	kubgoRequest.Id = uuid.NewUUID()

	// Act
	ctl := NewController(dispatcher, logrus.New())
	var response, error = ctl.DeleteKubgo(ctx, kubgoRequest)

	//Assert
	assert.NotNil(t, response)
	assert.Nil(t, error)
}

func TestNewControllerLoadAllKubgo(t *testing.T) {
	// Arrange
	dispatcher := cqrs.NewInMemoryDispatcher(&cqrs.DomainEventsDispatcher{})
	ctx := context.Background()
	var protoReq *emptypb.Empty

	kubgoRequest := &kubgov1.GetKubgoRequest{}
	kubgoRequest.Id = uuid.NewUUID()

	// Act
	ctl := NewController(dispatcher, logrus.New())
	var response, error = ctl.FindKubgos(ctx, protoReq)

	//Assert
	assert.Nil(t, response)
	assert.NotNil(t, error)
}
