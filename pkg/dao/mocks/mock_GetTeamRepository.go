// Code generated by mockery v2.46.0. DO NOT EDIT.

package daomocks

import (
	context "context"

	entities "github.com/in-rich/uservice-teams/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockGetTeamRepository is an autogenerated mock type for the GetTeamRepository type
type MockGetTeamRepository struct {
	mock.Mock
}

type MockGetTeamRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGetTeamRepository) EXPECT() *MockGetTeamRepository_Expecter {
	return &MockGetTeamRepository_Expecter{mock: &_m.Mock}
}

// GetTeam provides a mock function with given fields: ctx, teamID
func (_m *MockGetTeamRepository) GetTeam(ctx context.Context, teamID uuid.UUID) (*entities.Team, error) {
	ret := _m.Called(ctx, teamID)

	if len(ret) == 0 {
		panic("no return value specified for GetTeam")
	}

	var r0 *entities.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entities.Team, error)); ok {
		return rf(ctx, teamID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entities.Team); ok {
		r0 = rf(ctx, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGetTeamRepository_GetTeam_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTeam'
type MockGetTeamRepository_GetTeam_Call struct {
	*mock.Call
}

// GetTeam is a helper method to define mock.On call
//   - ctx context.Context
//   - teamID uuid.UUID
func (_e *MockGetTeamRepository_Expecter) GetTeam(ctx interface{}, teamID interface{}) *MockGetTeamRepository_GetTeam_Call {
	return &MockGetTeamRepository_GetTeam_Call{Call: _e.mock.On("GetTeam", ctx, teamID)}
}

func (_c *MockGetTeamRepository_GetTeam_Call) Run(run func(ctx context.Context, teamID uuid.UUID)) *MockGetTeamRepository_GetTeam_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockGetTeamRepository_GetTeam_Call) Return(_a0 *entities.Team, _a1 error) *MockGetTeamRepository_GetTeam_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGetTeamRepository_GetTeam_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*entities.Team, error)) *MockGetTeamRepository_GetTeam_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGetTeamRepository creates a new instance of MockGetTeamRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetTeamRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetTeamRepository {
	mock := &MockGetTeamRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
