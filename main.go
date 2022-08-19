package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"log"
	"os"
	"siigo.com/kubgo/src/api/boot"
	"siigo.com/kubgo/src/api/config"
	"siigo.com/kubgo/src/api/controller"
	fxmodule2 "siigo.com/kubgo/src/api/fxmodule"
	"siigo.com/kubgo/src/api/logger"
)

func init() {

	dd := os.Getenv("DD_AGENT_HOST")
	if len(dd) == 0 {
		return
	}

	// DataDog
	tracer.Start()
	defer tracer.Stop()
	err := profiler.Start(
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()
}
func main() {
	// start app
	newFxApp().Run()
}

func newFxApp() *fx.App {

	return fx.New(

		// Create Struct FX Providers
		fx.Provide(
			config.NewConfiguration,
			logger.NewLogrus,
		),

		fx.Provide(
			boot.CreateGrpcServer,
			boot.CreateGrpcClient,
			boot.NewNetListener,
			runtime.NewServeMux,
			controller.NewController,
		),

		// Load Module in bottom order
		fxmodule2.CQRSDDDModule,
		fxmodule2.BrokerModule,
		fxmodule2.InfrastructureModule,
		fxmodule2.ApplicationModule,

		// Invoke to init functions to start
		fx.Invoke(
			boot.RegisterGrpcServers,
			boot.StartGrpcServer,
			boot.RegisterGrpcHandlers,
			boot.StartHttpServer,
		),
	)
}
