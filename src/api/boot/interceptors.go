package boot

import (
	"context"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// RequestIDInterceptor is a interceptor of access control list.
func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestID := requestIDFromContext(ctx)

		header := metadata.Pairs(XRequestIDKey, requestID)
		grpc.SendHeader(ctx, header)

		ctx = context.WithValue(ctx, XRequestIDKey, requestID)
		return handler(ctx, req)
	}
}

//
func requestIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return unknownRequestID
	}

	header, ok := md[XRequestIDKey]
	if !ok || len(header) == 0 {
		//  generate request id if not exist
		return uuid.NewUUID()
	}

	requestID := header[0]
	if requestID == "" {
		return unknownRequestID
	}

	return requestID
}

// ProtocValidationInterceptor validate protobuf definition rules
// https://github.com/envoyproxy/protoc-gen-validate
func ProtocValidationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	ptr, itsVerifiable := req.(interface{ Validate() error })
	if !itsVerifiable {
		return handler(ctx, req)
	}

	validationError := ptr.Validate()
	if validationError == nil {
		return handler(ctx, req)
	}

	stat, e := status.
		New(codes.InvalidArgument, "Invalid request").
		WithDetails(&errdetails.ErrorInfo{Domain: "Api", Reason: validationError.Error()})

	if e != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error adding details: %s", e)
	}

	return nil, stat.Err()

}
