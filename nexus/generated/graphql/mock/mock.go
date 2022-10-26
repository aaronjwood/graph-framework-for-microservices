// Code generated by MockGen. DO NOT EDIT.
// Source: generated/graphql/query_grpc.pb.go

// Package mock_graphql is a generated GoMock package.
package mock_graphql

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	graphql "github.com/vmware-tanzu/graph-framework-for-microservices/nexus/generated/graphql"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockServerClient is a mock of ServerClient interface
type MockServerClient struct {
	ctrl     *gomock.Controller
	recorder *MockServerClientMockRecorder
}

// MockServerClientMockRecorder is the mock recorder for MockServerClient
type MockServerClientMockRecorder struct {
	mock *MockServerClient
}

// NewMockServerClient creates a new mock instance
func NewMockServerClient(ctrl *gomock.Controller) *MockServerClient {
	mock := &MockServerClient{ctrl: ctrl}
	mock.recorder = &MockServerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerClient) EXPECT() *MockServerClientMockRecorder {
	return m.recorder
}

// Query mocks base method
func (m *MockServerClient) Query(ctx context.Context, in *graphql.GraphQLQuery, opts ...grpc.CallOption) (*graphql.GraphQLResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*graphql.GraphQLResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockServerClientMockRecorder) Query(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockServerClient)(nil).Query), varargs...)
}

// MockUnstableServerService is a mock of UnstableServerService interface
type MockUnstableServerService struct {
	ctrl     *gomock.Controller
	recorder *MockUnstableServerServiceMockRecorder
}

// MockUnstableServerServiceMockRecorder is the mock recorder for MockUnstableServerService
type MockUnstableServerServiceMockRecorder struct {
	mock *MockUnstableServerService
}

// NewMockUnstableServerService creates a new mock instance
func NewMockUnstableServerService(ctrl *gomock.Controller) *MockUnstableServerService {
	mock := &MockUnstableServerService{ctrl: ctrl}
	mock.recorder = &MockUnstableServerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnstableServerService) EXPECT() *MockUnstableServerServiceMockRecorder {
	return m.recorder
}

// Query mocks base method
func (m *MockUnstableServerService) Query(arg0 context.Context, arg1 *graphql.GraphQLQuery) (*graphql.GraphQLResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].(*graphql.GraphQLResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockUnstableServerServiceMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockUnstableServerService)(nil).Query), arg0, arg1)
}