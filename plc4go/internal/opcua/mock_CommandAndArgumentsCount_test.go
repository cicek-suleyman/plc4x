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

package opcua

import mock "github.com/stretchr/testify/mock"

// MockCommandAndArgumentsCount is an autogenerated mock type for the CommandAndArgumentsCount type
type MockCommandAndArgumentsCount struct {
	mock.Mock
}

type MockCommandAndArgumentsCount_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCommandAndArgumentsCount) EXPECT() *MockCommandAndArgumentsCount_Expecter {
	return &MockCommandAndArgumentsCount_Expecter{mock: &_m.Mock}
}

// NumberOfArguments provides a mock function with given fields:
func (_m *MockCommandAndArgumentsCount) NumberOfArguments() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// MockCommandAndArgumentsCount_NumberOfArguments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NumberOfArguments'
type MockCommandAndArgumentsCount_NumberOfArguments_Call struct {
	*mock.Call
}

// NumberOfArguments is a helper method to define mock.On call
func (_e *MockCommandAndArgumentsCount_Expecter) NumberOfArguments() *MockCommandAndArgumentsCount_NumberOfArguments_Call {
	return &MockCommandAndArgumentsCount_NumberOfArguments_Call{Call: _e.mock.On("NumberOfArguments")}
}

func (_c *MockCommandAndArgumentsCount_NumberOfArguments_Call) Run(run func()) *MockCommandAndArgumentsCount_NumberOfArguments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandAndArgumentsCount_NumberOfArguments_Call) Return(_a0 uint8) *MockCommandAndArgumentsCount_NumberOfArguments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommandAndArgumentsCount_NumberOfArguments_Call) RunAndReturn(run func() uint8) *MockCommandAndArgumentsCount_NumberOfArguments_Call {
	_c.Call.Return(run)
	return _c
}

// PLC4XEnumName provides a mock function with given fields:
func (_m *MockCommandAndArgumentsCount) PLC4XEnumName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCommandAndArgumentsCount_PLC4XEnumName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PLC4XEnumName'
type MockCommandAndArgumentsCount_PLC4XEnumName_Call struct {
	*mock.Call
}

// PLC4XEnumName is a helper method to define mock.On call
func (_e *MockCommandAndArgumentsCount_Expecter) PLC4XEnumName() *MockCommandAndArgumentsCount_PLC4XEnumName_Call {
	return &MockCommandAndArgumentsCount_PLC4XEnumName_Call{Call: _e.mock.On("PLC4XEnumName")}
}

func (_c *MockCommandAndArgumentsCount_PLC4XEnumName_Call) Run(run func()) *MockCommandAndArgumentsCount_PLC4XEnumName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandAndArgumentsCount_PLC4XEnumName_Call) Return(_a0 string) *MockCommandAndArgumentsCount_PLC4XEnumName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommandAndArgumentsCount_PLC4XEnumName_Call) RunAndReturn(run func() string) *MockCommandAndArgumentsCount_PLC4XEnumName_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockCommandAndArgumentsCount) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCommandAndArgumentsCount_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockCommandAndArgumentsCount_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockCommandAndArgumentsCount_Expecter) String() *MockCommandAndArgumentsCount_String_Call {
	return &MockCommandAndArgumentsCount_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockCommandAndArgumentsCount_String_Call) Run(run func()) *MockCommandAndArgumentsCount_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandAndArgumentsCount_String_Call) Return(_a0 string) *MockCommandAndArgumentsCount_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommandAndArgumentsCount_String_Call) RunAndReturn(run func() string) *MockCommandAndArgumentsCount_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCommandAndArgumentsCount creates a new instance of MockCommandAndArgumentsCount. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommandAndArgumentsCount(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommandAndArgumentsCount {
	mock := &MockCommandAndArgumentsCount{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
