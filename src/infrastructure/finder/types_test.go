package finder

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestNewKubgoFinder(t *testing.T) {
	// Arrange
	mongo := new(*mongo.Collection)
	//Act
	finderKubgo := NewKubgoFinder(context.Background(), *mongo, logrus.New())
	//Assert
	assert.NotNil(t, finderKubgo)
}
