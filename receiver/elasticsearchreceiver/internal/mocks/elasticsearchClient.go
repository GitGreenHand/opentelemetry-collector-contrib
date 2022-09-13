// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver/internal/model"
)

// MockElasticsearchClient is an autogenerated mock type for the elasticsearchClient type
type MockElasticsearchClient struct {
	mock.Mock
}

// ClusterHealth provides a mock function with given fields: ctx
func (_m *MockElasticsearchClient) ClusterHealth(ctx context.Context) (*model.ClusterHealth, error) {
	ret := _m.Called(ctx)

	var r0 *model.ClusterHealth
	if rf, ok := ret.Get(0).(func(context.Context) *model.ClusterHealth); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ClusterHealth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeStats provides a mock function with given fields: ctx, nodes
func (_m *MockElasticsearchClient) NodeStats(ctx context.Context, nodes []string) (*model.NodeStats, error) {
	ret := _m.Called(ctx, nodes)

	var r0 *model.NodeStats
	if rf, ok := ret.Get(0).(func(context.Context, []string) *model.NodeStats); ok {
		r0 = rf(ctx, nodes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NodeStats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, nodes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields: ctx
func (_m *MockElasticsearchClient) Version(ctx context.Context) (*model.VersionResponse, error) {
	ret := _m.Called(ctx)

	var r0 *model.VersionResponse
	if rf, ok := ret.Get(0).(func(context.Context) *model.VersionResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VersionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
