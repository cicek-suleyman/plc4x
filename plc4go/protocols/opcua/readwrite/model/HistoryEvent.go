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

// HistoryEvent is the corresponding interface of HistoryEvent
type HistoryEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfEvents returns NoOfEvents (property field)
	GetNoOfEvents() int32
	// GetEvents returns Events (property field)
	GetEvents() []ExtensionObjectDefinition
}

// HistoryEventExactly can be used when we want exactly this type and not a type which fulfills HistoryEvent.
// This is useful for switch cases.
type HistoryEventExactly interface {
	HistoryEvent
	isHistoryEvent() bool
}

// _HistoryEvent is the data-structure of this message
type _HistoryEvent struct {
	*_ExtensionObjectDefinition
	NoOfEvents int32
	Events     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryEvent) GetIdentifier() string {
	return "661"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryEvent) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_HistoryEvent) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryEvent) GetNoOfEvents() int32 {
	return m.NoOfEvents
}

func (m *_HistoryEvent) GetEvents() []ExtensionObjectDefinition {
	return m.Events
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHistoryEvent factory function for _HistoryEvent
func NewHistoryEvent(noOfEvents int32, events []ExtensionObjectDefinition) *_HistoryEvent {
	_result := &_HistoryEvent{
		NoOfEvents:                 noOfEvents,
		Events:                     events,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryEvent(structType any) HistoryEvent {
	if casted, ok := structType.(HistoryEvent); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryEvent); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryEvent) GetTypeName() string {
	return "HistoryEvent"
}

func (m *_HistoryEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfEvents)
	lengthInBits += 32

	// Array field
	if len(m.Events) > 0 {
		for _curItem, element := range m.Events {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Events), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryEventParse(ctx context.Context, theBytes []byte, identifier string) (HistoryEvent, error) {
	return HistoryEventParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func HistoryEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (HistoryEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HistoryEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (noOfEvents)
	_noOfEvents, _noOfEventsErr := readBuffer.ReadInt32("noOfEvents", 32)
	if _noOfEventsErr != nil {
		return nil, errors.Wrap(_noOfEventsErr, "Error parsing 'noOfEvents' field of HistoryEvent")
	}
	noOfEvents := _noOfEvents

	// Array field (events)
	if pullErr := readBuffer.PullContext("events", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for events")
	}
	// Count array
	events := make([]ExtensionObjectDefinition, utils.Max(noOfEvents, 0))
	// This happens when the size is set conditional to 0
	if len(events) == 0 {
		events = nil
	}
	{
		_numItems := uint16(utils.Max(noOfEvents, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "922")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'events' field of HistoryEvent")
			}
			events[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("events", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for events")
	}

	if closeErr := readBuffer.CloseContext("HistoryEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryEvent")
	}

	// Create a partially initialized instance
	_child := &_HistoryEvent{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfEvents:                 noOfEvents,
		Events:                     events,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_HistoryEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryEvent")
		}

		// Simple Field (noOfEvents)
		noOfEvents := int32(m.GetNoOfEvents())
		_noOfEventsErr := writeBuffer.WriteInt32("noOfEvents", 32, (noOfEvents))
		if _noOfEventsErr != nil {
			return errors.Wrap(_noOfEventsErr, "Error serializing 'noOfEvents' field")
		}

		// Array Field (events)
		if pushErr := writeBuffer.PushContext("events", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for events")
		}
		for _curItem, _element := range m.GetEvents() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetEvents()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'events' field")
			}
		}
		if popErr := writeBuffer.PopContext("events", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for events")
		}

		if popErr := writeBuffer.PopContext("HistoryEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryEvent")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryEvent) isHistoryEvent() bool {
	return true
}

func (m *_HistoryEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
