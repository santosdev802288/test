package finder

import (
	"context"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"siigo.com/kubgo/src/domain/kubgo"
)

func (c KubgoFinder) GetAll() chan kubgo.KubgosResponse {

	result := make(chan kubgo.KubgosResponse)
	kubgosResponse := kubgo.KubgosResponse{}

	go func() {
		defer close(result)

		cursor, err := c.collection.Find(context.Background(), bson.D{{}})
		if err != nil {
			c.logger.Error("[Mongo] Error find kubgos:", err)
			kubgosResponse.Error = err
		} else {
			defer cursor.Close(context.Background())
			var kubgos []*kubgo.Kubgo

			// iterate for each kubgo
			for cursor.Next(context.Background()) {
				var ct *kubgo.Kubgo
				err := cursor.Decode(&ct)
				if err != nil {
					c.logger.Error("[Mongo] Error decoding kubgos:", err)
					kubgosResponse.Error = err
					result <- kubgosResponse
				} else {
					kubgos = append(kubgos, ct)
				}
			}

			kubgosResponse.Kubgos = kubgos
		}

		result <- kubgosResponse

	}()

	return result
}

func (c KubgoFinder) Get(id uuid.UUID) chan *kubgo.KubgoResponse {

	result := make(chan *kubgo.KubgoResponse)
	response := &kubgo.KubgoResponse{}

	go func() {
		defer close(result)

		kubgoResponse := &kubgo.Kubgo{}

		if findError := c.collection.
			FindOne(context.Background(), bson.D{{"_id", id.String()}}).
			Decode(kubgoResponse); findError != nil {

			c.logger.Error(findError)

			if errors.Is(findError, mongo.ErrNoDocuments) {
				response.Error = errors.New("kubgo not found")
			} else {
				response.Error = findError
			}
		}

		response.Kubgo = kubgoResponse
		result <- response
	}()

	return result

}
