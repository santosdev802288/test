package finder

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"siigo.com/kubgo/src/domain/kubgo"
)

// KubgoFinder KubgoRepository provides methods for processing mongo kubgo queries
type KubgoFinder struct {
	collection *mongo.Collection
	context    context.Context
	logger     *logrus.Logger
}

// NewKubgoFinder a new KubgoFinder
func NewKubgoFinder(context context.Context, collection *mongo.Collection, logger *logrus.Logger) kubgo.IKubgoFinder {
	return KubgoFinder{context: context, collection: collection, logger: logger}
}
