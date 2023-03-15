// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/IgorAleksandroff/agent-status/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Status is an autogenerated mock type for the Status type
type Status struct {
	mock.Mock
}

type Status_Expecter struct {
	mock *mock.Mock
}

func (_m *Status) EXPECT() *Status_Expecter {
	return &Status_Expecter{mock: &_m.Mock}
}

// AgentSetStatus provides a mock function with given fields: ctx, agent, mode
func (_m *Status) AgentSetStatus(ctx context.Context, agent entity.Agent, mode entity.Mode) error {
	ret := _m.Called(ctx, agent, mode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Agent, entity.Mode) error); ok {
		r0 = rf(ctx, agent, mode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status_AgentSetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AgentSetStatus'
type Status_AgentSetStatus_Call struct {
	*mock.Call
}

// AgentSetStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - agent entity.Agent
//   - mode entity.Mode
func (_e *Status_Expecter) AgentSetStatus(ctx interface{}, agent interface{}, mode interface{}) *Status_AgentSetStatus_Call {
	return &Status_AgentSetStatus_Call{Call: _e.mock.On("AgentSetStatus", ctx, agent, mode)}
}

func (_c *Status_AgentSetStatus_Call) Run(run func(ctx context.Context, agent entity.Agent, mode entity.Mode)) *Status_AgentSetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Agent), args[2].(entity.Mode))
	})
	return _c
}

func (_c *Status_AgentSetStatus_Call) Return(_a0 error) *Status_AgentSetStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewStatus interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatus creates a new instance of Status. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatus(t mockConstructorTestingTNewStatus) *Status {
	mock := &Status{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}