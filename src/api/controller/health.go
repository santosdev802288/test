// Copyright 2020 Siigo. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"siigo.com/kubgo/src/api/proto/kubgo/v1"
)

func (controller *Controller) Health(ctx context.Context, r *emptypb.Empty) (*kubgov1.HealthResponse, error) {
	return &kubgov1.HealthResponse{Status: "Ok"}, nil
}
