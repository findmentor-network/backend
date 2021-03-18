package person

import (
	context "context"

	pagination "github.com/findmentor-network/backend/pkg/pagination"
	mock "github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func NewMockRepository(m mock.Mock) Repository {

	return &mockRepository{
		m,
	}
}
func (_m *mockRepository) Get(_a0 context.Context, _a1 *pagination.Pages) ([]*Person, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*Person
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pages) []*Person); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pagination.Pages) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
