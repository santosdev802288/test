package controller

import (
	"context"
	"encoding/json"

	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"siigo.com/kubgo/src/api/logger"
	"siigo.com/kubgo/src/api/proto/kubgo/v1"
	"siigo.com/kubgo/src/application/command"
	"siigo.com/kubgo/src/application/query"
	"siigo.com/kubgo/src/domain/kubgo"
)

func (controller *Controller) FindKubgos(ctx context.Context, empty *emptypb.Empty) (*kubgov1.FindKubgosResponse, error) {
	controller.Logger.Info("Kubgo Find all.")

	// create query message
	em := cqrs.NewQueryMessage(&query.LoadAllKubgoQuery{})
	kubgos, err := controller.Bus.Send(em)
	if err != nil {
		return nil, err
	}

	var cts []*kubgov1.Kubgo
	bytes, _ := json.Marshal(kubgos)
	_ = json.Unmarshal(bytes, &cts)

	return &kubgov1.FindKubgosResponse{Kubgos: cts}, nil
}

func (controller *Controller) AddKubgo(ctx context.Context, grpcKubgo *kubgov1.AddKubgoRequest) (*kubgov1.AddKubgoResponse, error) {

	// generate uuid
	id := cqrs.NewUUID()

	controller.
		Logger.
		WithFields(
			logger.WithBusinessFields(map[string]interface{}{
				"MetricName":  "KubgoCreated",
				"Layer":       "Api",
				"MetricValue": id,
			}),
		).
		Info("Kubgo created .")

	// new aggregate
	ct := kubgo.NewKubgo(id)
	grpcKubgo.Kubgo.Id = id
	grpcKubgo.Kubgo.OccurredAt = ct.OccurredAt

	// map grpc to aggregate
	err := mapper(ct, grpcKubgo.Kubgo)
	if err != nil {
		return nil, err
	}

	// create command message
	em := cqrs.NewCommandMessage(id, &command.CreateKubgoCommand{
		Kubgo: ct,
	})

	// send command to handler(s)
	_, err = controller.Bus.Send(em)

	if err != nil {
		return nil, err
	}

	return &kubgov1.AddKubgoResponse{Kubgo: grpcKubgo.Kubgo}, nil
}

func (controller *Controller) GetKubgo(ctx context.Context, queryKubgo *kubgov1.GetKubgoRequest) (*kubgov1.GetKubgoResponse, error) {

	id, invalidIdError := uuid.FromString(queryKubgo.Id)
	if invalidIdError != nil {
		controller.Logger.Error(invalidIdError)
		return nil, invalidIdError
	}

	controller.Logger.Infof("find kubgo by id %s", id)

	// create query message
	em := cqrs.NewQueryMessage(&query.LoadKubgoQuery{Id: id})

	ct, err := controller.Bus.Send(em)
	if err != nil {
		return nil, err
	}

	// map aggregate to grpc
	response := &kubgov1.Kubgo{}
	err = mapper(response, ct)
	if err != nil {
		return nil, err
	}

	return &kubgov1.GetKubgoResponse{Kubgo: response}, err

}

func (controller *Controller) DeleteKubgo(ctx context.Context, request *kubgov1.DeleteKubgoRequest) (*kubgov1.DeleteKubgoResponse, error) {

	controller.Logger.Infof("find kubgo by id %s", request.Id)

	//Delete ELement
	em := cqrs.NewCommandMessage(request.Id, &command.DeleteKubgoCommand{
		Id: request.Id,
	})

	// send command to handler(s)
	controller.Bus.Send(em)
	return &kubgov1.DeleteKubgoResponse{Kubgo: nil}, nil
}

func (controller *Controller) UpdateKubgo(ctx context.Context, request *kubgov1.UpdateKubgoRequest) (*kubgov1.UpdateKubgoResponse, error) {
	controller.
		Logger.
		WithFields(
			logger.WithBusinessFields(map[string]interface{}{
				"MetricName":  "KubgoUpdate",
				"Layer":       "Api",
				"MetricValue": request.Kubgo.Id,
			}),
		).
		Info("Kubgo Update .")

	// map grpc to aggregate
	ct := kubgo.NewKubgo(request.Kubgo.Id)
	err := mapper(ct, request.Kubgo)
	if err != nil {
		return nil, err
	}

	// create command message
	em := cqrs.NewCommandMessage(request.Kubgo.Id, &command.UpdateKubgoCommand{
		Kubgo: ct,
	})

	// send command to handler(s)
	_, err = controller.Bus.Send(em)

	if err != nil {
		return nil, err
	}

	return &kubgov1.UpdateKubgoResponse{Kubgo: request.Kubgo}, nil

}
