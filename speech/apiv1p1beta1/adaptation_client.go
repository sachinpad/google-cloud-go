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

package speech

import (
	"context"
	"fmt"
	"math"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newAdaptationClientHook clientHook

// AdaptationCallOptions contains the retry settings for each method of AdaptationClient.
type AdaptationCallOptions struct {
	CreatePhraseSet   []gax.CallOption
	GetPhraseSet      []gax.CallOption
	ListPhraseSet     []gax.CallOption
	UpdatePhraseSet   []gax.CallOption
	DeletePhraseSet   []gax.CallOption
	CreateCustomClass []gax.CallOption
	GetCustomClass    []gax.CallOption
	ListCustomClasses []gax.CallOption
	UpdateCustomClass []gax.CallOption
	DeleteCustomClass []gax.CallOption
}

func defaultAdaptationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("speech.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("speech.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://speech.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAdaptationCallOptions() *AdaptationCallOptions {
	return &AdaptationCallOptions{
		CreatePhraseSet:   []gax.CallOption{},
		GetPhraseSet:      []gax.CallOption{},
		ListPhraseSet:     []gax.CallOption{},
		UpdatePhraseSet:   []gax.CallOption{},
		DeletePhraseSet:   []gax.CallOption{},
		CreateCustomClass: []gax.CallOption{},
		GetCustomClass:    []gax.CallOption{},
		ListCustomClasses: []gax.CallOption{},
		UpdateCustomClass: []gax.CallOption{},
		DeleteCustomClass: []gax.CallOption{},
	}
}

// internalAdaptationClient is an interface that defines the methods availaible from Cloud Speech-to-Text API.
type internalAdaptationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreatePhraseSet(context.Context, *speechpb.CreatePhraseSetRequest, ...gax.CallOption) (*speechpb.PhraseSet, error)
	GetPhraseSet(context.Context, *speechpb.GetPhraseSetRequest, ...gax.CallOption) (*speechpb.PhraseSet, error)
	ListPhraseSet(context.Context, *speechpb.ListPhraseSetRequest, ...gax.CallOption) *PhraseSetIterator
	UpdatePhraseSet(context.Context, *speechpb.UpdatePhraseSetRequest, ...gax.CallOption) (*speechpb.PhraseSet, error)
	DeletePhraseSet(context.Context, *speechpb.DeletePhraseSetRequest, ...gax.CallOption) error
	CreateCustomClass(context.Context, *speechpb.CreateCustomClassRequest, ...gax.CallOption) (*speechpb.CustomClass, error)
	GetCustomClass(context.Context, *speechpb.GetCustomClassRequest, ...gax.CallOption) (*speechpb.CustomClass, error)
	ListCustomClasses(context.Context, *speechpb.ListCustomClassesRequest, ...gax.CallOption) *CustomClassIterator
	UpdateCustomClass(context.Context, *speechpb.UpdateCustomClassRequest, ...gax.CallOption) (*speechpb.CustomClass, error)
	DeleteCustomClass(context.Context, *speechpb.DeleteCustomClassRequest, ...gax.CallOption) error
}

