// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/IgorAleksandroff/agent-status/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Authorization is an autogenerated mock type for the Authorization type
type Authorization struct {
	mock.Mock
}

type Authorization_Expecter struct {
	mock *mock.Mock
}

func (_m *Authorization) EXPECT() *Authorization_Expecter {
	return &Authorization_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Authorization) CreateUser(ctx context.Context, user entity.Agent) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Agent) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Authorization_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type Authorization_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user entity.Agent
func (_e *Authorization_Expecter) CreateUser(ctx interface{}, user interface{}) *Authorization_CreateUser_Call {
	return &Authorization_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *Authorization_CreateUser_Call) Run(run func(ctx context.Context, user entity.Agent)) *Authorization_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Agent))
	})
	return _c
}

func (_c *Authorization_CreateUser_Call) Return(_a0 error) *Authorization_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

// GenerateToken provides a mock function with given fields: ctx, username, password
func (_m *Authorization) GenerateToken(ctx context.Context, username string, password string) (string, error) {
	ret := _m.Called(ctx, username, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Authorization_GenerateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateToken'
type Authorization_GenerateToken_Call struct {
	*mock.Call
}

// GenerateToken is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - password string
func (_e *Authorization_Expecter) GenerateToken(ctx interface{}, username interface{}, password interface{}) *Authorization_GenerateToken_Call {
	return &Authorization_GenerateToken_Call{Call: _e.mock.On("GenerateToken", ctx, username, password)}
}

func (_c *Authorization_GenerateToken_Call) Run(run func(ctx context.Context, username string, password string)) *Authorization_GenerateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Authorization_GenerateToken_Call) Return(_a0 string, _a1 error) *Authorization_GenerateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ParseToken provides a mock function with given fields: token
func (_m *Authorization) ParseToken(token string) (string, error) {
	ret := _m.Called(token)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Authorization_ParseToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseToken'
type Authorization_ParseToken_Call struct {
	*mock.Call
}

// ParseToken is a helper method to define mock.On call
//   - token string
func (_e *Authorization_Expecter) ParseToken(token interface{}) *Authorization_ParseToken_Call {
	return &Authorization_ParseToken_Call{Call: _e.mock.On("ParseToken", token)}
}

func (_c *Authorization_ParseToken_Call) Run(run func(token string)) *Authorization_ParseToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Authorization_ParseToken_Call) Return(_a0 string, _a1 error) *Authorization_ParseToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewAuthorization interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthorization creates a new instance of Authorization. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthorization(t mockConstructorTestingTNewAuthorization) *Authorization {
	mock := &Authorization{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}