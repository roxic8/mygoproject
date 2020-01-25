// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/bxcodec/go-clean-arch/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, a
func (_m *Repository) Delete(ctx context.Context, a *models.User) error {
	ret := _m.Called(ctx, a)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx
func (_m *Repository) Fetch(ctx context.Context) ([]models.User, error) {
	ret := _m.Called(ctx)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context) []models.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *Repository) FindAll() ([]*models.User, error) {
	ret := _m.Called()

	var r0 []*models.User
	if rf, ok := ret.Get(0).(func() []*models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmpNumber provides a mock function with given fields: ctx, empNumber
func (_m *Repository) GetByEmpNumber(ctx context.Context, empNumber string) (models.User, error) {
	ret := _m.Called(ctx, empNumber)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, empNumber)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, empNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserName provides a mock function with given fields: ctx, userName
func (_m *Repository) GetByUserName(ctx context.Context, userName string) (models.User, error) {
	ret := _m.Called(ctx, userName)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, userName)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, a
func (_m *Repository) Store(ctx context.Context, a *models.User) (int, error) {
	ret := _m.Called(ctx, a)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) int); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ar, id
func (_m *Repository) Update(ctx context.Context, ar *models.User, id int) error {
	ret := _m.Called(ctx, ar, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User, int) error); ok {
		r0 = rf(ctx, ar, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}