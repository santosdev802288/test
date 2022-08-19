package kubgo

import "dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"

// Events are just plain structs

// CreatedEvent KubgoCreatedEvent KubgoCreated event
type CreatedEvent struct {
	Kubgo *Kubgo
}

// UpdatedEvent CreatedEvent KubgoCreatedEvent KubgoCreated event
type UpdatedEvent struct {
	Kubgo *Kubgo
}

func NewKubgoCreatedEvent(aggregate cqrs.AggregateRoot) *cqrs.EventDescriptor {
	return cqrs.NewEventMessage(
		aggregate.AggregateID(),
		&CreatedEvent{Kubgo: aggregate.(*Kubgo)},
		cqrs.Int(aggregate.CurrentVersion()),
	)
}

func NewKubgoUpdateEvent(aggregate cqrs.AggregateRoot) *cqrs.EventDescriptor {
	return cqrs.NewEventMessage(
		aggregate.AggregateID(),
		&UpdatedEvent{Kubgo: aggregate.(*Kubgo)},
		cqrs.Int(aggregate.CurrentVersion()),
	)
}
