// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import emptypb "google.golang.org/protobuf/types/known/emptypb"
import kubgov1 "siigo.com/kubgo/src/api/proto/kubgo/v1"
import mock "github.com/stretchr/testify/mock"

// KubgoServiceServer is an autogenerated mock type for the KubgoServiceServer type
type KubgoServiceServer struct {
	mock.Mock
}

// AddKubgo provides a mock function with given fields: _a0, _a1
func (_m *KubgoServiceServer) AddKubgo(_a0 context.Context, _a1 *kubgov1.AddKubgoRequest) (*kubgov1.AddKubgoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *kubgov1.AddKubgoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *kubgov1.AddKubgoRequest) *kubgov1.AddKubgoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kubgov1.AddKubgoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *kubgov1.AddKubgoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteKubgo provides a mock function with given fields: _a0, _a1
func (_m *KubgoServiceServer) DeleteKubgo(_a0 context.Context, _a1 *kubgov1.DeleteKubgoRequest) (*kubgov1.DeleteKubgoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *kubgov1.DeleteKubgoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *kubgov1.DeleteKubgoRequest) *kubgov1.DeleteKubgoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kubgov1.DeleteKubgoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *kubgov1.DeleteKubgoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindKubgos provides a mock function with given fields: _a0, _a1
func (_m *KubgoServiceServer) FindKubgos(_a0 context.Context, _a1 *emptypb.Empty) (*kubgov1.FindKubgosResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *kubgov1.FindKubgosResponse
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *kubgov1.FindKubgosResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kubgov1.FindKubgosResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKubgo provides a mock function with given fields: _a0, _a1
func (_m *KubgoServiceServer) GetKubgo(_a0 context.Context, _a1 *kubgov1.GetKubgoRequest) (*kubgov1.GetKubgoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *kubgov1.GetKubgoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *kubgov1.GetKubgoRequest) *kubgov1.GetKubgoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kubgov1.GetKubgoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *kubgov1.GetKubgoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateKubgo provides a mock function with given fields: _a0, _a1
func (_m *KubgoServiceServer) UpdateKubgo(_a0 context.Context, _a1 *kubgov1.UpdateKubgoRequest) (*kubgov1.UpdateKubgoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *kubgov1.UpdateKubgoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *kubgov1.UpdateKubgoRequest) *kubgov1.UpdateKubgoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kubgov1.UpdateKubgoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *kubgov1.UpdateKubgoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
