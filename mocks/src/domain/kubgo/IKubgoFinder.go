// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import kubgo "siigo.com/kubgo/src/domain/kubgo"
import mock "github.com/stretchr/testify/mock"
import uuid "dev.azure.com/SiigoDevOps/Siigo/_git/go-cqrs.git/cqrs/uuid"

// IKubgoFinder is an autogenerated mock type for the IKubgoFinder type
type IKubgoFinder struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *IKubgoFinder) Get(id uuid.UUID) chan *kubgo.KubgoResponse {
	ret := _m.Called(id)

	var r0 chan *kubgo.KubgoResponse
	if rf, ok := ret.Get(0).(func(uuid.UUID) chan *kubgo.KubgoResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *kubgo.KubgoResponse)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *IKubgoFinder) GetAll() chan kubgo.KubgosResponse {
	ret := _m.Called()

	var r0 chan kubgo.KubgosResponse
	if rf, ok := ret.Get(0).(func() chan kubgo.KubgosResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan kubgo.KubgosResponse)
		}
	}

	return r0
}
