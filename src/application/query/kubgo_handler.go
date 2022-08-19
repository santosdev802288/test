package query

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"errors"
	"reflect"
)

// Handle processes kubgo commands.
func (handler *KubgoQueryHandler) Handle(message cqrs.RequestMessage) (interface{}, error) {
	fn, ok := handler.handlersByStructType[reflect.TypeOf(message.Request()).String()]
	if !ok {
		return nil, errors.New("query type font found")
	}
	return fn(handler, message)
}

func LoadAllKubgoQueryHandle(handler *KubgoQueryHandler, message cqrs.RequestMessage) (interface{}, error) {
	kubgoResult, err := handler.kubgoService.LoadAll()
	if err != nil {
		return nil, err
	}
	return kubgoResult, nil

}

func LoadKubgoQueryHandle(handler *KubgoQueryHandler, message cqrs.RequestMessage) (interface{}, error) {

	cmd := message.Request().(*LoadKubgoQuery)

	kubgoResult, err := handler.kubgoService.Get(cmd.Id)
	if err != nil {
		return nil, err
	}

	return kubgoResult, nil

}
