package kubgo

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Arrange
	id := cqrs.NewUUID()
	var kubgo = NewKubgo(id)

	//Act

	//Assert
	assert.NotNil(t, kubgo)
}

func TestIsWithinBudget(t *testing.T) {
	// Arrange
	id := cqrs.NewUUID()
	var kubgo = NewKubgo(id)
	kubgo.Cost = 10
	//Act
	kubgo.IsWithinBudget()
	//Assert
	assert.NotNil(t, kubgo)
}

func TestIsWithinBudgetExceedsCostKubgoError(t *testing.T) {
	// Arrange
	id := cqrs.NewUUID()
	var kubgo = NewKubgo(id)
	kubgo.Cost = 20
	//Act
	error := kubgo.IsWithinBudget()
	//Assert
	assert.NotNil(t, kubgo)
	assert.NotNil(t, error)
}
