package services

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	finder "siigo.com/kubgo/mocks/src/domain/kubgo"
	repository "siigo.com/kubgo/mocks/src/domain/kubgo"
	"siigo.com/kubgo/src/domain/kubgo"
	"testing"
	"time"
)

func TestNewKubgoService(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	repository := &repository.IKubgoRepository{}
	finder := &finder.IKubgoFinder{}

	id := cqrs.NewUUID()
	kubgoMock := kubgo.NewKubgo(id)

	repository.
		On("Save", mock.Anything).
		Return(nil, nil)

	finder.
		On("Get", mock.Anything).
		Return(kubgoMock, nil)

	//Act
	var inter = NewKubgoService(dispatcher, nil, repository, finder)

	//Assert
	assert.NotNil(t, inter)
	assert.NotNil(t, repository)
	assert.NotNil(t, finder)

}

func TestNewKubgoDispatcherNil(t *testing.T) {
	// Arrange

	repository := &repository.IKubgoRepository{}
	finder := &finder.IKubgoFinder{}

	//Act
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("TestUserFail should have panicked!")
			}
		}()
		// This function should cause a panic
		NewKubgoService(nil, nil, repository, finder)
	}()
}

func TestNewKubgoServiceGet(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	responsemock := make(chan *kubgo.KubgoResponse)
	defer close(responsemock)

	id := cqrs.NewUUID()
	uid, _ := uuid.FromString(id)
	//kubgoMock := kubgo.NewKubgo(id)

	finderMock.On("Get", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, nil, finderMock)
	contracresponse := &kubgo.KubgoResponse{Kubgo: kubgo.NewKubgo(id), Error: nil}

	go func() {
		time.Sleep(time.Second)
		responsemock <- contracresponse
	}()

	var response, error = service.Get(uid)

	//Assert
	assert.NotNil(t, response)
	assert.Nil(t, error)
	finderMock.AssertNumberOfCalls(t, "Get", 1)
	finderMock.AssertCalled(t, "Get", uid)
	finderMock.AssertExpectations(t)
}

func TestNewKubgoServiceLoadAll(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	responsemock := make(chan kubgo.KubgosResponse)
	defer close(responsemock)

	//kubgoMock := kubgo.NewKubgo(id)

	finderMock.On("GetAll", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, nil, finderMock)
	contracresponse := kubgo.KubgosResponse{}

	go func() {
		time.Sleep(time.Second)
		responsemock <- contracresponse
	}()

	var response, error = service.LoadAll()

	//Assert
	assert.Nil(t, response)
	assert.Nil(t, error)
}

func TestNewKubgoServiceLoadAllError(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	responsemock := make(chan kubgo.KubgosResponse)
	defer close(responsemock)

	//kubgoMock := kubgo.NewKubgo(id)

	finderMock.On("GetAll", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, nil, finderMock)
	contracresponse := kubgo.KubgosResponse{Error: errors.New("kubgo not found")}

	go func() {
		time.Sleep(time.Second)
		responsemock <- contracresponse
	}()

	var response, error = service.LoadAll()

	//Assert
	assert.Nil(t, response)
	assert.NotNil(t, error)
}

func TestNewKubgoServiceLoadErrorResponse(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	responsemock := make(chan *kubgo.KubgoResponse)
	defer close(responsemock)

	id := cqrs.NewUUID()
	uid, _ := uuid.FromString(id)
	//kubgoMock := kubgo.NewKubgo(id)

	finderMock.On("Get", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, nil, finderMock)

	go func() {
		time.Sleep(time.Second)
		responsemock <- &kubgo.KubgoResponse{Error: errors.New("")}
	}()

	var response, error = service.Get(uid)

	//Assert
	assert.Nil(t, response)
	assert.NotNil(t, error)
	finderMock.AssertNumberOfCalls(t, "Get", 1)
	finderMock.AssertCalled(t, "Get", uid)
	finderMock.AssertExpectations(t)
}

func TestNewKubgoServiceSave(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)
	ct := &kubgo.Kubgo{}

	repositoryMock.On("Save", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, nil)

	go func() {
		time.Sleep(time.Second)
		responsemock <- nil
	}()

	var result = service.Save(ct, nil)

	//Assert
	assert.Nil(t, result)
}

func TestNewKubgoServiceSaveErrorGeneral(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)
	ct := &kubgo.Kubgo{}

	repositoryMock.On("Save", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, nil)

	go func() {
		time.Sleep(time.Second)
		responsemock <- errors.New("Fail General")
	}()

	var result = service.Save(ct, nil)

	//Assert
	assert.NotNil(t, result)
}

func TestNewKubgoServiceErrorIsWithinBudget(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	repositoryMock := &repository.IKubgoRepository{}

	ct := &kubgo.Kubgo{
		Cost: 20,
	}

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, nil)

	var result = service.Save(ct, nil)

	//Assert
	assert.NotNil(t, result)
	assert.Equal(t, result.Error(), "rpc error: code = InvalidArgument desc = cost 20 exceeds the budget.")
}

func TestNewKubgoServiceDelete(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)

	responsemockLoad := make(chan *kubgo.KubgoResponse)
	defer close(responsemockLoad)

	id := cqrs.NewUUID()

	finderMock.On("Get", mock.Anything).
		Return(responsemockLoad)
	contracresponseLoad := &kubgo.KubgoResponse{Kubgo: kubgo.NewKubgo(id), Error: nil}

	repositoryMock.On("Delete", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	go func() {
		time.Sleep(time.Second)
		responsemockLoad <- contracresponseLoad
		responsemock <- nil
	}()

	var error = service.Delete(id)

	//Assert
	assert.Nil(t, error)
}

