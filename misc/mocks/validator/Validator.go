// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	internaltransferssystem "github.com/wdevarshi/InternalTransfersSystem/proto"
)

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// ValidateCreateAccountRequest provides a mock function with given fields: _a0
func (_m *Validator) ValidateCreateAccountRequest(_a0 *internaltransferssystem.CreateAccountRequest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internaltransferssystem.CreateAccountRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}