package gateway

import (
	"context"
	"fmt"
	"github.com/felixge/httpsnoop"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"siigo.com/kubgo/src/api/config"
	"siigo.com/kubgo/third_party"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/grpclog"
	"io/fs"
)

// RunServerMux runs the gRPC-Gateway, dialling the provided address.
func RunServerMux(config *config.Configuration, serverMux *runtime.ServeMux) {

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	handler := cors.Default().Handler(serverMux)

	oa := getOpenAPIHandler()

	gatewayAddr := fmt.Sprintf("[::]:%d", config.HttpServer.Port)
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				m := httpsnoop.CaptureMetrics(handler, w, r)
				log.Infof("http[%d]-- %s -- %s\n", m.Code, m.Duration, r.URL.Path)
				return
			}
			oa.ServeHTTP(w, r)
		}),
	}

	go func() {
		log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
		err := gwServer.ListenAndServe()
		panic(err)
	}()

	/*lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
			go gwServer.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stop gRPC-Gateway.")
			return gwServer.Shutdown(ctx)
		},
	})*/

}

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		panic(err.Error())
	}
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

func RegisterHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) func(args ...Handler) {
	return func(args ...Handler) {
		for _, handler := range args {
			err := handler(ctx, mux, conn)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}

type Handler = func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
