package services

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"encoding/json"
	"github.com/go-redis/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"siigo.com/kubgo/src/domain/kubgo"
)

// KubgoService is a services specialized for persistence events of Kubgos and invoke infrastructure implementation.
// While it is not required to construct a services specialized for a
// specific aggregate type, it is better to do so. There can be quite a lot of
// services configuration that is specific to a type and it is cleaner if that
// code is contained in a specialized repository as shown here.
// Also because the CommonDomainRepository Load method returns an interface{}, a
// type assertion is required. Here the type assertion is contained in this specialized
// eventStore and a *Kubgo is returned from the eventStore.
type KubgoService struct {
	events         *cqrs.DomainEventsDispatcher
	redis          *redis.Client
	kubgoRepo   kubgo.IKubgoRepository
	kubgoFinder kubgo.IKubgoFinder
}

type IKubgoService interface {
	LoadAll() ([]*kubgo.Kubgo, error)
	Get(id uuid.UUID) (*kubgo.Kubgo, error)
	Save(aggregate cqrs.AggregateRoot, expectedVersion *int) error
	Delete(id string) error
	Update(aggregate cqrs.AggregateRoot, expectedVersion *int) error
}

// NewKubgoService constructs a new NewKubgoService.
func NewKubgoService(eventStore *cqrs.DomainEventsDispatcher, redisCliente *redis.Client, kubgoRepo kubgo.IKubgoRepository, kubgoFinder kubgo.IKubgoFinder) IKubgoService {

	ret := &KubgoService{
		events:         eventStore,
		redis:          redisCliente,
		kubgoRepo:   kubgoRepo,
		kubgoFinder: kubgoFinder,
	}

	// An aggregate factory creates an aggregate instance given the name of an aggregate.
	aggregateFactory := cqrs.NewDelegateAggregateFactory()
	err := aggregateFactory.RegisterDelegate(&kubgo.Kubgo{}, func(id string) cqrs.AggregateRoot { return kubgo.NewKubgo(id) })
	if err != nil {
		panic(err)
	}
	if ret.events.EventStore != nil {
		ret.events.EventStore.SetAggregateFactory(aggregateFactory)
	}

	// A stream name delegate constructs a stream name.
	// A common way to construct a stream name is to use a bounded context and
	// an aggregate id.
	// The interface for a stream name delegate takes a two strings. One may be
	// the aggregate type and the other the aggregate id. In this case the first
	// argument and the second argument are concatenated with a hyphen.
	streamNameDelegate := cqrs.NewDelegateStreamNamer()
	err = streamNameDelegate.RegisterDelegate(func(t string, id string) string {
		return t + "-" + id
	}, &kubgo.Kubgo{})
	if err != nil {
		panic(err)
	}
	if ret.events.EventStore != nil {
		ret.events.EventStore.SetStreamNameDelegate(streamNameDelegate)
	}

	// An event factory creates an instance of an event given the name of an event
	// as a string.
	eventFactory := cqrs.NewDelegateEventFactory()
	err = eventFactory.RegisterDelegate(&kubgo.CreatedEvent{}, func() interface{} { return &kubgo.CreatedEvent{} })
	if err != nil {
		panic(err)
	}

	err = eventFactory.RegisterDelegate(&kubgo.UpdatedEvent{}, func() interface{} { return &kubgo.UpdatedEvent{} })
	if err != nil {
		panic(err)
	}

	if ret.events.EventStore != nil {
		ret.events.EventStore.SetEventFactory(eventFactory)
	}

	return ret
}

// LoadAll Load
func (r *KubgoService) LoadAll() ([]*kubgo.Kubgo, error) {

	kubgoResponse := <-r.kubgoFinder.GetAll()

	if kubgoResponse.Error != nil {
		return nil, kubgoResponse.Error
	}

	return kubgoResponse.Kubgos, nil
}

// Get Load Returns an *Aggregate.
func (r *KubgoService) Get(id uuid.UUID) (*kubgo.Kubgo, error) {

	//Evaluate redis properties
	if r.redis != nil {
		uid := id.String()
		val, err := r.redis.Get(uid).Result()
		if err == nil {
			subscribers := kubgo.KubgoResponse{}
			errorReflect := json.Unmarshal([]byte(val), &subscribers)
			if errorReflect == nil && subscribers.Kubgo != nil {
				return subscribers.Kubgo, nil
			}
		}

	}

	// load document by bson id
	kubgoResponse := <-r.kubgoFinder.Get(id)

	if kubgoResponse.Error != nil {
		return nil, kubgoResponse.Error
	}

	//Put item redis
	json, err := json.Marshal(kubgoResponse)
	if err == nil {
		if r.redis != nil {
			go r.redis.Set(kubgoResponse.Kubgo.Id, json, 0).Err()
		}
	}

	return kubgoResponse.Kubgo, nil
}

// Save persists an aggregate.
func (r *KubgoService) Save(aggregate cqrs.AggregateRoot, expectedVersion *int) error {

	ctr := aggregate.(*kubgo.Kubgo)

	// validate cost of kubgo
	if err := ctr.IsWithinBudget(); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	// use infrastructure
	errno := <-r.kubgoRepo.Save(ctr)
	if errno != nil {
		return status.Error(codes.FailedPrecondition, errno.Error())
	}

	// add domain event
	ctr.Apply(kubgo.NewKubgoCreatedEvent(aggregate), true)

	// dispatch domain events
	if r.events.EventStore != nil {
		go r.events.SaveAndPublish(ctr, nil)
	}

	return nil
}

// Delete element.
func (r *KubgoService) Delete(Id string) error {

	id, invalidIdError := uuid.FromString(Id)
	if invalidIdError != nil {
		return invalidIdError
	}

	kubgoDocument, err := r.Get(id)
	if err != nil {
		return err
	}

	// use infrastructure
	errno := <-r.kubgoRepo.Delete(kubgoDocument)
	if errno != nil {
		return errno
	}

	//Delete reference redis
	if r.redis != nil {
		go r.redis.Del(Id)
	}

	return nil
}

// Update element.
func (r *KubgoService) Update(aggregate cqrs.AggregateRoot, expectedVersion *int) error {

	ctr := aggregate.(*kubgo.Kubgo)

	// validate cost of kubgo
	if err := ctr.IsWithinBudget(); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	//Validate item exist
	id, invalidIdError := uuid.FromString(ctr.Id)
	if invalidIdError != nil {
		return invalidIdError
	}
	_, err := r.Get(id)
	if err != nil {
		return status.Error(codes.NotFound, "kubgo not found")
	}

	// use infrastructure
	errno := <-r.kubgoRepo.Update(ctr)
	if errno != nil {
		return errno
	}

	// add domain events
	ctr.Apply(kubgo.NewKubgoUpdateEvent(aggregate), false)

	// dispatch domain events
	if r.events.EventStore != nil {
		go r.events.SaveAndPublish(ctr, nil)
	}

	//Delete reference redis
	if r.redis != nil {
		go r.redis.Del(ctr.Id)
	}
	return nil
}
