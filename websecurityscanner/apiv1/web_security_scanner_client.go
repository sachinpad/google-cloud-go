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

package websecurityscanner

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	websecurityscannerpb "google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateScanConfig     []gax.CallOption
	DeleteScanConfig     []gax.CallOption
	GetScanConfig        []gax.CallOption
	ListScanConfigs      []gax.CallOption
	UpdateScanConfig     []gax.CallOption
	StartScanRun         []gax.CallOption
	GetScanRun           []gax.CallOption
	ListScanRuns         []gax.CallOption
	StopScanRun          []gax.CallOption
	ListCrawledUrls      []gax.CallOption
	GetFinding           []gax.CallOption
	ListFindings         []gax.CallOption
	ListFindingTypeStats []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("websecurityscanner.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("websecurityscanner.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://websecurityscanner.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateScanConfig: []gax.CallOption{},
		DeleteScanConfig: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetScanConfig: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListScanConfigs: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateScanConfig: []gax.CallOption{},
		StartScanRun:     []gax.CallOption{},
		GetScanRun: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListScanRuns: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		StopScanRun: []gax.CallOption{},
		ListCrawledUrls: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetFinding: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListFindings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListFindingTypeStats: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalClient is an interface that defines the methods availaible from Web Security Scanner API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateScanConfig(context.Context, *websecurityscannerpb.CreateScanConfigRequest, ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error)
	DeleteScanConfig(context.Context, *websecurityscannerpb.DeleteScanConfigRequest, ...gax.CallOption) error
	GetScanConfig(context.Context, *websecurityscannerpb.GetScanConfigRequest, ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error)
	ListScanConfigs(context.Context, *websecurityscannerpb.ListScanConfigsRequest, ...gax.CallOption) *ScanConfigIterator
	UpdateScanConfig(context.Context, *websecurityscannerpb.UpdateScanConfigRequest, ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error)
	StartScanRun(context.Context, *websecurityscannerpb.StartScanRunRequest, ...gax.CallOption) (*websecurityscannerpb.ScanRun, error)
	GetScanRun(context.Context, *websecurityscannerpb.GetScanRunRequest, ...gax.CallOption) (*websecurityscannerpb.ScanRun, error)
	ListScanRuns(context.Context, *websecurityscannerpb.ListScanRunsRequest, ...gax.CallOption) *ScanRunIterator
	StopScanRun(context.Context, *websecurityscannerpb.StopScanRunRequest, ...gax.CallOption) (*websecurityscannerpb.ScanRun, error)
	ListCrawledUrls(context.Context, *websecurityscannerpb.ListCrawledUrlsRequest, ...gax.CallOption) *CrawledUrlIterator
	GetFinding(context.Context, *websecurityscannerpb.GetFindingRequest, ...gax.CallOption) (*websecurityscannerpb.Finding, error)
	ListFindings(context.Context, *websecurityscannerpb.ListFindingsRequest, ...gax.CallOption) *FindingIterator
	ListFindingTypeStats(context.Context, *websecurityscannerpb.ListFindingTypeStatsRequest, ...gax.CallOption) (*websecurityscannerpb.ListFindingTypeStatsResponse, error)
}

// Client is a client for interacting with Web Security Scanner API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Web Security Scanner Service identifies security vulnerabilities in web
// applications hosted on Google Cloud. It crawls your application, and
// attempts to exercise as many user inputs and event handlers as possible.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateScanConfig creates a new ScanConfig.
func (c *Client) CreateScanConfig(ctx context.Context, req *websecurityscannerpb.CreateScanConfigRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error) {
	return c.internalClient.CreateScanConfig(ctx, req, opts...)
}

// DeleteScanConfig deletes an existing ScanConfig and its child resources.
func (c *Client) DeleteScanConfig(ctx context.Context, req *websecurityscannerpb.DeleteScanConfigRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteScanConfig(ctx, req, opts...)
}

// GetScanConfig gets a ScanConfig.
func (c *Client) GetScanConfig(ctx context.Context, req *websecurityscannerpb.GetScanConfigRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error) {
	return c.internalClient.GetScanConfig(ctx, req, opts...)
}

// ListScanConfigs lists ScanConfigs under a given project.
func (c *Client) ListScanConfigs(ctx context.Context, req *websecurityscannerpb.ListScanConfigsRequest, opts ...gax.CallOption) *ScanConfigIterator {
	return c.internalClient.ListScanConfigs(ctx, req, opts...)
}

// UpdateScanConfig updates a ScanConfig. This method support partial update of a ScanConfig.
func (c *Client) UpdateScanConfig(ctx context.Context, req *websecurityscannerpb.UpdateScanConfigRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error) {
	return c.internalClient.UpdateScanConfig(ctx, req, opts...)
}

// StartScanRun start a ScanRun according to the given ScanConfig.
func (c *Client) StartScanRun(ctx context.Context, req *websecurityscannerpb.StartScanRunRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanRun, error) {
	return c.internalClient.StartScanRun(ctx, req, opts...)
}

