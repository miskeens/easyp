// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Mod is an autogenerated mock type for the Mod type
type Mod struct {
	mock.Mock
}

type Mod_Expecter struct {
	mock *mock.Mock
}

func (_m *Mod) EXPECT() *Mod_Expecter {
	return &Mod_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, requestedDependency
func (_m *Mod) Get(ctx context.Context, requestedDependency string) error {
	ret := _m.Called(ctx, requestedDependency)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, requestedDependency)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mod_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Mod_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - requestedDependency string
func (_e *Mod_Expecter) Get(ctx interface{}, requestedDependency interface{}) *Mod_Get_Call {
	return &Mod_Get_Call{Call: _e.mock.On("Get", ctx, requestedDependency)}
}

func (_c *Mod_Get_Call) Run(run func(ctx context.Context, requestedDependency string)) *Mod_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Mod_Get_Call) Return(_a0 error) *Mod_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mod_Get_Call) RunAndReturn(run func(context.Context, string) error) *Mod_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewMod creates a new instance of Mod. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMod(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mod {
	mock := &Mod{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
