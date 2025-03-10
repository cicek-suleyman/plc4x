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

// BACnetTimeStampSequence is the corresponding interface of BACnetTimeStampSequence
type BACnetTimeStampSequence interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimeStamp
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() BACnetContextTagUnsignedInteger
}

// BACnetTimeStampSequenceExactly can be used when we want exactly this type and not a type which fulfills BACnetTimeStampSequence.
// This is useful for switch cases.
type BACnetTimeStampSequenceExactly interface {
	BACnetTimeStampSequence
	isBACnetTimeStampSequence() bool
}

// _BACnetTimeStampSequence is the data-structure of this message
type _BACnetTimeStampSequence struct {
	*_BACnetTimeStamp
	SequenceNumber BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimeStampSequence) InitializeParent(parent BACnetTimeStamp, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimeStampSequence) GetParent() BACnetTimeStamp {
	return m._BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampSequence) GetSequenceNumber() BACnetContextTagUnsignedInteger {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampSequence factory function for _BACnetTimeStampSequence
func NewBACnetTimeStampSequence(sequenceNumber BACnetContextTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetTimeStampSequence {
	_result := &_BACnetTimeStampSequence{
		SequenceNumber:   sequenceNumber,
		_BACnetTimeStamp: NewBACnetTimeStamp(peekedTagHeader),
	}
	_result._BACnetTimeStamp._BACnetTimeStampChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampSequence(structType any) BACnetTimeStampSequence {
	if casted, ok := structType.(BACnetTimeStampSequence); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampSequence); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampSequence) GetTypeName() string {
	return "BACnetTimeStampSequence"
}

func (m *_BACnetTimeStampSequence) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (sequenceNumber)
	lengthInBits += m.SequenceNumber.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampSequence) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeStampSequenceParse(ctx context.Context, theBytes []byte) (BACnetTimeStampSequence, error) {
	return BACnetTimeStampSequenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimeStampSequenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampSequence, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetTimeStampSequence"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampSequence")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sequenceNumber)
	if pullErr := readBuffer.PullContext("sequenceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sequenceNumber")
	}
	_sequenceNumber, _sequenceNumberErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of BACnetTimeStampSequence")
	}
	sequenceNumber := _sequenceNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("sequenceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sequenceNumber")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampSequence"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampSequence")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimeStampSequence{
		_BACnetTimeStamp: &_BACnetTimeStamp{},
		SequenceNumber:   sequenceNumber,
	}
	_child._BACnetTimeStamp._BACnetTimeStampChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimeStampSequence) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampSequence) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampSequence"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampSequence")
		}

		// Simple Field (sequenceNumber)
		if pushErr := writeBuffer.PushContext("sequenceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sequenceNumber")
		}
		_sequenceNumberErr := writeBuffer.WriteSerializable(ctx, m.GetSequenceNumber())
		if popErr := writeBuffer.PopContext("sequenceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sequenceNumber")
		}
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampSequence"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimeStampSequence")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimeStampSequence) isBACnetTimeStampSequence() bool {
	return true
}

func (m *_BACnetTimeStampSequence) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
