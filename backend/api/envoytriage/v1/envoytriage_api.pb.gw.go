// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: envoytriage/v1/envoytriage_api.proto

/*
Package envoytriagev1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package envoytriagev1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_EnvoyTriageAPI_Read_0(ctx context.Context, marshaler runtime.Marshaler, client EnvoyTriageAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ReadRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Read(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_EnvoyTriageAPI_Read_0(ctx context.Context, marshaler runtime.Marshaler, server EnvoyTriageAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ReadRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Read(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterEnvoyTriageAPIHandlerServer registers the http handlers for service EnvoyTriageAPI to "mux".
// UnaryRPC     :call EnvoyTriageAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterEnvoyTriageAPIHandlerFromEndpoint instead.
func RegisterEnvoyTriageAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server EnvoyTriageAPIServer) error {

	mux.Handle("POST", pattern_EnvoyTriageAPI_Read_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/clutch.envoytriage.v1.EnvoyTriageAPI/Read", runtime.WithHTTPPathPattern("/v1/envoytriage/read"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_EnvoyTriageAPI_Read_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EnvoyTriageAPI_Read_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterEnvoyTriageAPIHandlerFromEndpoint is same as RegisterEnvoyTriageAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterEnvoyTriageAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterEnvoyTriageAPIHandler(ctx, mux, conn)
}

// RegisterEnvoyTriageAPIHandler registers the http handlers for service EnvoyTriageAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterEnvoyTriageAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterEnvoyTriageAPIHandlerClient(ctx, mux, NewEnvoyTriageAPIClient(conn))
}

// RegisterEnvoyTriageAPIHandlerClient registers the http handlers for service EnvoyTriageAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "EnvoyTriageAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "EnvoyTriageAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "EnvoyTriageAPIClient" to call the correct interceptors.
func RegisterEnvoyTriageAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client EnvoyTriageAPIClient) error {

	mux.Handle("POST", pattern_EnvoyTriageAPI_Read_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/clutch.envoytriage.v1.EnvoyTriageAPI/Read", runtime.WithHTTPPathPattern("/v1/envoytriage/read"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_EnvoyTriageAPI_Read_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EnvoyTriageAPI_Read_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_EnvoyTriageAPI_Read_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "envoytriage", "read"}, ""))
)

var (
	forward_EnvoyTriageAPI_Read_0 = runtime.ForwardResponseMessage
)
