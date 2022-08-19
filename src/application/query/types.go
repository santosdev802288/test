package query

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"reflect"
	"siigo.com/kubgo/src/domain/services"
)

//LoadAllKubgo
type LoadAllKubgoQuery struct {
}

// LoadKubgoQuery create a new kubgo item
type LoadKubgoQuery struct {
	Id uuid.UUID
}

type KubgoQueryHandler struct {
	kubgoService      services.IKubgoService
	handlersByStructType map[string]func(handler *KubgoQueryHandler, message cqrs.RequestMessage) (interface{}, error)
}

// NewKubgoQueryHandler a new Kubgo query Handler
func NewKubgoQueryHandler(kubgoService services.IKubgoService) *KubgoQueryHandler {

	// register handler functions
	handlers := map[string]func(handler *KubgoQueryHandler, message cqrs.RequestMessage) (interface{}, error){}
	handlers[reflect.TypeOf(&LoadKubgoQuery{}).String()] = LoadKubgoQueryHandle
	handlers[reflect.TypeOf(&LoadAllKubgoQuery{}).String()] = LoadAllKubgoQueryHandle

	return &KubgoQueryHandler{
		kubgoService:      kubgoService,
		handlersByStructType: handlers,
	}
}
