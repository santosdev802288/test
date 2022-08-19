package kubgo

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKubgoCreated(t *testing.T) {

	//Act
	var kubgoEvent = &CreatedEvent{Kubgo: NewKubgo(cqrs.NewUUID())}

	//Asserts
	assert.NotNil(t, kubgoEvent.Kubgo.Id)
	assert.NotNil(t, kubgoEvent.Kubgo.OccurredAt)

}

func TestKubgoEvent(t *testing.T) {

	//Act
	var kubgoEvent = &UpdatedEvent{Kubgo: NewKubgo(cqrs.NewUUID())}

	//Asserts
	assert.NotNil(t, kubgoEvent.Kubgo.Id)
	assert.NotNil(t, kubgoEvent.Kubgo.OccurredAt)
}
