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

// ModbusPDUReadCoilsResponse is the corresponding interface of ModbusPDUReadCoilsResponse
type ModbusPDUReadCoilsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUReadCoilsResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadCoilsResponse.
// This is useful for switch cases.
type ModbusPDUReadCoilsResponseExactly interface {
	ModbusPDUReadCoilsResponse
	isModbusPDUReadCoilsResponse() bool
}

// _ModbusPDUReadCoilsResponse is the data-structure of this message
type _ModbusPDUReadCoilsResponse struct {
	*_ModbusPDU
	Value []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadCoilsResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadCoilsResponse) GetFunctionFlag() uint8 {
	return 0x01
}

func (m *_ModbusPDUReadCoilsResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadCoilsResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadCoilsResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadCoilsResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadCoilsResponse factory function for _ModbusPDUReadCoilsResponse
func NewModbusPDUReadCoilsResponse(value []byte) *_ModbusPDUReadCoilsResponse {
	_result := &_ModbusPDUReadCoilsResponse{
		Value:      value,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadCoilsResponse(structType any) ModbusPDUReadCoilsResponse {
	if casted, ok := structType.(ModbusPDUReadCoilsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadCoilsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadCoilsResponse) GetTypeName() string {
	return "ModbusPDUReadCoilsResponse"
}

func (m *_ModbusPDUReadCoilsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadCoilsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadCoilsResponseParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReadCoilsResponse, error) {
	return ModbusPDUReadCoilsResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadCoilsResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadCoilsResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusPDUReadCoilsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadCoilsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUReadCoilsResponse")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field of ModbusPDUReadCoilsResponse")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadCoilsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadCoilsResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadCoilsResponse{
		_ModbusPDU: &_ModbusPDU{},
		Value:      value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadCoilsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadCoilsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadCoilsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadCoilsResponse")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.GetValue())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (value)
		// Byte Array field (value)
		if err := writeBuffer.WriteByteArray("value", m.GetValue()); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadCoilsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadCoilsResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadCoilsResponse) isModbusPDUReadCoilsResponse() bool {
	return true
}

func (m *_ModbusPDUReadCoilsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