// AdaptationClient is a client for interacting with Cloud Speech-to-Text API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service that implements Google Cloud Speech Adaptation API.
type AdaptationClient struct {
	// The internal transport-dependent client.
	internalClient internalAdaptationClient

	// The call options for this service.
	CallOptions *AdaptationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdaptationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdaptationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AdaptationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreatePhraseSet create a set of phrase hints. Each item in the set can be a single word or
// a multi-word phrase. The items in the PhraseSet are favored by the
// recognition model when you send a call that includes the PhraseSet.
func (c *AdaptationClient) CreatePhraseSet(ctx context.Context, req *speechpb.CreatePhraseSetRequest, opts ...gax.CallOption) (*speechpb.PhraseSet, error) {
	return c.internalClient.CreatePhraseSet(ctx, req, opts...)
}

// GetPhraseSet get a phrase set.
func (c *AdaptationClient) GetPhraseSet(ctx context.Context, req *speechpb.GetPhraseSetRequest, opts ...gax.CallOption) (*speechpb.PhraseSet, error) {
	return c.internalClient.GetPhraseSet(ctx, req, opts...)
}

// ListPhraseSet list phrase sets.
func (c *AdaptationClient) ListPhraseSet(ctx context.Context, req *speechpb.ListPhraseSetRequest, opts ...gax.CallOption) *PhraseSetIterator {
	return c.internalClient.ListPhraseSet(ctx, req, opts...)
}

// UpdatePhraseSet update a phrase set.
func (c *AdaptationClient) UpdatePhraseSet(ctx context.Context, req *speechpb.UpdatePhraseSetRequest, opts ...gax.CallOption) (*speechpb.PhraseSet, error) {
	return c.internalClient.UpdatePhraseSet(ctx, req, opts...)
}

// DeletePhraseSet delete a phrase set.
func (c *AdaptationClient) DeletePhraseSet(ctx context.Context, req *speechpb.DeletePhraseSetRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeletePhraseSet(ctx, req, opts...)
}

// CreateCustomClass create a custom class.
func (c *AdaptationClient) CreateCustomClass(ctx context.Context, req *speechpb.CreateCustomClassRequest, opts ...gax.CallOption) (*speechpb.CustomClass, error) {
	return c.internalClient.CreateCustomClass(ctx, req, opts...)
}

// GetCustomClass get a custom class.
func (c *AdaptationClient) GetCustomClass(ctx context.Context, req *speechpb.GetCustomClassRequest, opts ...gax.CallOption) (*speechpb.CustomClass, error) {
	return c.internalClient.GetCustomClass(ctx, req, opts...)
}

// ListCustomClasses list custom classes.
func (c *AdaptationClient) ListCustomClasses(ctx context.Context, req *speechpb.ListCustomClassesRequest, opts ...gax.CallOption) *CustomClassIterator {
	return c.internalClient.ListCustomClasses(ctx, req, opts...)
}

// UpdateCustomClass update a custom class.
func (c *AdaptationClient) UpdateCustomClass(ctx context.Context, req *speechpb.UpdateCustomClassRequest, opts ...gax.CallOption) (*speechpb.CustomClass, error) {
	return c.internalClient.UpdateCustomClass(ctx, req, opts...)
}

// DeleteCustomClass delete a custom class.
func (c *AdaptationClient) DeleteCustomClass(ctx context.Context, req *speechpb.DeleteCustomClassRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteCustomClass(ctx, req, opts...)
}

// adaptationGRPCClient is a client for interacting with Cloud Speech-to-Text API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adaptationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AdaptationClient
	CallOptions **AdaptationCallOptions

	// The gRPC API client.
	adaptationClient speechpb.AdaptationClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdaptationClient creates a new adaptation client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service that implements Google Cloud Speech Adaptation API.
func NewAdaptationClient(ctx context.Context, opts ...option.ClientOption) (*AdaptationClient, error) {
	clientOpts := defaultAdaptationGRPCClientOptions()
	if newAdaptationClientHook != nil {
		hookOpts, err := newAdaptationClientHook(ctx, clientHookParams{})
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
	client := AdaptationClient{CallOptions: defaultAdaptationCallOptions()}

	c := &adaptationGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		adaptationClient: speechpb.NewAdaptationClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *adaptationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adaptationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adaptationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adaptationGRPCClient) CreatePhraseSet(ctx context.Context, req *speechpb.CreatePhraseSetRequest, opts ...gax.CallOption) (*speechpb.PhraseSet, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreatePhraseSet[0:len((*c.CallOptions).CreatePhraseSet):len((*c.CallOptions).CreatePhraseSet)], opts...)
	var resp *speechpb.PhraseSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adaptationClient.CreatePhraseSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adaptationGRPCClient) GetPhraseSet(ctx context.Context, req *speechpb.GetPhraseSetRequest, opts ...gax.CallOption) (*speechpb.PhraseSet, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetPhraseSet[0:len((*c.CallOptions).GetPhraseSet):len((*c.CallOptions).GetPhraseSet)], opts...)
	var resp *speechpb.PhraseSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adaptationClient.GetPhraseSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adaptationGRPCClient) ListPhraseSet(ctx context.Context, req *speechpb.ListPhraseSetRequest, opts ...gax.CallOption) *PhraseSetIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListPhraseSet[0:len((*c.CallOptions).ListPhraseSet):len((*c.CallOptions).ListPhraseSet)], opts...)
	it := &PhraseSetIterator{}
	req = proto.Clone(req).(*speechpb.ListPhraseSetRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*speechpb.PhraseSet, string, error) {
		resp := &speechpb.ListPhraseSetResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.adaptationClient.ListPhraseSet(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPhraseSets(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *adaptationGRPCClient) UpdatePhraseSet(ctx context.Context, req *speechpb.UpdatePhraseSetRequest, opts ...gax.CallOption) (*speechpb.PhraseSet, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "phrase_set.name", url.QueryEscape(req.GetPhraseSet().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdatePhraseSet[0:len((*c.CallOptions).UpdatePhraseSet):len((*c.CallOptions).UpdatePhraseSet)], opts...)
	var resp *speechpb.PhraseSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adaptationClient.UpdatePhraseSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adaptationGRPCClient) DeletePhraseSet(ctx context.Context, req *speechpb.DeletePhraseSetRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeletePhraseSet[0:len((*c.CallOptions).DeletePhraseSet):len((*c.CallOptions).DeletePhraseSet)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.adaptationClient.DeletePhraseSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *adaptationGRPCClient) CreateCustomClass(ctx context.Context, req *speechpb.CreateCustomClassRequest, opts ...gax.CallOption) (*speechpb.CustomClass, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateCustomClass[0:len((*c.CallOptions).CreateCustomClass):len((*c.CallOptions).CreateCustomClass)], opts...)
	var resp *speechpb.CustomClass
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adaptationClient.CreateCustomClass(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adaptationGRPCClient) GetCustomClass(ctx context.Context, req *speechpb.GetCustomClassRequest, opts ...gax.CallOption) (*speechpb.CustomClass, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCustomClass[0:len((*c.CallOptions).GetCustomClass):len((*c.CallOptions).GetCustomClass)], opts...)
	var resp *speechpb.CustomClass
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adaptationClient.GetCustomClass(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adaptationGRPCClient) ListCustomClasses(ctx context.Context, req *speechpb.ListCustomClassesRequest, opts ...gax.CallOption) *CustomClassIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListCustomClasses[0:len((*c.CallOptions).ListCustomClasses):len((*c.CallOptions).ListCustomClasses)], opts...)
	it := &CustomClassIterator{}
	req = proto.Clone(req).(*speechpb.ListCustomClassesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*speechpb.CustomClass, string, error) {
		resp := &speechpb.ListCustomClassesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.adaptationClient.ListCustomClasses(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCustomClasses(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *adaptationGRPCClient) UpdateCustomClass(ctx context.Context, req *speechpb.UpdateCustomClassRequest, opts ...gax.CallOption) (*speechpb.CustomClass, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "custom_class.name", url.QueryEscape(req.GetCustomClass().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateCustomClass[0:len((*c.CallOptions).UpdateCustomClass):len((*c.CallOptions).UpdateCustomClass)], opts...)
	var resp *speechpb.CustomClass
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adaptationClient.UpdateCustomClass(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adaptationGRPCClient) DeleteCustomClass(ctx context.Context, req *speechpb.DeleteCustomClassRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteCustomClass[0:len((*c.CallOptions).DeleteCustomClass):len((*c.CallOptions).DeleteCustomClass)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.adaptationClient.DeleteCustomClass(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CustomClassIterator manages a stream of *speechpb.CustomClass.
type CustomClassIterator struct {
	items    []*speechpb.CustomClass
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*speechpb.CustomClass, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CustomClassIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CustomClassIterator) Next() (*speechpb.CustomClass, error) {
	var item *speechpb.CustomClass
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CustomClassIterator) bufLen() int {
	return len(it.items)
}

func (it *CustomClassIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// PhraseSetIterator manages a stream of *speechpb.PhraseSet.
type PhraseSetIterator struct {
	items    []*speechpb.PhraseSet
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*speechpb.PhraseSet, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PhraseSetIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PhraseSetIterator) Next() (*speechpb.PhraseSet, error) {
	var item *speechpb.PhraseSet
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PhraseSetIterator) bufLen() int {
	return len(it.items)
}

func (it *PhraseSetIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
