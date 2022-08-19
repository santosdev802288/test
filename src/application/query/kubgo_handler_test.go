package query_test

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockServices "siigo.com/kubgo/mocks/src/domain/services"
	"siigo.com/kubgo/src/application/query"
	"siigo.com/kubgo/src/domain/kubgo"
	"testing"
)

func TestLoadKubgoQueryHandlerDomainServiceError_Handle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	id, _ := uuid.FromString(cqrs.NewUUID())
	errorMessage := "domain error"

	domainKubgoService.
		On("Get", mock.Anything).
		Return(nil, errors.New(errorMessage))

	commandMessage := cqrs.NewQueryMessage(&query.LoadKubgoQuery{Id: id})

	// Act
	handler := query.NewKubgoQueryHandler(domainKubgoService)
	response, err := handler.Handle(commandMessage)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, response)

	domainKubgoService.AssertNumberOfCalls(t, "Get", 1)
	domainKubgoService.AssertCalled(t, "Get", id)
	domainKubgoService.AssertExpectations(t)
}

func TestLoadKubgoQueryHandler_Handle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	id := cqrs.NewUUID()
	uid, _ := uuid.FromString(id)
	kubgoMock := kubgo.NewKubgo(id)

	domainKubgoService.
		On("Get", mock.Anything).
		Return(kubgoMock, nil)

	commandMessage := cqrs.NewQueryMessage(&query.LoadKubgoQuery{Id: uid})

	// Act
	handler := query.NewKubgoQueryHandler(domainKubgoService)
	response, err := handler.Handle(commandMessage)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, response)

	domainKubgoService.AssertNumberOfCalls(t, "Get", 1)
	domainKubgoService.AssertCalled(t, "Get", uid)
	domainKubgoService.AssertExpectations(t)
}

func TestLoadAllKubgoQueryHandler_Handle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}

	domainKubgoService.
		On("LoadAll", mock.Anything).
		Return([]*kubgo.Kubgo{}, nil)

	commandMessage := cqrs.NewQueryMessage(&query.LoadAllKubgoQuery{})

	// Act
	handler := query.NewKubgoQueryHandler(domainKubgoService)
	response, err := handler.Handle(commandMessage)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, response)

	domainKubgoService.AssertNumberOfCalls(t, "LoadAll", 1)
	domainKubgoService.AssertExpectations(t)
}

func TestLoadAllKubgoQueryHandler_ErrorHandle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	errorMessage := "domain error"
	domainKubgoService.
		On("LoadAll", mock.Anything).
		Return(nil, errors.New(errorMessage))

	commandMessage := cqrs.NewQueryMessage(&query.LoadAllKubgoQuery{})

	// Act
	handler := query.NewKubgoQueryHandler(domainKubgoService)
	response, err := handler.Handle(commandMessage)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, response)

	domainKubgoService.AssertNumberOfCalls(t, "LoadAll", 1)
	domainKubgoService.AssertExpectations(t)
}
