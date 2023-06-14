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

import (
	spi "github.com/apache/plc4x/plc4go/spi"
	mock "github.com/stretchr/testify/mock"
)

// MockReaderExposer is an autogenerated mock type for the ReaderExposer type
type MockReaderExposer struct {
	mock.Mock
}

type MockReaderExposer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReaderExposer) EXPECT() *MockReaderExposer_Expecter {
	return &MockReaderExposer_Expecter{mock: &_m.Mock}
}

// GetReader provides a mock function with given fields:
func (_m *MockReaderExposer) GetReader() spi.PlcReader {
	ret := _m.Called()

	var r0 spi.PlcReader
	if rf, ok := ret.Get(0).(func() spi.PlcReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.PlcReader)
		}
	}

	return r0
}

// MockReaderExposer_GetReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReader'
type MockReaderExposer_GetReader_Call struct {
	*mock.Call
}

// GetReader is a helper method to define mock.On call
func (_e *MockReaderExposer_Expecter) GetReader() *MockReaderExposer_GetReader_Call {
	return &MockReaderExposer_GetReader_Call{Call: _e.mock.On("GetReader")}
}

func (_c *MockReaderExposer_GetReader_Call) Run(run func()) *MockReaderExposer_GetReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReaderExposer_GetReader_Call) Return(_a0 spi.PlcReader) *MockReaderExposer_GetReader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReaderExposer_GetReader_Call) RunAndReturn(run func() spi.PlcReader) *MockReaderExposer_GetReader_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReaderExposer creates a new instance of MockReaderExposer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReaderExposer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReaderExposer {
	mock := &MockReaderExposer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
