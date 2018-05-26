// Code generated by mockery v1.0.0
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import proto "github.com/gogo/protobuf/proto"

import types "github.com/ovrclk/akash/types"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Account provides a mock function with given fields: ctx, id
func (_m *Client) Account(ctx context.Context, id []byte) (*types.Account, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.Account
	if rf, ok := ret.Get(0).(func(context.Context, []byte) *types.Account); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deployment provides a mock function with given fields: ctx, id
func (_m *Client) Deployment(ctx context.Context, id []byte) (*types.Deployment, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, []byte) *types.Deployment); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeploymentGroup provides a mock function with given fields: ctx, id
func (_m *Client) DeploymentGroup(ctx context.Context, id types.DeploymentGroupID) (*types.DeploymentGroup, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.DeploymentGroup
	if rf, ok := ret.Get(0).(func(context.Context, types.DeploymentGroupID) *types.DeploymentGroup); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.DeploymentGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.DeploymentGroupID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeploymentLeases provides a mock function with given fields: ctx, id
func (_m *Client) DeploymentLeases(ctx context.Context, id []byte) (*types.Leases, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.Leases
	if rf, ok := ret.Get(0).(func(context.Context, []byte) *types.Leases); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Leases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deployments provides a mock function with given fields: ctx
func (_m *Client) Deployments(ctx context.Context) (*types.Deployments, error) {
	ret := _m.Called(ctx)

	var r0 *types.Deployments
	if rf, ok := ret.Get(0).(func(context.Context) *types.Deployments); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Deployments)
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

// Get provides a mock function with given fields: ctx, path, obj
func (_m *Client) Get(ctx context.Context, path string, obj proto.Message) error {
	ret := _m.Called(ctx, path, obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, proto.Message) error); ok {
		r0 = rf(ctx, path, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Lease provides a mock function with given fields: ctx, id
func (_m *Client) Lease(ctx context.Context, id types.LeaseID) (*types.Lease, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.Lease
	if rf, ok := ret.Get(0).(func(context.Context, types.LeaseID) *types.Lease); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.LeaseID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Leases provides a mock function with given fields: ctx
func (_m *Client) Leases(ctx context.Context) (*types.Leases, error) {
	ret := _m.Called(ctx)

	var r0 *types.Leases
	if rf, ok := ret.Get(0).(func(context.Context) *types.Leases); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Leases)
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

// Order provides a mock function with given fields: ctx, id
func (_m *Client) Order(ctx context.Context, id types.OrderID) (*types.Order, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.Order
	if rf, ok := ret.Get(0).(func(context.Context, types.OrderID) *types.Order); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.OrderID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Orders provides a mock function with given fields: ctx
func (_m *Client) Orders(ctx context.Context) (*types.Orders, error) {
	ret := _m.Called(ctx)

	var r0 *types.Orders
	if rf, ok := ret.Get(0).(func(context.Context) *types.Orders); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Orders)
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

// Provider provides a mock function with given fields: ctx, id
func (_m *Client) Provider(ctx context.Context, id []byte) (*types.Provider, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.Provider
	if rf, ok := ret.Get(0).(func(context.Context, []byte) *types.Provider); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Providers provides a mock function with given fields: ctx
func (_m *Client) Providers(ctx context.Context) (*types.Providers, error) {
	ret := _m.Called(ctx)

	var r0 *types.Providers
	if rf, ok := ret.Get(0).(func(context.Context) *types.Providers); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Providers)
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
