// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"
	categories "go-ads-management/businesses/categories"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, categoryReq
func (_m *Repository) Create(ctx context.Context, categoryReq *categories.Domain) (categories.Domain, error) {
	ret := _m.Called(ctx, categoryReq)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 categories.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *categories.Domain) (categories.Domain, error)); ok {
		return rf(ctx, categoryReq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *categories.Domain) categories.Domain); ok {
		r0 = rf(ctx, categoryReq)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *categories.Domain) error); ok {
		r1 = rf(ctx, categoryReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]categories.Domain, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []categories.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]categories.Domain, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []categories.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]categories.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id int) (categories.Domain, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 categories.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (categories.Domain, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) categories.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, categoryReq, id
func (_m *Repository) Update(ctx context.Context, categoryReq *categories.Domain, id int) (categories.Domain, error) {
	ret := _m.Called(ctx, categoryReq, id)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 categories.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *categories.Domain, int) (categories.Domain, error)); ok {
		return rf(ctx, categoryReq, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *categories.Domain, int) categories.Domain); ok {
		r0 = rf(ctx, categoryReq, id)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *categories.Domain, int) error); ok {
		r1 = rf(ctx, categoryReq, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}