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

package spi

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockPlcValueHandler is an autogenerated mock type for the PlcValueHandler type
type MockPlcValueHandler struct {
	mock.Mock
}

type MockPlcValueHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcValueHandler) EXPECT() *MockPlcValueHandler_Expecter {
	return &MockPlcValueHandler_Expecter{mock: &_m.Mock}
}

// NewPlcValue provides a mock function with given fields: tag, value
func (_m *MockPlcValueHandler) NewPlcValue(tag model.PlcTag, value interface{}) (values.PlcValue, error) {
	ret := _m.Called(tag, value)

	var r0 values.PlcValue
	var r1 error
	if rf, ok := ret.Get(0).(func(model.PlcTag, interface{}) (values.PlcValue, error)); ok {
		return rf(tag, value)
	}
	if rf, ok := ret.Get(0).(func(model.PlcTag, interface{}) values.PlcValue); ok {
		r0 = rf(tag, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(values.PlcValue)
		}
	}

	if rf, ok := ret.Get(1).(func(model.PlcTag, interface{}) error); ok {
		r1 = rf(tag, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcValueHandler_NewPlcValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewPlcValue'
type MockPlcValueHandler_NewPlcValue_Call struct {
	*mock.Call
}

// NewPlcValue is a helper method to define mock.On call
//   - tag model.PlcTag
//   - value interface{}
func (_e *MockPlcValueHandler_Expecter) NewPlcValue(tag interface{}, value interface{}) *MockPlcValueHandler_NewPlcValue_Call {
	return &MockPlcValueHandler_NewPlcValue_Call{Call: _e.mock.On("NewPlcValue", tag, value)}
}

func (_c *MockPlcValueHandler_NewPlcValue_Call) Run(run func(tag model.PlcTag, value interface{})) *MockPlcValueHandler_NewPlcValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.PlcTag), args[1].(interface{}))
	})
	return _c
}

func (_c *MockPlcValueHandler_NewPlcValue_Call) Return(_a0 values.PlcValue, _a1 error) *MockPlcValueHandler_NewPlcValue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcValueHandler_NewPlcValue_Call) RunAndReturn(run func(model.PlcTag, interface{}) (values.PlcValue, error)) *MockPlcValueHandler_NewPlcValue_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcValueHandler creates a new instance of MockPlcValueHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcValueHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcValueHandler {
	mock := &MockPlcValueHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
