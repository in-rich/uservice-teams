// Code generated by mockery v2.46.0. DO NOT EDIT.

package daomocks

import (
	context "context"

	entities "github.com/in-rich/uservice-teams/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockSetTeamOwnerRepository is an autogenerated mock type for the SetTeamOwnerRepository type
type MockSetTeamOwnerRepository struct {
	mock.Mock
}

type MockSetTeamOwnerRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSetTeamOwnerRepository) EXPECT() *MockSetTeamOwnerRepository_Expecter {
	return &MockSetTeamOwnerRepository_Expecter{mock: &_m.Mock}
}

// SetTeamOwner provides a mock function with given fields: ctx, teamID, ownerID
func (_m *MockSetTeamOwnerRepository) SetTeamOwner(ctx context.Context, teamID uuid.UUID, ownerID string) (*entities.Team, error) {
	ret := _m.Called(ctx, teamID, ownerID)

	if len(ret) == 0 {
		panic("no return value specified for SetTeamOwner")
	}

	var r0 *entities.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) (*entities.Team, error)); ok {
		return rf(ctx, teamID, ownerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) *entities.Team); ok {
		r0 = rf(ctx, teamID, ownerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) error); ok {
		r1 = rf(ctx, teamID, ownerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSetTeamOwnerRepository_SetTeamOwner_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTeamOwner'
type MockSetTeamOwnerRepository_SetTeamOwner_Call struct {
	*mock.Call
}

// SetTeamOwner is a helper method to define mock.On call
//   - ctx context.Context
//   - teamID uuid.UUID
//   - ownerID string
func (_e *MockSetTeamOwnerRepository_Expecter) SetTeamOwner(ctx interface{}, teamID interface{}, ownerID interface{}) *MockSetTeamOwnerRepository_SetTeamOwner_Call {
	return &MockSetTeamOwnerRepository_SetTeamOwner_Call{Call: _e.mock.On("SetTeamOwner", ctx, teamID, ownerID)}
}

func (_c *MockSetTeamOwnerRepository_SetTeamOwner_Call) Run(run func(ctx context.Context, teamID uuid.UUID, ownerID string)) *MockSetTeamOwnerRepository_SetTeamOwner_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *MockSetTeamOwnerRepository_SetTeamOwner_Call) Return(_a0 *entities.Team, _a1 error) *MockSetTeamOwnerRepository_SetTeamOwner_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSetTeamOwnerRepository_SetTeamOwner_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) (*entities.Team, error)) *MockSetTeamOwnerRepository_SetTeamOwner_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSetTeamOwnerRepository creates a new instance of MockSetTeamOwnerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSetTeamOwnerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSetTeamOwnerRepository {
	mock := &MockSetTeamOwnerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
