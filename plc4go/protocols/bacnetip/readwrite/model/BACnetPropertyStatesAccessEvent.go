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

// BACnetPropertyStatesAccessEvent is the corresponding interface of BACnetPropertyStatesAccessEvent
type BACnetPropertyStatesAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetAccessEvent returns AccessEvent (property field)
	GetAccessEvent() BACnetAccessEventTagged
}

// BACnetPropertyStatesAccessEventExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesAccessEvent.
// This is useful for switch cases.
type BACnetPropertyStatesAccessEventExactly interface {
	BACnetPropertyStatesAccessEvent
	isBACnetPropertyStatesAccessEvent() bool
}

// _BACnetPropertyStatesAccessEvent is the data-structure of this message
type _BACnetPropertyStatesAccessEvent struct {
	*_BACnetPropertyStates
	AccessEvent BACnetAccessEventTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesAccessEvent) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesAccessEvent) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesAccessEvent) GetAccessEvent() BACnetAccessEventTagged {
	return m.AccessEvent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesAccessEvent factory function for _BACnetPropertyStatesAccessEvent
func NewBACnetPropertyStatesAccessEvent(accessEvent BACnetAccessEventTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesAccessEvent {
	_result := &_BACnetPropertyStatesAccessEvent{
		AccessEvent:           accessEvent,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesAccessEvent(structType any) BACnetPropertyStatesAccessEvent {
	if casted, ok := structType.(BACnetPropertyStatesAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesAccessEvent) GetTypeName() string {
	return "BACnetPropertyStatesAccessEvent"
}

func (m *_BACnetPropertyStatesAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (accessEvent)
	lengthInBits += m.AccessEvent.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesAccessEventParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesAccessEvent, error) {
	return BACnetPropertyStatesAccessEventParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesAccessEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesAccessEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessEvent)
	if pullErr := readBuffer.PullContext("accessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEvent")
	}
	_accessEvent, _accessEventErr := BACnetAccessEventTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _accessEventErr != nil {
		return nil, errors.Wrap(_accessEventErr, "Error parsing 'accessEvent' field of BACnetPropertyStatesAccessEvent")
	}
	accessEvent := _accessEvent.(BACnetAccessEventTagged)
	if closeErr := readBuffer.CloseContext("accessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEvent")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAccessEvent")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesAccessEvent{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		AccessEvent:           accessEvent,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAccessEvent")
		}

		// Simple Field (accessEvent)
		if pushErr := writeBuffer.PushContext("accessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEvent")
		}
		_accessEventErr := writeBuffer.WriteSerializable(ctx, m.GetAccessEvent())
		if popErr := writeBuffer.PopContext("accessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEvent")
		}
		if _accessEventErr != nil {
			return errors.Wrap(_accessEventErr, "Error serializing 'accessEvent' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAccessEvent")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesAccessEvent) isBACnetPropertyStatesAccessEvent() bool {
	return true
}

func (m *_BACnetPropertyStatesAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
