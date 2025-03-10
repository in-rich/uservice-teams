// Code generated by mockery v2.46.0. DO NOT EDIT.

package daomocks

import (
	context "context"

	dao "github.com/in-rich/uservice-teams/pkg/dao"
	entities "github.com/in-rich/uservice-teams/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockUpdateTeamMemberRepository is an autogenerated mock type for the UpdateTeamMemberRepository type
type MockUpdateTeamMemberRepository struct {
	mock.Mock
}

type MockUpdateTeamMemberRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdateTeamMemberRepository) EXPECT() *MockUpdateTeamMemberRepository_Expecter {
	return &MockUpdateTeamMemberRepository_Expecter{mock: &_m.Mock}
}

// UpdateTeamMember provides a mock function with given fields: ctx, team, userID, data
func (_m *MockUpdateTeamMemberRepository) UpdateTeamMember(ctx context.Context, team uuid.UUID, userID string, data *dao.UpdateTeamMemberData) (*entities.TeamMember, error) {
	ret := _m.Called(ctx, team, userID, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTeamMember")
	}

	var r0 *entities.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string, *dao.UpdateTeamMemberData) (*entities.TeamMember, error)); ok {
		return rf(ctx, team, userID, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string, *dao.UpdateTeamMemberData) *entities.TeamMember); ok {
		r0 = rf(ctx, team, userID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string, *dao.UpdateTeamMemberData) error); ok {
		r1 = rf(ctx, team, userID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdateTeamMemberRepository_UpdateTeamMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTeamMember'
type MockUpdateTeamMemberRepository_UpdateTeamMember_Call struct {
	*mock.Call
}

// UpdateTeamMember is a helper method to define mock.On call
//   - ctx context.Context
//   - team uuid.UUID
//   - userID string
//   - data *dao.UpdateTeamMemberData
func (_e *MockUpdateTeamMemberRepository_Expecter) UpdateTeamMember(ctx interface{}, team interface{}, userID interface{}, data interface{}) *MockUpdateTeamMemberRepository_UpdateTeamMember_Call {
	return &MockUpdateTeamMemberRepository_UpdateTeamMember_Call{Call: _e.mock.On("UpdateTeamMember", ctx, team, userID, data)}
}

func (_c *MockUpdateTeamMemberRepository_UpdateTeamMember_Call) Run(run func(ctx context.Context, team uuid.UUID, userID string, data *dao.UpdateTeamMemberData)) *MockUpdateTeamMemberRepository_UpdateTeamMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string), args[3].(*dao.UpdateTeamMemberData))
	})
	return _c
}

func (_c *MockUpdateTeamMemberRepository_UpdateTeamMember_Call) Return(_a0 *entities.TeamMember, _a1 error) *MockUpdateTeamMemberRepository_UpdateTeamMember_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdateTeamMemberRepository_UpdateTeamMember_Call) RunAndReturn(run func(context.Context, uuid.UUID, string, *dao.UpdateTeamMemberData) (*entities.TeamMember, error)) *MockUpdateTeamMemberRepository_UpdateTeamMember_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdateTeamMemberRepository creates a new instance of MockUpdateTeamMemberRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdateTeamMemberRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdateTeamMemberRepository {
	mock := &MockUpdateTeamMemberRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
