// Copyright 2020 Siigo. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"siigo.com/kubgo/src/domain/kubgo"
)

func (c KubgoRepository) Save(kubgo *kubgo.Kubgo) chan error {

	result := make(chan error)

	go func() {
		defer close(result)

		_, insertErr := c.collection.InsertOne(context.TODO(), kubgo)

		if insertErr != nil {
			c.logger.Error("InsertOne Kubgo ERROR:", insertErr)
		}

		result <- insertErr
	}()

	return result

}

func (c KubgoRepository) Delete(kubgo *kubgo.Kubgo) chan error {

	result := make(chan error)

	go func() {
		defer close(result)
		_, deleteErr := c.collection.DeleteOne(context.TODO(), kubgo)

		if deleteErr != nil {
			c.logger.Error("DeleteOne Kubgo ERROR:", deleteErr)
		}

		result <- deleteErr
	}()

	return result

}

func (c KubgoRepository) Update(kubgo *kubgo.Kubgo) chan error {

	result := make(chan error)

	go func() {
		defer close(result)

		_, updateErr := c.collection.UpdateOne(context.TODO(),
			bson.M{"_id": kubgo.Id},
			bson.D{{"$set", kubgo}})

		if updateErr != nil {
			c.logger.Error("UpdateOne Kubgo ERROR:", updateErr)
		}

		result <- updateErr
	}()

	return result

}