// GetScanRun gets a ScanRun.
func (c *Client) GetScanRun(ctx context.Context, req *websecurityscannerpb.GetScanRunRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanRun, error) {
	return c.internalClient.GetScanRun(ctx, req, opts...)
}

// ListScanRuns lists ScanRuns under a given ScanConfig, in descending order of ScanRun
// stop time.
func (c *Client) ListScanRuns(ctx context.Context, req *websecurityscannerpb.ListScanRunsRequest, opts ...gax.CallOption) *ScanRunIterator {
	return c.internalClient.ListScanRuns(ctx, req, opts...)
}

// StopScanRun stops a ScanRun. The stopped ScanRun is returned.
func (c *Client) StopScanRun(ctx context.Context, req *websecurityscannerpb.StopScanRunRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanRun, error) {
	return c.internalClient.StopScanRun(ctx, req, opts...)
}

// ListCrawledUrls list CrawledUrls under a given ScanRun.
func (c *Client) ListCrawledUrls(ctx context.Context, req *websecurityscannerpb.ListCrawledUrlsRequest, opts ...gax.CallOption) *CrawledUrlIterator {
	return c.internalClient.ListCrawledUrls(ctx, req, opts...)
}

// GetFinding gets a Finding.
func (c *Client) GetFinding(ctx context.Context, req *websecurityscannerpb.GetFindingRequest, opts ...gax.CallOption) (*websecurityscannerpb.Finding, error) {
	return c.internalClient.GetFinding(ctx, req, opts...)
}

// ListFindings list Findings under a given ScanRun.
func (c *Client) ListFindings(ctx context.Context, req *websecurityscannerpb.ListFindingsRequest, opts ...gax.CallOption) *FindingIterator {
	return c.internalClient.ListFindings(ctx, req, opts...)
}

