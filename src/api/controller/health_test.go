// Copyright 2020 Siigo. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"context"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"siigo.com/kubgo/src/api/proto/kubgo/v1"
	"testing"
)

func TestHealth(t *testing.T) {
	// Arrange
	var bus = *new(cqrs.Dispatcher)
	ctl := NewController(bus, logrus.New())
	var expected = kubgov1.HealthResponse{Status: "Ok"}
	n := new(emptypb.Empty)
	ctx := context.Background()

	//Act
	result, err := ctl.Health(ctx, n)

	//Assert
	assert.NotNil(t, ctl)
	assert.NotNil(t, expected)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
