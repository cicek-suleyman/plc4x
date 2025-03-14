/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.32.4. DO NOT EDIT.

package cbus

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	readwritemodel "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockCALIdentifyTag is an autogenerated mock type for the CALIdentifyTag type
type MockCALIdentifyTag struct {
	mock.Mock
}

type MockCALIdentifyTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCALIdentifyTag) EXPECT() *MockCALIdentifyTag_Expecter {
	return &MockCALIdentifyTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockCALIdentifyTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCALIdentifyTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockCALIdentifyTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockCALIdentifyTag_Expecter) GetAddressString() *MockCALIdentifyTag_GetAddressString_Call {
	return &MockCALIdentifyTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockCALIdentifyTag_GetAddressString_Call) Run(run func()) *MockCALIdentifyTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALIdentifyTag_GetAddressString_Call) Return(_a0 string) *MockCALIdentifyTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALIdentifyTag_GetAddressString_Call) RunAndReturn(run func() string) *MockCALIdentifyTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockCALIdentifyTag) GetArrayInfo() []model.ArrayInfo {
	ret := _m.Called()

	var r0 []model.ArrayInfo
	if rf, ok := ret.Get(0).(func() []model.ArrayInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ArrayInfo)
		}
	}

	return r0
}

// MockCALIdentifyTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockCALIdentifyTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockCALIdentifyTag_Expecter) GetArrayInfo() *MockCALIdentifyTag_GetArrayInfo_Call {
	return &MockCALIdentifyTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockCALIdentifyTag_GetArrayInfo_Call) Run(run func()) *MockCALIdentifyTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALIdentifyTag_GetArrayInfo_Call) Return(_a0 []model.ArrayInfo) *MockCALIdentifyTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALIdentifyTag_GetArrayInfo_Call) RunAndReturn(run func() []model.ArrayInfo) *MockCALIdentifyTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetAttribute provides a mock function with given fields:
func (_m *MockCALIdentifyTag) GetAttribute() readwritemodel.Attribute {
	ret := _m.Called()

	var r0 readwritemodel.Attribute
	if rf, ok := ret.Get(0).(func() readwritemodel.Attribute); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(readwritemodel.Attribute)
	}

	return r0
}

// MockCALIdentifyTag_GetAttribute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAttribute'
type MockCALIdentifyTag_GetAttribute_Call struct {
	*mock.Call
}

// GetAttribute is a helper method to define mock.On call
func (_e *MockCALIdentifyTag_Expecter) GetAttribute() *MockCALIdentifyTag_GetAttribute_Call {
	return &MockCALIdentifyTag_GetAttribute_Call{Call: _e.mock.On("GetAttribute")}
}

func (_c *MockCALIdentifyTag_GetAttribute_Call) Run(run func()) *MockCALIdentifyTag_GetAttribute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALIdentifyTag_GetAttribute_Call) Return(_a0 readwritemodel.Attribute) *MockCALIdentifyTag_GetAttribute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALIdentifyTag_GetAttribute_Call) RunAndReturn(run func() readwritemodel.Attribute) *MockCALIdentifyTag_GetAttribute_Call {
	_c.Call.Return(run)
	return _c
}

// GetBridgeAddresses provides a mock function with given fields:
func (_m *MockCALIdentifyTag) GetBridgeAddresses() []readwritemodel.BridgeAddress {
	ret := _m.Called()

	var r0 []readwritemodel.BridgeAddress
	if rf, ok := ret.Get(0).(func() []readwritemodel.BridgeAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]readwritemodel.BridgeAddress)
		}
	}

	return r0
}

// MockCALIdentifyTag_GetBridgeAddresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBridgeAddresses'
type MockCALIdentifyTag_GetBridgeAddresses_Call struct {
	*mock.Call
}

// GetBridgeAddresses is a helper method to define mock.On call
func (_e *MockCALIdentifyTag_Expecter) GetBridgeAddresses() *MockCALIdentifyTag_GetBridgeAddresses_Call {
	return &MockCALIdentifyTag_GetBridgeAddresses_Call{Call: _e.mock.On("GetBridgeAddresses")}
}