func TestNewKubgoServiceDeleteErrorId(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	var error = service.Delete("id")

	//Assert
	assert.NotNil(t, error)
}

func TestNewKubgoServiceDeleteContracNotFound(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)

	responsemockLoad := make(chan *kubgo.KubgoResponse)
	defer close(responsemockLoad)

	id := cqrs.NewUUID()

	finderMock.On("Get", mock.Anything).
		Return(responsemockLoad)
	contracresponseLoad := &kubgo.KubgoResponse{Kubgo: kubgo.NewKubgo(id), Error: errors.New("kubgo not found")}

	repositoryMock.On("Delete", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	go func() {
		time.Sleep(time.Second)
		responsemockLoad <- contracresponseLoad
	}()

	var error = service.Delete(id)

	//Assert
	assert.NotNil(t, error)
}

func TestNewKubgoServiceDeleteErrorGeneral(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)

	responsemockLoad := make(chan *kubgo.KubgoResponse)
	defer close(responsemockLoad)

	id := cqrs.NewUUID()

	finderMock.On("Get", mock.Anything).
		Return(responsemockLoad)
	contracresponseLoad := &kubgo.KubgoResponse{Kubgo: kubgo.NewKubgo(id), Error: nil}

	repositoryMock.On("Delete", mock.Anything).
		Return(responsemock)

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	go func() {
		time.Sleep(time.Second)
		responsemockLoad <- contracresponseLoad
		responsemock <- errors.New("Fail General")
	}()

	var error = service.Delete(id)

	//Assert
	assert.NotNil(t, error)
}

func TestNewKubgoServiceUpdate(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)
	ct := &kubgo.Kubgo{}
	id := cqrs.NewUUID()
	ct.Id = id

	responsemockLoad := make(chan *kubgo.KubgoResponse)
	defer close(responsemockLoad)

	repositoryMock.On("Update", mock.Anything).
		Return(responsemock)

	finderMock.On("Get", mock.Anything).
		Return(responsemockLoad)
	contracresponseLoad := &kubgo.KubgoResponse{Kubgo: kubgo.NewKubgo(id), Error: nil}

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	go func() {
		time.Sleep(time.Second)
		responsemockLoad <- contracresponseLoad
		responsemock <- nil
	}()

	var error = service.Update(ct, nil)

	//Assert
	assert.Nil(t, error)
}

func TestNewKubgoServiceUpdateIsWithinBudget(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)
	ct := &kubgo.Kubgo{}
	id := cqrs.NewUUID()
	ct.Id = id
	ct.Cost = 20

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	var error = service.Update(ct, nil)

	//Assert
	assert.NotNil(t, error)
}

func TestNewKubgoServiceUpdateKubgoNotFound(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)
	ct := &kubgo.Kubgo{}
	id := cqrs.NewUUID()
	ct.Id = id

	responsemockLoad := make(chan *kubgo.KubgoResponse)
	defer close(responsemockLoad)

	repositoryMock.On("Update", mock.Anything).
		Return(responsemock)

	finderMock.On("Get", mock.Anything).
		Return(responsemockLoad)
	contracresponseLoad := &kubgo.KubgoResponse{Kubgo: kubgo.NewKubgo(id), Error: errors.New("kubgo not found")}

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	go func() {
		time.Sleep(time.Second)
		responsemockLoad <- contracresponseLoad
	}()

	var error = service.Update(ct, nil)

	//Assert
	assert.NotNil(t, error)
}

func TestNewKubgoServiceUpdateKubgoErrorTransaction(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)
	ct := &kubgo.Kubgo{}
	id := cqrs.NewUUID()
	ct.Id = id

	responsemockLoad := make(chan *kubgo.KubgoResponse)
	defer close(responsemockLoad)

	repositoryMock.On("Update", mock.Anything).
		Return(responsemock)

	finderMock.On("Get", mock.Anything).
		Return(responsemockLoad)
	contracresponseLoad := &kubgo.KubgoResponse{Kubgo: kubgo.NewKubgo(id), Error: nil}

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	go func() {
		time.Sleep(time.Second)
		responsemockLoad <- contracresponseLoad
		responsemock <- errors.New("Fail General")
	}()

	var error = service.Update(ct, nil)

	//Assert
	assert.NotNil(t, error)
}

func TestNewKubgoServiceUpdateErrorId(t *testing.T) {
	// Arrange
	dispatcher := &cqrs.DomainEventsDispatcher{}
	finderMock := &finder.IKubgoFinder{}
	repositoryMock := &repository.IKubgoRepository{}
	responsemock := make(chan error)
	defer close(responsemock)
	ct := &kubgo.Kubgo{}
	id := cqrs.NewUUID()

	responsemockLoad := make(chan *kubgo.KubgoResponse)
	defer close(responsemockLoad)

	repositoryMock.On("Update", mock.Anything).
		Return(responsemock)

	finderMock.On("Get", mock.Anything).
		Return(responsemockLoad)
	contracresponseLoad := &kubgo.KubgoResponse{Kubgo: kubgo.NewKubgo(id), Error: nil}

	//Act
	var service = NewKubgoService(dispatcher, nil, repositoryMock, finderMock)

	go func() {
		time.Sleep(time.Second)
		responsemockLoad <- contracresponseLoad
		responsemock <- nil
	}()

	var error = service.Update(ct, nil)

	//Assert
	assert.NotNil(t, error)
}
