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

// ApduDataGroupValueWrite is the corresponding interface of ApduDataGroupValueWrite
type ApduDataGroupValueWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduData
	// GetDataFirstByte returns DataFirstByte (property field)
	GetDataFirstByte() int8
	// GetData returns Data (property field)
	GetData() []byte
}

// ApduDataGroupValueWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataGroupValueWrite.
// This is useful for switch cases.
type ApduDataGroupValueWriteExactly interface {
	ApduDataGroupValueWrite
	isApduDataGroupValueWrite() bool
}

// _ApduDataGroupValueWrite is the data-structure of this message
type _ApduDataGroupValueWrite struct {
	*_ApduData
	DataFirstByte int8
	Data          []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataGroupValueWrite) GetApciType() uint8 {
	return 0x2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataGroupValueWrite) InitializeParent(parent ApduData) {}

func (m *_ApduDataGroupValueWrite) GetParent() ApduData {
	return m._ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataGroupValueWrite) GetDataFirstByte() int8 {
	return m.DataFirstByte
}

func (m *_ApduDataGroupValueWrite) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataGroupValueWrite factory function for _ApduDataGroupValueWrite
func NewApduDataGroupValueWrite(dataFirstByte int8, data []byte, dataLength uint8) *_ApduDataGroupValueWrite {
	_result := &_ApduDataGroupValueWrite{
		DataFirstByte: dataFirstByte,
		Data:          data,
		_ApduData:     NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataGroupValueWrite(structType any) ApduDataGroupValueWrite {
	if casted, ok := structType.(ApduDataGroupValueWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataGroupValueWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataGroupValueWrite) GetTypeName() string {
	return "ApduDataGroupValueWrite"
}

func (m *_ApduDataGroupValueWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (dataFirstByte)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataGroupValueWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataGroupValueWriteParse(ctx context.Context, theBytes []byte, dataLength uint8) (ApduDataGroupValueWrite, error) {
	return ApduDataGroupValueWriteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataGroupValueWriteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataGroupValueWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataGroupValueWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataGroupValueWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dataFirstByte)
	_dataFirstByte, _dataFirstByteErr := readBuffer.ReadInt8("dataFirstByte", 6)
	if _dataFirstByteErr != nil {
		return nil, errors.Wrap(_dataFirstByteErr, "Error parsing 'dataFirstByte' field of ApduDataGroupValueWrite")
	}
	dataFirstByte := _dataFirstByte
	// Byte Array field (data)
	numberOfBytesdata := int(utils.InlineIf((bool((dataLength) < (1))), func() any { return uint16(uint16(0)) }, func() any { return uint16(uint16(dataLength) - uint16(uint16(1))) }).(uint16))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ApduDataGroupValueWrite")
	}

	if closeErr := readBuffer.CloseContext("ApduDataGroupValueWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataGroupValueWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataGroupValueWrite{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
		DataFirstByte: dataFirstByte,
		Data:          data,
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataGroupValueWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataGroupValueWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataGroupValueWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataGroupValueWrite")
		}

		// Simple Field (dataFirstByte)
		dataFirstByte := int8(m.GetDataFirstByte())
		_dataFirstByteErr := writeBuffer.WriteInt8("dataFirstByte", 6, (dataFirstByte))
		if _dataFirstByteErr != nil {
			return errors.Wrap(_dataFirstByteErr, "Error serializing 'dataFirstByte' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataGroupValueWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataGroupValueWrite")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataGroupValueWrite) isApduDataGroupValueWrite() bool {
	return true
}

func (m *_ApduDataGroupValueWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
