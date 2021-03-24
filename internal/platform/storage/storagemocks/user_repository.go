// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package storagemocks

import (
	context "context"

	noter "github.com/romycode/bank-manager/internal"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *UserRepository) All(ctx context.Context) []noter.UserInfo {
	ret := _m.Called(ctx)

	var r0 []noter.UserInfo
	if rf, ok := ret.Get(0).(func(context.Context) []noter.UserInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]noter.UserInfo)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *UserRepository) Delete(ctx context.Context, id string) {
	_m.Called(ctx, id)
}

// Save provides a mock function with given fields: ctx, u
func (_m *UserRepository) Save(ctx context.Context, u noter.User) {
	_m.Called(ctx, u)
}