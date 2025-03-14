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

// MemberID is the corresponding interface of MemberID
type MemberID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LogicalSegmentType
	// GetFormat returns Format (property field)
	GetFormat() uint8
	// GetInstance returns Instance (property field)
	GetInstance() uint8
}

// MemberIDExactly can be used when we want exactly this type and not a type which fulfills MemberID.
// This is useful for switch cases.
type MemberIDExactly interface {
	MemberID
	isMemberID() bool
}

// _MemberID is the data-structure of this message
type _MemberID struct {
	*_LogicalSegmentType
	Format   uint8
	Instance uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MemberID) GetLogicalSegmentType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MemberID) InitializeParent(parent LogicalSegmentType) {}

func (m *_MemberID) GetParent() LogicalSegmentType {
	return m._LogicalSegmentType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MemberID) GetFormat() uint8 {
	return m.Format
}

func (m *_MemberID) GetInstance() uint8 {
	return m.Instance
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMemberID factory function for _MemberID
func NewMemberID(format uint8, instance uint8) *_MemberID {
	_result := &_MemberID{
		Format:              format,
		Instance:            instance,
		_LogicalSegmentType: NewLogicalSegmentType(),
	}
	_result._LogicalSegmentType._LogicalSegmentTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMemberID(structType any) MemberID {
	if casted, ok := structType.(MemberID); ok {
		return casted
	}
	if casted, ok := structType.(*MemberID); ok {
		return *casted
	}
	return nil
}

func (m *_MemberID) GetTypeName() string {
	return "MemberID"
}

func (m *_MemberID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (format)
	lengthInBits += 2

	// Simple field (instance)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MemberID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MemberIDParse(ctx context.Context, theBytes []byte) (MemberID, error) {
	return MemberIDParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MemberIDParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MemberID, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MemberID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MemberID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (format)
	_format, _formatErr := readBuffer.ReadUint8("format", 2)
	if _formatErr != nil {
		return nil, errors.Wrap(_formatErr, "Error parsing 'format' field of MemberID")
	}
	format := _format

	// Simple Field (instance)
	_instance, _instanceErr := readBuffer.ReadUint8("instance", 8)
	if _instanceErr != nil {
		return nil, errors.Wrap(_instanceErr, "Error parsing 'instance' field of MemberID")
	}
	instance := _instance

	if closeErr := readBuffer.CloseContext("MemberID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MemberID")
	}

	// Create a partially initialized instance
	_child := &_MemberID{
		_LogicalSegmentType: &_LogicalSegmentType{},
		Format:              format,
		Instance:            instance,
	}
	_child._LogicalSegmentType._LogicalSegmentTypeChildRequirements = _child
	return _child, nil
}

func (m *_MemberID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MemberID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MemberID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MemberID")
		}

		// Simple Field (format)
		format := uint8(m.GetFormat())
		_formatErr := writeBuffer.WriteUint8("format", 2, (format))
		if _formatErr != nil {
			return errors.Wrap(_formatErr, "Error serializing 'format' field")
		}

		// Simple Field (instance)
		instance := uint8(m.GetInstance())
		_instanceErr := writeBuffer.WriteUint8("instance", 8, (instance))
		if _instanceErr != nil {
			return errors.Wrap(_instanceErr, "Error serializing 'instance' field")
		}

		if popErr := writeBuffer.PopContext("MemberID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MemberID")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MemberID) isMemberID() bool {
	return true
}

func (m *_MemberID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