func (_c *MockCALIdentifyTag_GetBridgeAddresses_Call) Run(run func()) *MockCALIdentifyTag_GetBridgeAddresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALIdentifyTag_GetBridgeAddresses_Call) Return(_a0 []readwritemodel.BridgeAddress) *MockCALIdentifyTag_GetBridgeAddresses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALIdentifyTag_GetBridgeAddresses_Call) RunAndReturn(run func() []readwritemodel.BridgeAddress) *MockCALIdentifyTag_GetBridgeAddresses_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagType provides a mock function with given fields:
func (_m *MockCALIdentifyTag) GetTagType() TagType {
	ret := _m.Called()

	var r0 TagType
	if rf, ok := ret.Get(0).(func() TagType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(TagType)
	}

	return r0
}

// MockCALIdentifyTag_GetTagType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagType'
type MockCALIdentifyTag_GetTagType_Call struct {
	*mock.Call
}

// GetTagType is a helper method to define mock.On call
func (_e *MockCALIdentifyTag_Expecter) GetTagType() *MockCALIdentifyTag_GetTagType_Call {
	return &MockCALIdentifyTag_GetTagType_Call{Call: _e.mock.On("GetTagType")}
}

func (_c *MockCALIdentifyTag_GetTagType_Call) Run(run func()) *MockCALIdentifyTag_GetTagType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALIdentifyTag_GetTagType_Call) Return(_a0 TagType) *MockCALIdentifyTag_GetTagType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALIdentifyTag_GetTagType_Call) RunAndReturn(run func() TagType) *MockCALIdentifyTag_GetTagType_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnitAddress provides a mock function with given fields:
func (_m *MockCALIdentifyTag) GetUnitAddress() readwritemodel.UnitAddress {
	ret := _m.Called()

	var r0 readwritemodel.UnitAddress
	if rf, ok := ret.Get(0).(func() readwritemodel.UnitAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(readwritemodel.UnitAddress)
		}
	}

	return r0
}

// MockCALIdentifyTag_GetUnitAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnitAddress'
type MockCALIdentifyTag_GetUnitAddress_Call struct {
	*mock.Call
}

// GetUnitAddress is a helper method to define mock.On call
func (_e *MockCALIdentifyTag_Expecter) GetUnitAddress() *MockCALIdentifyTag_GetUnitAddress_Call {
	return &MockCALIdentifyTag_GetUnitAddress_Call{Call: _e.mock.On("GetUnitAddress")}
}

func (_c *MockCALIdentifyTag_GetUnitAddress_Call) Run(run func()) *MockCALIdentifyTag_GetUnitAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALIdentifyTag_GetUnitAddress_Call) Return(_a0 readwritemodel.UnitAddress) *MockCALIdentifyTag_GetUnitAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALIdentifyTag_GetUnitAddress_Call) RunAndReturn(run func() readwritemodel.UnitAddress) *MockCALIdentifyTag_GetUnitAddress_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockCALIdentifyTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockCALIdentifyTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockCALIdentifyTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockCALIdentifyTag_Expecter) GetValueType() *MockCALIdentifyTag_GetValueType_Call {
	return &MockCALIdentifyTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockCALIdentifyTag_GetValueType_Call) Run(run func()) *MockCALIdentifyTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALIdentifyTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockCALIdentifyTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALIdentifyTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockCALIdentifyTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockCALIdentifyTag) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCALIdentifyTag_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockCALIdentifyTag_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockCALIdentifyTag_Expecter) String() *MockCALIdentifyTag_String_Call {
	return &MockCALIdentifyTag_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockCALIdentifyTag_String_Call) Run(run func()) *MockCALIdentifyTag_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALIdentifyTag_String_Call) Return(_a0 string) *MockCALIdentifyTag_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALIdentifyTag_String_Call) RunAndReturn(run func() string) *MockCALIdentifyTag_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCALIdentifyTag creates a new instance of MockCALIdentifyTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCALIdentifyTag(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCALIdentifyTag {
	mock := &MockCALIdentifyTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
