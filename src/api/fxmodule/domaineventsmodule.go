package fxmodule

import (
	"fmt"

	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	slim "dev.azure.com/SiigoDevOps/Siigo/_git/go-slim.git/slim"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	goes "github.com/jetbasrawi/go.geteventstore"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"siigo.com/kubgo/src/api/config"
	"siigo.com/kubgo/src/domain/services"
)

var validate = validator.New()

var CQRSDDDModule = fx.Options(
	fx.Provide(
		cqrs.NewInternalEventBus,
		NewEventStoreClient,
		NewEventStoreCommonDomainRepo,
		cqrs.NewInMemoryRepo,
		NewDomainEventsDispatcher,
		NewInMemoryDispatcherWithInterceptors,
		NewRedisClient,

		services.NewKubgoService,
	),
)

func NewInMemoryDispatcherWithInterceptors(eventDispatcher *cqrs.DomainEventsDispatcher) cqrs.Dispatcher {
	return cqrs.NewInMemoryDispatcher(
		eventDispatcher,
	)
}

func NewEventStoreClient(config *config.Configuration) *goes.Client {
	client, err := goes.NewClient(nil, config.EventStore.Url)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewEventStoreCommonDomainRepo(eventStore *goes.Client, eventBus *cqrs.InternalEventBus) *cqrs.GetEventStoreCommonDomainRepo {
	r, err := cqrs.NewCommonDomainRepository(eventStore, eventBus)
	if err != nil {
		panic(err)
	}
	return r
}

func NewDomainEventsDispatcher(eventStore *cqrs.GetEventStoreCommonDomainRepo, slim *slim.MessageBusBuilder) *cqrs.DomainEventsDispatcher {
	return &cqrs.DomainEventsDispatcher{
		EventStore: eventStore,
		EventPublisher: func(event interface{}) {

			if event == nil {
				return
			}

			slim.Publish(event)

		},
	}
}

// DispatcherValidationInterceptor validate commands and queries
func DispatcherValidationInterceptor(message cqrs.RequestMessage) error {

	// struct validations by tags https://github.com/go-playground/validator
	err := validate.Struct(message.Request())

	// validation ok
	if err == nil {
		return nil
	}

	// map errors to grpc-errors
	validationErrors := err.(validator.ValidationErrors)
	var details []proto.Message

	// map validation errors to proto message details
	for i := range validationErrors {
		detailError := &errdetails.ErrorInfo{Domain: "Api", Reason: validationErrors[i].Error()}
		details = append(details, detailError)
	}

	// create new grpc error with status and details
	stat, e := status.
		New(codes.InvalidArgument, "Invalid request").
		WithDetails(details...)

	if e != nil {
		return status.Errorf(codes.Internal, "unexpected error adding details: %s", e)
	}

	return stat.Err()

}

// NewRedisClient Create Connection redis
func NewRedisClient(config *config.Configuration) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:        config.Redis.Addr,
		Password:    config.Redis.Password,
		DB:          config.Redis.Db,
		ReadTimeout: config.Redis.TimeOut,
	})
	pong, err := client.Ping().Result()
	fmt.Println("Status Redis "+pong, err)
	return client
}