// ListFindingTypeStats list all FindingTypeStats under a given ScanRun.
func (c *Client) ListFindingTypeStats(ctx context.Context, req *websecurityscannerpb.ListFindingTypeStatsRequest, opts ...gax.CallOption) (*websecurityscannerpb.ListFindingTypeStatsResponse, error) {
	return c.internalClient.ListFindingTypeStats(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Web Security Scanner API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client websecurityscannerpb.WebSecurityScannerClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new web security scanner client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Web Security Scanner Service identifies security vulnerabilities in web
// applications hosted on Google Cloud. It crawls your application, and
// attempts to exercise as many user inputs and event handlers as possible.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
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
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           websecurityscannerpb.NewWebSecurityScannerClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) CreateScanConfig(ctx context.Context, req *websecurityscannerpb.CreateScanConfigRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateScanConfig[0:len((*c.CallOptions).CreateScanConfig):len((*c.CallOptions).CreateScanConfig)], opts...)
	var resp *websecurityscannerpb.ScanConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateScanConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteScanConfig(ctx context.Context, req *websecurityscannerpb.DeleteScanConfigRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteScanConfig[0:len((*c.CallOptions).DeleteScanConfig):len((*c.CallOptions).DeleteScanConfig)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteScanConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetScanConfig(ctx context.Context, req *websecurityscannerpb.GetScanConfigRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetScanConfig[0:len((*c.CallOptions).GetScanConfig):len((*c.CallOptions).GetScanConfig)], opts...)
	var resp *websecurityscannerpb.ScanConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetScanConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListScanConfigs(ctx context.Context, req *websecurityscannerpb.ListScanConfigsRequest, opts ...gax.CallOption) *ScanConfigIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListScanConfigs[0:len((*c.CallOptions).ListScanConfigs):len((*c.CallOptions).ListScanConfigs)], opts...)
	it := &ScanConfigIterator{}
	req = proto.Clone(req).(*websecurityscannerpb.ListScanConfigsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*websecurityscannerpb.ScanConfig, string, error) {
		resp := &websecurityscannerpb.ListScanConfigsResponse{}
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
			resp, err = c.client.ListScanConfigs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetScanConfigs(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) UpdateScanConfig(ctx context.Context, req *websecurityscannerpb.UpdateScanConfigRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "scan_config.name", url.QueryEscape(req.GetScanConfig().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateScanConfig[0:len((*c.CallOptions).UpdateScanConfig):len((*c.CallOptions).UpdateScanConfig)], opts...)
	var resp *websecurityscannerpb.ScanConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateScanConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) StartScanRun(ctx context.Context, req *websecurityscannerpb.StartScanRunRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanRun, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).StartScanRun[0:len((*c.CallOptions).StartScanRun):len((*c.CallOptions).StartScanRun)], opts...)
	var resp *websecurityscannerpb.ScanRun
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.StartScanRun(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetScanRun(ctx context.Context, req *websecurityscannerpb.GetScanRunRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanRun, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetScanRun[0:len((*c.CallOptions).GetScanRun):len((*c.CallOptions).GetScanRun)], opts...)
	var resp *websecurityscannerpb.ScanRun
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetScanRun(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListScanRuns(ctx context.Context, req *websecurityscannerpb.ListScanRunsRequest, opts ...gax.CallOption) *ScanRunIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListScanRuns[0:len((*c.CallOptions).ListScanRuns):len((*c.CallOptions).ListScanRuns)], opts...)
	it := &ScanRunIterator{}
	req = proto.Clone(req).(*websecurityscannerpb.ListScanRunsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*websecurityscannerpb.ScanRun, string, error) {
		resp := &websecurityscannerpb.ListScanRunsResponse{}
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
			resp, err = c.client.ListScanRuns(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetScanRuns(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) StopScanRun(ctx context.Context, req *websecurityscannerpb.StopScanRunRequest, opts ...gax.CallOption) (*websecurityscannerpb.ScanRun, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).StopScanRun[0:len((*c.CallOptions).StopScanRun):len((*c.CallOptions).StopScanRun)], opts...)
	var resp *websecurityscannerpb.ScanRun
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.StopScanRun(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListCrawledUrls(ctx context.Context, req *websecurityscannerpb.ListCrawledUrlsRequest, opts ...gax.CallOption) *CrawledUrlIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListCrawledUrls[0:len((*c.CallOptions).ListCrawledUrls):len((*c.CallOptions).ListCrawledUrls)], opts...)
	it := &CrawledUrlIterator{}
	req = proto.Clone(req).(*websecurityscannerpb.ListCrawledUrlsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*websecurityscannerpb.CrawledUrl, string, error) {
		resp := &websecurityscannerpb.ListCrawledUrlsResponse{}
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
			resp, err = c.client.ListCrawledUrls(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCrawledUrls(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetFinding(ctx context.Context, req *websecurityscannerpb.GetFindingRequest, opts ...gax.CallOption) (*websecurityscannerpb.Finding, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFinding[0:len((*c.CallOptions).GetFinding):len((*c.CallOptions).GetFinding)], opts...)
	var resp *websecurityscannerpb.Finding
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetFinding(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListFindings(ctx context.Context, req *websecurityscannerpb.ListFindingsRequest, opts ...gax.CallOption) *FindingIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListFindings[0:len((*c.CallOptions).ListFindings):len((*c.CallOptions).ListFindings)], opts...)
	it := &FindingIterator{}
	req = proto.Clone(req).(*websecurityscannerpb.ListFindingsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*websecurityscannerpb.Finding, string, error) {
		resp := &websecurityscannerpb.ListFindingsResponse{}
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
			resp, err = c.client.ListFindings(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetFindings(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) ListFindingTypeStats(ctx context.Context, req *websecurityscannerpb.ListFindingTypeStatsRequest, opts ...gax.CallOption) (*websecurityscannerpb.ListFindingTypeStatsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListFindingTypeStats[0:len((*c.CallOptions).ListFindingTypeStats):len((*c.CallOptions).ListFindingTypeStats)], opts...)
	var resp *websecurityscannerpb.ListFindingTypeStatsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ListFindingTypeStats(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CrawledUrlIterator manages a stream of *websecurityscannerpb.CrawledUrl.
type CrawledUrlIterator struct {
	items    []*websecurityscannerpb.CrawledUrl
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
	InternalFetch func(pageSize int, pageToken string) (results []*websecurityscannerpb.CrawledUrl, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CrawledUrlIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CrawledUrlIterator) Next() (*websecurityscannerpb.CrawledUrl, error) {
	var item *websecurityscannerpb.CrawledUrl
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CrawledUrlIterator) bufLen() int {
	return len(it.items)
}

func (it *CrawledUrlIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// FindingIterator manages a stream of *websecurityscannerpb.Finding.
type FindingIterator struct {
	items    []*websecurityscannerpb.Finding
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
	InternalFetch func(pageSize int, pageToken string) (results []*websecurityscannerpb.Finding, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *FindingIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *FindingIterator) Next() (*websecurityscannerpb.Finding, error) {
	var item *websecurityscannerpb.Finding
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *FindingIterator) bufLen() int {
	return len(it.items)
}

func (it *FindingIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ScanConfigIterator manages a stream of *websecurityscannerpb.ScanConfig.
type ScanConfigIterator struct {
	items    []*websecurityscannerpb.ScanConfig
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
	InternalFetch func(pageSize int, pageToken string) (results []*websecurityscannerpb.ScanConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ScanConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ScanConfigIterator) Next() (*websecurityscannerpb.ScanConfig, error) {
	var item *websecurityscannerpb.ScanConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ScanConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *ScanConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ScanRunIterator manages a stream of *websecurityscannerpb.ScanRun.
type ScanRunIterator struct {
	items    []*websecurityscannerpb.ScanRun
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
	InternalFetch func(pageSize int, pageToken string) (results []*websecurityscannerpb.ScanRun, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ScanRunIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ScanRunIterator) Next() (*websecurityscannerpb.ScanRun, error) {
	var item *websecurityscannerpb.ScanRun
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ScanRunIterator) bufLen() int {
	return len(it.items)
}

func (it *ScanRunIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
