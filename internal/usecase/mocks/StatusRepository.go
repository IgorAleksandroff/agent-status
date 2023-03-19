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
