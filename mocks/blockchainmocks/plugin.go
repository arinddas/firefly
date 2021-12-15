// Code generated by mockery v1.0.0. DO NOT EDIT.

package blockchainmocks

import (
	config "github.com/hyperledger/firefly/internal/config"
	blockchain "github.com/hyperledger/firefly/pkg/blockchain"

	context "context"

	fftypes "github.com/hyperledger/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// AddSubscription provides a mock function with given fields: ctx, subscription
func (_m *Plugin) AddSubscription(ctx context.Context, subscription *fftypes.ContractSubscriptionInput) error {
	ret := _m.Called(ctx, subscription)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.ContractSubscriptionInput) error); ok {
		r0 = rf(ctx, subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Capabilities provides a mock function with given fields:
func (_m *Plugin) Capabilities() *blockchain.Capabilities {
	ret := _m.Called()

	var r0 *blockchain.Capabilities
	if rf, ok := ret.Get(0).(func() *blockchain.Capabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*blockchain.Capabilities)
		}
	}

	return r0
}

// DeleteSubscription provides a mock function with given fields: ctx, subscription
func (_m *Plugin) DeleteSubscription(ctx context.Context, subscription *fftypes.ContractSubscription) error {
	ret := _m.Called(ctx, subscription)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.ContractSubscription) error); ok {
		r0 = rf(ctx, subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Init provides a mock function with given fields: ctx, prefix, callbacks
func (_m *Plugin) Init(ctx context.Context, prefix config.Prefix, callbacks blockchain.Callbacks) error {
	ret := _m.Called(ctx, prefix, callbacks)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Prefix, blockchain.Callbacks) error); ok {
		r0 = rf(ctx, prefix, callbacks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitPrefix provides a mock function with given fields: prefix
func (_m *Plugin) InitPrefix(prefix config.Prefix) {
	_m.Called(prefix)
}

// InvokeContract provides a mock function with given fields: ctx, operationID, signingKey, location, method, params
func (_m *Plugin) InvokeContract(ctx context.Context, operationID *fftypes.UUID, signingKey string, location fftypes.Byteable, method *fftypes.FFIMethod, params map[string]interface{}) (interface{}, error) {
	ret := _m.Called(ctx, operationID, signingKey, location, method, params)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, string, fftypes.Byteable, *fftypes.FFIMethod, map[string]interface{}) interface{}); ok {
		r0 = rf(ctx, operationID, signingKey, location, method, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, string, fftypes.Byteable, *fftypes.FFIMethod, map[string]interface{}) error); ok {
		r1 = rf(ctx, operationID, signingKey, location, method, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *Plugin) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResolveSigningKey provides a mock function with given fields: ctx, signingKey
func (_m *Plugin) ResolveSigningKey(ctx context.Context, signingKey string) (string, error) {
	ret := _m.Called(ctx, signingKey)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, signingKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, signingKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Plugin) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubmitBatchPin provides a mock function with given fields: ctx, operationID, ledgerID, signingKey, batch
func (_m *Plugin) SubmitBatchPin(ctx context.Context, operationID *fftypes.UUID, ledgerID *fftypes.UUID, signingKey string, batch *blockchain.BatchPin) error {
	ret := _m.Called(ctx, operationID, ledgerID, signingKey, batch)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, *fftypes.UUID, string, *blockchain.BatchPin) error); ok {
		r0 = rf(ctx, operationID, ledgerID, signingKey, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateContractLocation provides a mock function with given fields: ctx, location
func (_m *Plugin) ValidateContractLocation(ctx context.Context, location fftypes.Byteable) error {
	ret := _m.Called(ctx, location)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, fftypes.Byteable) error); ok {
		r0 = rf(ctx, location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateFFIParam provides a mock function with given fields: ctx, method
func (_m *Plugin) ValidateFFIParam(ctx context.Context, method *fftypes.FFIParam) error {
	ret := _m.Called(ctx, method)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIParam) error); ok {
		r0 = rf(ctx, method)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
