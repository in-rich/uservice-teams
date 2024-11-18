// Code generated by mockery v2.46.0. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockDeleteTeamMemberService is an autogenerated mock type for the DeleteTeamMemberService type
type MockDeleteTeamMemberService struct {
	mock.Mock
}

type MockDeleteTeamMemberService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeleteTeamMemberService) EXPECT() *MockDeleteTeamMemberService_Expecter {
	return &MockDeleteTeamMemberService_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, teamID, memberID
func (_m *MockDeleteTeamMemberService) Exec(ctx context.Context, teamID string, memberID string) error {
	ret := _m.Called(ctx, teamID, memberID)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, teamID, memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDeleteTeamMemberService_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockDeleteTeamMemberService_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - teamID string
//   - memberID string
func (_e *MockDeleteTeamMemberService_Expecter) Exec(ctx interface{}, teamID interface{}, memberID interface{}) *MockDeleteTeamMemberService_Exec_Call {
	return &MockDeleteTeamMemberService_Exec_Call{Call: _e.mock.On("Exec", ctx, teamID, memberID)}
}

func (_c *MockDeleteTeamMemberService_Exec_Call) Run(run func(ctx context.Context, teamID string, memberID string)) *MockDeleteTeamMemberService_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockDeleteTeamMemberService_Exec_Call) Return(_a0 error) *MockDeleteTeamMemberService_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeleteTeamMemberService_Exec_Call) RunAndReturn(run func(context.Context, string, string) error) *MockDeleteTeamMemberService_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDeleteTeamMemberService creates a new instance of MockDeleteTeamMemberService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeleteTeamMemberService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeleteTeamMemberService {
	mock := &MockDeleteTeamMemberService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
