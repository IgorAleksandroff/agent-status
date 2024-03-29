// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/IgorAleksandroff/agent-status/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// StatusRepository is an autogenerated mock type for the StatusRepository type
type StatusRepository struct {
	mock.Mock
}

type StatusRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *StatusRepository) EXPECT() *StatusRepository_Expecter {
	return &StatusRepository_Expecter{mock: &_m.Mock}
}

// AgentSetStatusTx provides a mock function with given fields: ctx, agent, mode
func (_m *StatusRepository) AgentSetStatusTx(ctx context.Context, agent entity.Agent, mode entity.Mode) (int64, error) {
	ret := _m.Called(ctx, agent, mode)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, entity.Agent, entity.Mode) int64); ok {
		r0 = rf(ctx, agent, mode)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.Agent, entity.Mode) error); ok {
		r1 = rf(ctx, agent, mode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusRepository_AgentSetStatusTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AgentSetStatusTx'
type StatusRepository_AgentSetStatusTx_Call struct {
	*mock.Call
}

// AgentSetStatusTx is a helper method to define mock.On call
//   - ctx context.Context
//   - agent entity.Agent
//   - mode entity.Mode
func (_e *StatusRepository_Expecter) AgentSetStatusTx(ctx interface{}, agent interface{}, mode interface{}) *StatusRepository_AgentSetStatusTx_Call {
	return &StatusRepository_AgentSetStatusTx_Call{Call: _e.mock.On("AgentSetStatusTx", ctx, agent, mode)}
}

func (_c *StatusRepository_AgentSetStatusTx_Call) Run(run func(ctx context.Context, agent entity.Agent, mode entity.Mode)) *StatusRepository_AgentSetStatusTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Agent), args[2].(entity.Mode))
	})
	return _c
}

func (_c *StatusRepository_AgentSetStatusTx_Call) Return(_a0 int64, _a1 error) *StatusRepository_AgentSetStatusTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLogsForAgent provides a mock function with given fields: ctx, login
func (_m *StatusRepository) GetLogsForAgent(ctx context.Context, login string) ([]entity.Logs, error) {
	ret := _m.Called(ctx, login)

	var r0 []entity.Logs
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Logs); ok {
		r0 = rf(ctx, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Logs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusRepository_GetLogsForAgent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLogsForAgent'
type StatusRepository_GetLogsForAgent_Call struct {
	*mock.Call
}

// GetLogsForAgent is a helper method to define mock.On call
//   - ctx context.Context
//   - login string
func (_e *StatusRepository_Expecter) GetLogsForAgent(ctx interface{}, login interface{}) *StatusRepository_GetLogsForAgent_Call {
	return &StatusRepository_GetLogsForAgent_Call{Call: _e.mock.On("GetLogsForAgent", ctx, login)}
}

func (_c *StatusRepository_GetLogsForAgent_Call) Run(run func(ctx context.Context, login string)) *StatusRepository_GetLogsForAgent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *StatusRepository_GetLogsForAgent_Call) Return(_a0 []entity.Logs, _a1 error) *StatusRepository_GetLogsForAgent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUser provides a mock function with given fields: ctx, login
func (_m *StatusRepository) GetUser(ctx context.Context, login string) (entity.Agent, error) {
	ret := _m.Called(ctx, login)

	var r0 entity.Agent
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Agent); ok {
		r0 = rf(ctx, login)
	} else {
		r0 = ret.Get(0).(entity.Agent)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusRepository_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type StatusRepository_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - ctx context.Context
//   - login string
func (_e *StatusRepository_Expecter) GetUser(ctx interface{}, login interface{}) *StatusRepository_GetUser_Call {
	return &StatusRepository_GetUser_Call{Call: _e.mock.On("GetUser", ctx, login)}
}

func (_c *StatusRepository_GetUser_Call) Run(run func(ctx context.Context, login string)) *StatusRepository_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *StatusRepository_GetUser_Call) Return(_a0 entity.Agent, _a1 error) *StatusRepository_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewStatusRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatusRepository creates a new instance of StatusRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatusRepository(t mockConstructorTestingTNewStatusRepository) *StatusRepository {
	mock := &StatusRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
