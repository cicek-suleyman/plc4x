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

// Code generated by mockery v2.30.1. DO NOT EDIT.

package interceptors

import mock "github.com/stretchr/testify/mock"

// MockWriteRequestInterceptorExposer is an autogenerated mock type for the WriteRequestInterceptorExposer type
type MockWriteRequestInterceptorExposer struct {
	mock.Mock
}

type MockWriteRequestInterceptorExposer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWriteRequestInterceptorExposer) EXPECT() *MockWriteRequestInterceptorExposer_Expecter {
	return &MockWriteRequestInterceptorExposer_Expecter{mock: &_m.Mock}
}

// GetWriteRequestInterceptor provides a mock function with given fields:
func (_m *MockWriteRequestInterceptorExposer) GetWriteRequestInterceptor() WriteRequestInterceptor {
	ret := _m.Called()

	var r0 WriteRequestInterceptor
	if rf, ok := ret.Get(0).(func() WriteRequestInterceptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(WriteRequestInterceptor)
		}
	}

	return r0
}

// MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWriteRequestInterceptor'
type MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call struct {
	*mock.Call
}

// GetWriteRequestInterceptor is a helper method to define mock.On call
func (_e *MockWriteRequestInterceptorExposer_Expecter) GetWriteRequestInterceptor() *MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call {
	return &MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call{Call: _e.mock.On("GetWriteRequestInterceptor")}
}

func (_c *MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call) Run(run func()) *MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call) Return(_a0 WriteRequestInterceptor) *MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call) RunAndReturn(run func() WriteRequestInterceptor) *MockWriteRequestInterceptorExposer_GetWriteRequestInterceptor_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWriteRequestInterceptorExposer creates a new instance of MockWriteRequestInterceptorExposer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWriteRequestInterceptorExposer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWriteRequestInterceptorExposer {
	mock := &MockWriteRequestInterceptorExposer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
