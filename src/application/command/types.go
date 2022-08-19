package command

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"reflect"
	"siigo.com/kubgo/src/domain/kubgo"
	"siigo.com/kubgo/src/domain/services"
)

// CreateKubgoCommand CreateKubgo create a new inventory item
type CreateKubgoCommand struct {
	Kubgo *kubgo.Kubgo
}

type DeleteKubgoCommand struct {
	Id string
}

type UpdateKubgoCommand struct {
	Kubgo *kubgo.Kubgo
}

type KubgoCommandHandler struct {
	kubgoService      services.IKubgoService
	handlersByStructType map[string]func(handler *KubgoCommandHandler, message cqrs.RequestMessage) (interface{}, error)
}

// NewKubgoCommandHandler a new Kubgo Command Handler
func NewKubgoCommandHandler(kubgoService services.IKubgoService) *KubgoCommandHandler {

	// register handler functions
	handlers := map[string]func(handler *KubgoCommandHandler, message cqrs.RequestMessage) (interface{}, error){}
	handlers[reflect.TypeOf(&CreateKubgoCommand{}).String()] = CreateKubgoCommandHandle
	handlers[reflect.TypeOf(&DeleteKubgoCommand{}).String()] = DeleteKubgoCommandHandle
	handlers[reflect.TypeOf(&UpdateKubgoCommand{}).String()] = UpdateKubgoCommandHandle

	return &KubgoCommandHandler{
		kubgoService:      kubgoService,
		handlersByStructType: handlers,
	}
}

type DeleteKubgoHandler struct {
	kubgoService      services.IKubgoService
	handlersByStructType map[string]func(handler *DeleteKubgoHandler, message cqrs.RequestMessage) (interface{}, error)
}
