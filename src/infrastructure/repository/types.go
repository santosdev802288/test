package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"siigo.com/kubgo/src/domain/kubgo"
)

// KubgoRepository provides methods for processing mongo Person Connections
type KubgoRepository struct {
	collection *mongo.Collection
	context    context.Context
	logger     *logrus.Logger
}

// NewKubgoRepository a new NewKubgoRepository
func NewKubgoRepository(context context.Context, collection *mongo.Collection, logger *logrus.Logger) kubgo.IKubgoRepository {
	return KubgoRepository{context: context, collection: collection, logger: logger}
}
