// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package talent

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newEventClientHook clientHook

// EventCallOptions contains the retry settings for each method of EventClient.
type EventCallOptions struct {
	CreateClientEvent []gax.CallOption
}

func defaultEventGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("jobs.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("jobs.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://jobs.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultEventCallOptions() *EventCallOptions {
	return &EventCallOptions{
		CreateClientEvent: []gax.CallOption{},
	}
}

// internalEventClient is an interface that defines the methods availaible from Cloud Talent Solution API.
type internalEventClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateClientEvent(context.Context, *talentpb.CreateClientEventRequest, ...gax.CallOption) (*talentpb.ClientEvent, error)
}

// EventClient is a client for interacting with Cloud Talent Solution API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service handles client event report.
type EventClient struct {
	// The internal transport-dependent client.
	internalClient internalEventClient

	// The call options for this service.
	CallOptions *EventCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *EventClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *EventClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *EventClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateClientEvent report events issued when end user interacts with customer’s application
// that uses Cloud Talent Solution. You may inspect the created events in
// self service
// tools (at https://console.cloud.google.com/talent-solution/overview).
// Learn
// more (at https://cloud.google.com/talent-solution/docs/management-tools)
// about self service tools.
func (c *EventClient) CreateClientEvent(ctx context.Context, req *talentpb.CreateClientEventRequest, opts ...gax.CallOption) (*talentpb.ClientEvent, error) {
	return c.internalClient.CreateClientEvent(ctx, req, opts...)
}

// eventGRPCClient is a client for interacting with Cloud Talent Solution API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type eventGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing EventClient
	CallOptions **EventCallOptions

	// The gRPC API client.
	eventClient talentpb.EventServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewEventClient creates a new event service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service handles client event report.
func NewEventClient(ctx context.Context, opts ...option.ClientOption) (*EventClient, error) {
	clientOpts := defaultEventGRPCClientOptions()
	if newEventClientHook != nil {
		hookOpts, err := newEventClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := EventClient{CallOptions: defaultEventCallOptions()}

	c := &eventGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		eventClient:      talentpb.NewEventServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *eventGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *eventGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *eventGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *eventGRPCClient) CreateClientEvent(ctx context.Context, req *talentpb.CreateClientEventRequest, opts ...gax.CallOption) (*talentpb.ClientEvent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateClientEvent[0:len((*c.CallOptions).CreateClientEvent):len((*c.CallOptions).CreateClientEvent)], opts...)
	var resp *talentpb.ClientEvent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.eventClient.CreateClientEvent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
