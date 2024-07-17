// Code generated by mockery v2.43.2. DO NOT EDIT.

package daomocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockDeleteTeamMemberRepository is an autogenerated mock type for the DeleteTeamMemberRepository type
type MockDeleteTeamMemberRepository struct {
	mock.Mock
}

type MockDeleteTeamMemberRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeleteTeamMemberRepository) EXPECT() *MockDeleteTeamMemberRepository_Expecter {
	return &MockDeleteTeamMemberRepository_Expecter{mock: &_m.Mock}
}

// DeleteTeamMember provides a mock function with given fields: ctx, team, userID
func (_m *MockDeleteTeamMemberRepository) DeleteTeamMember(ctx context.Context, team uuid.UUID, userID string) error {
	ret := _m.Called(ctx, team, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTeamMember")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, team, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDeleteTeamMemberRepository_DeleteTeamMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTeamMember'
type MockDeleteTeamMemberRepository_DeleteTeamMember_Call struct {
	*mock.Call
}

// DeleteTeamMember is a helper method to define mock.On call
//   - ctx context.Context
//   - team uuid.UUID
//   - userID string
func (_e *MockDeleteTeamMemberRepository_Expecter) DeleteTeamMember(ctx interface{}, team interface{}, userID interface{}) *MockDeleteTeamMemberRepository_DeleteTeamMember_Call {
	return &MockDeleteTeamMemberRepository_DeleteTeamMember_Call{Call: _e.mock.On("DeleteTeamMember", ctx, team, userID)}
}

func (_c *MockDeleteTeamMemberRepository_DeleteTeamMember_Call) Run(run func(ctx context.Context, team uuid.UUID, userID string)) *MockDeleteTeamMemberRepository_DeleteTeamMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *MockDeleteTeamMemberRepository_DeleteTeamMember_Call) Return(_a0 error) *MockDeleteTeamMemberRepository_DeleteTeamMember_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeleteTeamMemberRepository_DeleteTeamMember_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) error) *MockDeleteTeamMemberRepository_DeleteTeamMember_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDeleteTeamMemberRepository creates a new instance of MockDeleteTeamMemberRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeleteTeamMemberRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeleteTeamMemberRepository {
	mock := &MockDeleteTeamMemberRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
