// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	mock "github.com/stretchr/testify/mock"

	spec "github.com/kyma-project/kyma/components/console-backend-service/internal/domain/rafter/spec"
)

// SpecificationGetter is an autogenerated mock type for the SpecificationGetter type
type SpecificationGetter struct {
	mock.Mock
}

// AsyncAPI provides a mock function with given fields: baseURL, name
func (_m *SpecificationGetter) AsyncAPI(baseURL string, name string) (*spec.AsyncAPISpec, error) {
	ret := _m.Called(baseURL, name)

	var r0 *spec.AsyncAPISpec
	if rf, ok := ret.Get(0).(func(string, string) *spec.AsyncAPISpec); ok {
		r0 = rf(baseURL, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spec.AsyncAPISpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(baseURL, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
