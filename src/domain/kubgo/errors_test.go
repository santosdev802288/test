package kubgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorMessage(t *testing.T) {

	//Arrange
	var exceeds = new(ExceedsCostKubgoError)
	exceeds.Cost = 1

	//Act
	var msg = exceeds.Error()

	//Assert
	assert.NotNil(t, msg)
}
