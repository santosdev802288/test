package command_test

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockServices "siigo.com/kubgo/mocks/src/domain/services"
	"siigo.com/kubgo/src/application/command"
	"siigo.com/kubgo/src/domain/kubgo"
	"testing"
)

func TestCreateKubgoCommandHandlerWithBudgedError_Handle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	domainError := &kubgo.ExceedsCostKubgoError{}

	domainKubgoService.
		On("Save", mock.Anything, mock.Anything).
		Return(domainError)

	id := cqrs.NewUUID()
	ct := kubgo.NewKubgo(id)
	commandMessage := cqrs.NewCommandMessage(id, &command.CreateKubgoCommand{
		Kubgo: ct,
	})

	// Act
	handler := command.NewKubgoCommandHandler(domainKubgoService)
	resp, err := handler.Handle(commandMessage)

	// Assert
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, domainError))
	assert.EqualError(t, err, domainError.Error())
	assert.Nil(t, resp)

	domainKubgoService.AssertNumberOfCalls(t, "Save", 1)
	domainKubgoService.AssertCalled(t, "Save", ct, cqrs.Int(ct.OriginalVersion()))
	domainKubgoService.AssertExpectations(t)

}

func TestCreateKubgoCommandHandlerSuccess_Handle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	id := cqrs.NewUUID()
	ct := kubgo.NewKubgo(id)

	domainKubgoService.
		On("Save", mock.Anything, mock.Anything).
		Return(nil)

	commandMessage := cqrs.NewCommandMessage(id, &command.CreateKubgoCommand{
		Kubgo: ct,
	})

	// Act
	handler := command.NewKubgoCommandHandler(domainKubgoService)
	response, err := handler.Handle(commandMessage)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, id, response.(*kubgo.Kubgo).Id)

	domainKubgoService.AssertNumberOfCalls(t, "Save", 1)
	domainKubgoService.AssertCalled(t, "Save", ct, cqrs.Int(ct.OriginalVersion()))
	domainKubgoService.AssertExpectations(t)

}

func TestDeleteKubgoCommandHandlerSuccess_Handle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	id := cqrs.NewUUID()

	domainKubgoService.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil)

	commandMessage := cqrs.NewCommandMessage(id, &command.DeleteKubgoCommand{
		Id: id,
	})

	// Act
	handler := command.NewKubgoCommandHandler(domainKubgoService)
	_, err := handler.Handle(commandMessage)

	// Assert
	assert.Nil(t, err)

	domainKubgoService.AssertNumberOfCalls(t, "Delete", 1)
	domainKubgoService.AssertExpectations(t)

}

func TestDeleteKubgoCommandHandlerSuccess_ErrorHandle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	id := cqrs.NewUUID()
	domainError := errors.New("kubgo problems")

	domainKubgoService.
		On("Delete", mock.Anything, mock.Anything).
		Return(domainError)

	commandMessage := cqrs.NewCommandMessage(id, &command.DeleteKubgoCommand{
		Id: id,
	})

	// Act
	handler := command.NewKubgoCommandHandler(domainKubgoService)
	_, err := handler.Handle(commandMessage)

	// Assert
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, domainError))
	assert.EqualError(t, err, domainError.Error())

	domainKubgoService.AssertNumberOfCalls(t, "Delete", 1)
	domainKubgoService.AssertExpectations(t)

}

func TestUpdateKubgoCommandHandlerSuccess_Handle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	id := cqrs.NewUUID()
	ct := kubgo.NewKubgo(id)

	domainKubgoService.
		On("Update", mock.Anything, mock.Anything).
		Return(nil)

	commandMessage := cqrs.NewCommandMessage(id, &command.UpdateKubgoCommand{
		Kubgo: ct,
	})

	// Act
	handler := command.NewKubgoCommandHandler(domainKubgoService)
	response, err := handler.Handle(commandMessage)

	// Assert
	assert.Nil(t, err)
	assert.Nil(t, response)

	domainKubgoService.AssertNumberOfCalls(t, "Update", 1)
	domainKubgoService.AssertCalled(t, "Update", ct, cqrs.Int(ct.OriginalVersion()))
	domainKubgoService.AssertExpectations(t)

}

func TestUpdateKubgoCommandHandlerSuccess_ErrorHandle(t *testing.T) {

	// Arrange
	domainKubgoService := &mockServices.IKubgoService{}
	id := cqrs.NewUUID()
	ct := kubgo.NewKubgo(id)
	domainError := errors.New("kubgo problems")

	domainKubgoService.
		On("Update", mock.Anything, mock.Anything).
		Return(domainError)

	commandMessage := cqrs.NewCommandMessage(id, &command.UpdateKubgoCommand{
		Kubgo: ct,
	})

	// Act
	handler := command.NewKubgoCommandHandler(domainKubgoService)
	_, err := handler.Handle(commandMessage)

	// Assert
	assert.NotNil(t, err)

	domainKubgoService.AssertNumberOfCalls(t, "Update", 1)
	domainKubgoService.AssertCalled(t, "Update", ct, cqrs.Int(ct.OriginalVersion()))
	domainKubgoService.AssertExpectations(t)

}
