package command

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"errors"
	"reflect"
)

// Handle processes kubgo commands.
func (handler *KubgoCommandHandler) Handle(message cqrs.RequestMessage) (interface{}, error) {
	fn, ok := handler.handlersByStructType[reflect.TypeOf(message.Request()).String()]
	if !ok {
		return nil, errors.New("command type font found")
	}
	return fn(handler, message)
}

func CreateKubgoCommandHandle(handler *KubgoCommandHandler, message cqrs.RequestMessage) (interface{}, error) {

	cmd := message.Request().(*CreateKubgoCommand)

	err := handler.kubgoService.Save(cmd.Kubgo, cqrs.Int(cmd.Kubgo.OriginalVersion()))
	if err != nil {
		return nil, err
	}

	return cmd.Kubgo, nil

}

func DeleteKubgoCommandHandle(handler *KubgoCommandHandler, message cqrs.RequestMessage) (interface{}, error) {

	cmd := message.Request().(*DeleteKubgoCommand)

	err := handler.kubgoService.Delete(cmd.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func UpdateKubgoCommandHandle(handler *KubgoCommandHandler, message cqrs.RequestMessage) (interface{}, error) {

	cmd := message.Request().(*UpdateKubgoCommand)

	err := handler.kubgoService.Update(cmd.Kubgo, cqrs.Int(cmd.Kubgo.OriginalVersion()))
	if err != nil {
		return nil, err
	}

	return nil, nil

}
