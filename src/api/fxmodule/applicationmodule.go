// Copyright 2020 Siigo. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package fxmodule

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs"
	"go.uber.org/fx"
	"siigo.com/kubgo/src/application/command"
	"siigo.com/kubgo/src/application/query"
)

// ApplicationModule Application Module Commands And Queries
var ApplicationModule = fx.Options(
	fx.Provide(
		command.NewKubgoCommandHandler,
		query.NewKubgoQueryHandler,
	),
	fx.Invoke(
		RegisterHandlers,
	),
)

// RegisterHandlers Register Commands and Queries
func RegisterHandlers(dispatcher cqrs.Dispatcher,
	kubgoCommandHandler *command.KubgoCommandHandler,
	kubgoQueryHandler *query.KubgoQueryHandler,
) {

	// Configure Commands
	HandleErrorRegister(
		dispatcher.RegisterHandler(kubgoCommandHandler, &command.CreateKubgoCommand{}),
		dispatcher.RegisterHandler(kubgoCommandHandler, &command.DeleteKubgoCommand{}),
		dispatcher.RegisterHandler(kubgoCommandHandler, &command.UpdateKubgoCommand{}),
	)

	// Configure Queries
	HandleErrorRegister(
		dispatcher.RegisterHandler(kubgoQueryHandler, &query.LoadKubgoQuery{}),
		dispatcher.RegisterHandler(kubgoQueryHandler, &query.LoadAllKubgoQuery{}),
	)

}

func HandleErrorRegister(errors ...error) {
	for _, err := range errors {
		if err != nil {
			panic(err)
		}
	}
}
