// Code generated by mockery v2.46.0. DO NOT EDIT.

package daomocks

import (
	context "context"

	entities "github.com/in-rich/uservice-teams/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockGetUserRoleRepository is an autogenerated mock type for the GetUserRoleRepository type
type MockGetUserRoleRepository struct {
	mock.Mock
}

type MockGetUserRoleRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGetUserRoleRepository) EXPECT() *MockGetUserRoleRepository_Expecter {
	return &MockGetUserRoleRepository_Expecter{mock: &_m.Mock}
}

// GetUserRole provides a mock function with given fields: ctx, teamID, userID
func (_m *MockGetUserRoleRepository) GetUserRole(ctx context.Context, teamID uuid.UUID, userID string) (*entities.Role, error) {
	ret := _m.Called(ctx, teamID, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserRole")
	}

	var r0 *entities.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) (*entities.Role, error)); ok {
		return rf(ctx, teamID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) *entities.Role); ok {
		r0 = rf(ctx, teamID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) error); ok {
		r1 = rf(ctx, teamID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGetUserRoleRepository_GetUserRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserRole'
type MockGetUserRoleRepository_GetUserRole_Call struct {
	*mock.Call
}

// GetUserRole is a helper method to define mock.On call
//   - ctx context.Context
//   - teamID uuid.UUID
//   - userID string
func (_e *MockGetUserRoleRepository_Expecter) GetUserRole(ctx interface{}, teamID interface{}, userID interface{}) *MockGetUserRoleRepository_GetUserRole_Call {
	return &MockGetUserRoleRepository_GetUserRole_Call{Call: _e.mock.On("GetUserRole", ctx, teamID, userID)}
}

func (_c *MockGetUserRoleRepository_GetUserRole_Call) Run(run func(ctx context.Context, teamID uuid.UUID, userID string)) *MockGetUserRoleRepository_GetUserRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *MockGetUserRoleRepository_GetUserRole_Call) Return(_a0 *entities.Role, _a1 error) *MockGetUserRoleRepository_GetUserRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGetUserRoleRepository_GetUserRole_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) (*entities.Role, error)) *MockGetUserRoleRepository_GetUserRole_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGetUserRoleRepository creates a new instance of MockGetUserRoleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetUserRoleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetUserRoleRepository {
	mock := &MockGetUserRoleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
