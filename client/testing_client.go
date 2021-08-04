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

package client

import (
	"context"
	"fmt"
	"math"
	"net/url"

	genprotopb "github.com/googleapis/gapic-showcase/server/genproto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newTestingClientHook clientHook

// TestingCallOptions contains the retry settings for each method of TestingClient.
type TestingCallOptions struct {
	CreateSession      []gax.CallOption
	GetSession         []gax.CallOption
	ListSessions       []gax.CallOption
	DeleteSession      []gax.CallOption
	ReportSession      []gax.CallOption
	ListTests          []gax.CallOption
	DeleteTest         []gax.CallOption
	VerifyTest         []gax.CallOption
	ListLocations      []gax.CallOption
	GetLocation        []gax.CallOption
	SetIamPolicy       []gax.CallOption
	GetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
	ListOperations     []gax.CallOption
	GetOperation       []gax.CallOption
	DeleteOperation    []gax.CallOption
	CancelOperation    []gax.CallOption
}

func defaultTestingGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("localhost:7469"),
		internaloption.WithDefaultMTLSEndpoint("localhost:7469"),
		internaloption.WithDefaultAudience("https://localhost/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTestingCallOptions() *TestingCallOptions {
	return &TestingCallOptions{
		CreateSession:      []gax.CallOption{},
		GetSession:         []gax.CallOption{},
		ListSessions:       []gax.CallOption{},
		DeleteSession:      []gax.CallOption{},
		ReportSession:      []gax.CallOption{},
		ListTests:          []gax.CallOption{},
		DeleteTest:         []gax.CallOption{},
		VerifyTest:         []gax.CallOption{},
		ListLocations:      []gax.CallOption{},
		GetLocation:        []gax.CallOption{},
		SetIamPolicy:       []gax.CallOption{},
		GetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
		ListOperations:     []gax.CallOption{},
		GetOperation:       []gax.CallOption{},
		DeleteOperation:    []gax.CallOption{},
		CancelOperation:    []gax.CallOption{},
	}
}

// internalTestingClient is an interface that defines the methods availaible from Client Libraries Showcase API.
type internalTestingClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateSession(context.Context, *genprotopb.CreateSessionRequest, ...gax.CallOption) (*genprotopb.Session, error)
	GetSession(context.Context, *genprotopb.GetSessionRequest, ...gax.CallOption) (*genprotopb.Session, error)
	ListSessions(context.Context, *genprotopb.ListSessionsRequest, ...gax.CallOption) *SessionIterator
	DeleteSession(context.Context, *genprotopb.DeleteSessionRequest, ...gax.CallOption) error
	ReportSession(context.Context, *genprotopb.ReportSessionRequest, ...gax.CallOption) (*genprotopb.ReportSessionResponse, error)
	ListTests(context.Context, *genprotopb.ListTestsRequest, ...gax.CallOption) *TestIterator
	DeleteTest(context.Context, *genprotopb.DeleteTestRequest, ...gax.CallOption) error
	VerifyTest(context.Context, *genprotopb.VerifyTestRequest, ...gax.CallOption) (*genprotopb.VerifyTestResponse, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
}

// TestingClient is a client for interacting with Client Libraries Showcase API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service to facilitate running discrete sets of tests
// against Showcase.
type TestingClient struct {
	// The internal transport-dependent client.
	internalClient internalTestingClient

	// The call options for this service.
	CallOptions *TestingCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TestingClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TestingClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TestingClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateSession creates a new testing session.
func (c *TestingClient) CreateSession(ctx context.Context, req *genprotopb.CreateSessionRequest, opts ...gax.CallOption) (*genprotopb.Session, error) {
	return c.internalClient.CreateSession(ctx, req, opts...)
}

// GetSession gets a testing session.
func (c *TestingClient) GetSession(ctx context.Context, req *genprotopb.GetSessionRequest, opts ...gax.CallOption) (*genprotopb.Session, error) {
	return c.internalClient.GetSession(ctx, req, opts...)
}

// ListSessions lists the current test sessions.
func (c *TestingClient) ListSessions(ctx context.Context, req *genprotopb.ListSessionsRequest, opts ...gax.CallOption) *SessionIterator {
	return c.internalClient.ListSessions(ctx, req, opts...)
}

// DeleteSession delete a test session.
func (c *TestingClient) DeleteSession(ctx context.Context, req *genprotopb.DeleteSessionRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteSession(ctx, req, opts...)
}

// ReportSession report on the status of a session.
// This generates a report detailing which tests have been completed,
// and an overall rollup.
func (c *TestingClient) ReportSession(ctx context.Context, req *genprotopb.ReportSessionRequest, opts ...gax.CallOption) (*genprotopb.ReportSessionResponse, error) {
	return c.internalClient.ReportSession(ctx, req, opts...)
}

// ListTests list the tests of a sessesion.
func (c *TestingClient) ListTests(ctx context.Context, req *genprotopb.ListTestsRequest, opts ...gax.CallOption) *TestIterator {
	return c.internalClient.ListTests(ctx, req, opts...)
}

// DeleteTest explicitly decline to implement a test.
//
// This removes the test from subsequent ListTests calls, and
// attempting to do the test will error.
//
// This method will error if attempting to delete a required test.
func (c *TestingClient) DeleteTest(ctx context.Context, req *genprotopb.DeleteTestRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteTest(ctx, req, opts...)
}

// VerifyTest register a response to a test.
//
// In cases where a test involves registering a final answer at the
// end of the test, this method provides the means to do so.
func (c *TestingClient) VerifyTest(ctx context.Context, req *genprotopb.VerifyTestRequest, opts ...gax.CallOption) (*genprotopb.VerifyTestResponse, error) {
	return c.internalClient.VerifyTest(ctx, req, opts...)
}

// ListLocations is a utility method from google.cloud.location.Locations.
func (c *TestingClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetLocation is a utility method from google.cloud.location.Locations.
func (c *TestingClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// SetIamPolicy is a utility method from google.iam.v1.IAMPolicy.
func (c *TestingClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// GetIamPolicy is a utility method from google.iam.v1.IAMPolicy.
func (c *TestingClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions is a utility method from google.iam.v1.IAMPolicy.
func (c *TestingClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *TestingClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *TestingClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *TestingClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *TestingClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// testingGRPCClient is a client for interacting with Client Libraries Showcase API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type testingGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing TestingClient
	CallOptions **TestingCallOptions

	// The gRPC API client.
	testingClient genprotopb.TestingClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTestingClient creates a new testing client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service to facilitate running discrete sets of tests
// against Showcase.
func NewTestingClient(ctx context.Context, opts ...option.ClientOption) (*TestingClient, error) {
	clientOpts := defaultTestingGRPCClientOptions()
	if newTestingClientHook != nil {
		hookOpts, err := newTestingClientHook(ctx, clientHookParams{})
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
	client := TestingClient{CallOptions: defaultTestingCallOptions()}

	c := &testingGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		testingClient:    genprotopb.NewTestingClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient:  iampb.NewIAMPolicyClient(connPool),
		locationsClient:  locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *testingGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *testingGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *testingGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *testingGRPCClient) CreateSession(ctx context.Context, req *genprotopb.CreateSessionRequest, opts ...gax.CallOption) (*genprotopb.Session, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateSession[0:len((*c.CallOptions).CreateSession):len((*c.CallOptions).CreateSession)], opts...)
	var resp *genprotopb.Session
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testingClient.CreateSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) GetSession(ctx context.Context, req *genprotopb.GetSessionRequest, opts ...gax.CallOption) (*genprotopb.Session, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSession[0:len((*c.CallOptions).GetSession):len((*c.CallOptions).GetSession)], opts...)
	var resp *genprotopb.Session
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testingClient.GetSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) ListSessions(ctx context.Context, req *genprotopb.ListSessionsRequest, opts ...gax.CallOption) *SessionIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListSessions[0:len((*c.CallOptions).ListSessions):len((*c.CallOptions).ListSessions)], opts...)
	it := &SessionIterator{}
	req = proto.Clone(req).(*genprotopb.ListSessionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*genprotopb.Session, string, error) {
		resp := &genprotopb.ListSessionsResponse{}
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
			resp, err = c.testingClient.ListSessions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSessions(), resp.GetNextPageToken(), nil
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

func (c *testingGRPCClient) DeleteSession(ctx context.Context, req *genprotopb.DeleteSessionRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteSession[0:len((*c.CallOptions).DeleteSession):len((*c.CallOptions).DeleteSession)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.testingClient.DeleteSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *testingGRPCClient) ReportSession(ctx context.Context, req *genprotopb.ReportSessionRequest, opts ...gax.CallOption) (*genprotopb.ReportSessionResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReportSession[0:len((*c.CallOptions).ReportSession):len((*c.CallOptions).ReportSession)], opts...)
	var resp *genprotopb.ReportSessionResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testingClient.ReportSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) ListTests(ctx context.Context, req *genprotopb.ListTestsRequest, opts ...gax.CallOption) *TestIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListTests[0:len((*c.CallOptions).ListTests):len((*c.CallOptions).ListTests)], opts...)
	it := &TestIterator{}
	req = proto.Clone(req).(*genprotopb.ListTestsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*genprotopb.Test, string, error) {
		resp := &genprotopb.ListTestsResponse{}
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
			resp, err = c.testingClient.ListTests(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTests(), resp.GetNextPageToken(), nil
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

func (c *testingGRPCClient) DeleteTest(ctx context.Context, req *genprotopb.DeleteTestRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTest[0:len((*c.CallOptions).DeleteTest):len((*c.CallOptions).DeleteTest)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.testingClient.DeleteTest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *testingGRPCClient) VerifyTest(ctx context.Context, req *genprotopb.VerifyTestRequest, opts ...gax.CallOption) (*genprotopb.VerifyTestResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).VerifyTest[0:len((*c.CallOptions).VerifyTest):len((*c.CallOptions).VerifyTest)], opts...)
	var resp *genprotopb.VerifyTestResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testingClient.VerifyTest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
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
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
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

func (c *testingGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

func (c *testingGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *testingGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *testingGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SessionIterator manages a stream of *genprotopb.Session.
type SessionIterator struct {
	items    []*genprotopb.Session
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
	InternalFetch func(pageSize int, pageToken string) (results []*genprotopb.Session, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SessionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SessionIterator) Next() (*genprotopb.Session, error) {
	var item *genprotopb.Session
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SessionIterator) bufLen() int {
	return len(it.items)
}

func (it *SessionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TestIterator manages a stream of *genprotopb.Test.
type TestIterator struct {
	items    []*genprotopb.Test
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
	InternalFetch func(pageSize int, pageToken string) (results []*genprotopb.Test, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TestIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TestIterator) Next() (*genprotopb.Test, error) {
	var item *genprotopb.Test
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TestIterator) bufLen() int {
	return len(it.items)
}

func (it *TestIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
