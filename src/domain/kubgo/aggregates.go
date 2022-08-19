package kubgo

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Kubgo struct {
	Id                 string `bson:"_id"`
	cqrs.AggregateBase `bson:"-"`
	Activated          bool   `json:"activated"`
	Cost               int32  `json:"cost"`
	Address            string `json:"address" validate:"required"`
	OccurredAt         *timestamppb.Timestamp
	Email              string
}

type KubgoResponse struct {
	Error    error
	Kubgo *Kubgo
}

type KubgosResponse struct {
	Error     error
	Kubgos []*Kubgo
}

// NewKubgo constructs a new inventory item aggregate.
//
// Importantly it embeds a new AggregateBase.
func NewKubgo(id string) *Kubgo {
	aggregate := &Kubgo{
		Id:            id,
		AggregateBase: *cqrs.NewAggregateBase(id),
		OccurredAt:    timestamppb.Now(),
	}

	return aggregate
}

func (a *Kubgo) IsWithinBudget() error {
	if a.Cost > 10 {
		return &ExceedsCostKubgoError{Cost: a.Cost}
	}
	return nil
}

// Apply handles the logic of events on the aggregate.
func (a *Kubgo) Apply(message cqrs.EventMessage, isNew bool) {
	a.TrackChange(message)
}
