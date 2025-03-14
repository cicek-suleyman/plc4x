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

package model

import (
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableException is an enum
type OpcuaNodeIdServicesVariableException int32

type IOpcuaNodeIdServicesVariableException interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableException_ExceptionDeviationFormat_EnumStrings OpcuaNodeIdServicesVariableException = 7614
)

var OpcuaNodeIdServicesVariableExceptionValues []OpcuaNodeIdServicesVariableException

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableExceptionValues = []OpcuaNodeIdServicesVariableException{
		OpcuaNodeIdServicesVariableException_ExceptionDeviationFormat_EnumStrings,
	}
}

func OpcuaNodeIdServicesVariableExceptionByValue(value int32) (enum OpcuaNodeIdServicesVariableException, ok bool) {
	switch value {
	case 7614:
		return OpcuaNodeIdServicesVariableException_ExceptionDeviationFormat_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableExceptionByName(value string) (enum OpcuaNodeIdServicesVariableException, ok bool) {
	switch value {
	case "ExceptionDeviationFormat_EnumStrings":
		return OpcuaNodeIdServicesVariableException_ExceptionDeviationFormat_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableExceptionKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableExceptionValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableException(structType any) OpcuaNodeIdServicesVariableException {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableException {
		if sOpcuaNodeIdServicesVariableException, ok := typ.(OpcuaNodeIdServicesVariableException); ok {
			return sOpcuaNodeIdServicesVariableException
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableException) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableException) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableExceptionParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableException, error) {
	return OpcuaNodeIdServicesVariableExceptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableExceptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableException, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableException", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableException")
	}
	if enum, ok := OpcuaNodeIdServicesVariableExceptionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableException")
		return OpcuaNodeIdServicesVariableException(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableException) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableException) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableException", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableException) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableException_ExceptionDeviationFormat_EnumStrings:
		return "ExceptionDeviationFormat_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableException) String() string {
	return e.PLC4XEnumName()
}
