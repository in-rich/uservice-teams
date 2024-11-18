// Code generated by mockery v2.46.0. DO NOT EDIT.

package daomocks

import (
	context "context"

	dao "github.com/in-rich/uservice-teams/pkg/dao"
	entities "github.com/in-rich/uservice-teams/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockUpdateTeamRepository is an autogenerated mock type for the UpdateTeamRepository type
type MockUpdateTeamRepository struct {
	mock.Mock
}

type MockUpdateTeamRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdateTeamRepository) EXPECT() *MockUpdateTeamRepository_Expecter {
	return &MockUpdateTeamRepository_Expecter{mock: &_m.Mock}
}

// UpdateTeam provides a mock function with given fields: ctx, teamID, data
func (_m *MockUpdateTeamRepository) UpdateTeam(ctx context.Context, teamID uuid.UUID, data *dao.UpdateTeamData) (*entities.Team, error) {
	ret := _m.Called(ctx, teamID, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTeam")
	}

	var r0 *entities.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *dao.UpdateTeamData) (*entities.Team, error)); ok {
		return rf(ctx, teamID, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *dao.UpdateTeamData) *entities.Team); ok {
		r0 = rf(ctx, teamID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, *dao.UpdateTeamData) error); ok {
		r1 = rf(ctx, teamID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdateTeamRepository_UpdateTeam_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTeam'
type MockUpdateTeamRepository_UpdateTeam_Call struct {
	*mock.Call
}

// UpdateTeam is a helper method to define mock.On call
//   - ctx context.Context
//   - teamID uuid.UUID
//   - data *dao.UpdateTeamData
func (_e *MockUpdateTeamRepository_Expecter) UpdateTeam(ctx interface{}, teamID interface{}, data interface{}) *MockUpdateTeamRepository_UpdateTeam_Call {
	return &MockUpdateTeamRepository_UpdateTeam_Call{Call: _e.mock.On("UpdateTeam", ctx, teamID, data)}
}

func (_c *MockUpdateTeamRepository_UpdateTeam_Call) Run(run func(ctx context.Context, teamID uuid.UUID, data *dao.UpdateTeamData)) *MockUpdateTeamRepository_UpdateTeam_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(*dao.UpdateTeamData))
	})
	return _c
}

func (_c *MockUpdateTeamRepository_UpdateTeam_Call) Return(_a0 *entities.Team, _a1 error) *MockUpdateTeamRepository_UpdateTeam_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdateTeamRepository_UpdateTeam_Call) RunAndReturn(run func(context.Context, uuid.UUID, *dao.UpdateTeamData) (*entities.Team, error)) *MockUpdateTeamRepository_UpdateTeam_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdateTeamRepository creates a new instance of MockUpdateTeamRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdateTeamRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdateTeamRepository {
	mock := &MockUpdateTeamRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
