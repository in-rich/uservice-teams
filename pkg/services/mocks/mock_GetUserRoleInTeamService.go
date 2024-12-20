// Code generated by mockery v2.46.0. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/in-rich/uservice-teams/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// MockGetUserRoleInTeamService is an autogenerated mock type for the GetUserRoleInTeamService type
type MockGetUserRoleInTeamService struct {
	mock.Mock
}

type MockGetUserRoleInTeamService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGetUserRoleInTeamService) EXPECT() *MockGetUserRoleInTeamService_Expecter {
	return &MockGetUserRoleInTeamService_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, in
func (_m *MockGetUserRoleInTeamService) Exec(ctx context.Context, in *models.GetUserRoleInTeamRequest) (string, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetUserRoleInTeamRequest) (string, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetUserRoleInTeamRequest) string); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.GetUserRoleInTeamRequest) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGetUserRoleInTeamService_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockGetUserRoleInTeamService_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - in *models.GetUserRoleInTeamRequest
func (_e *MockGetUserRoleInTeamService_Expecter) Exec(ctx interface{}, in interface{}) *MockGetUserRoleInTeamService_Exec_Call {
	return &MockGetUserRoleInTeamService_Exec_Call{Call: _e.mock.On("Exec", ctx, in)}
}

func (_c *MockGetUserRoleInTeamService_Exec_Call) Run(run func(ctx context.Context, in *models.GetUserRoleInTeamRequest)) *MockGetUserRoleInTeamService_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.GetUserRoleInTeamRequest))
	})
	return _c
}

func (_c *MockGetUserRoleInTeamService_Exec_Call) Return(_a0 string, _a1 error) *MockGetUserRoleInTeamService_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGetUserRoleInTeamService_Exec_Call) RunAndReturn(run func(context.Context, *models.GetUserRoleInTeamRequest) (string, error)) *MockGetUserRoleInTeamService_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGetUserRoleInTeamService creates a new instance of MockGetUserRoleInTeamService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetUserRoleInTeamService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetUserRoleInTeamService {
	mock := &MockGetUserRoleInTeamService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
