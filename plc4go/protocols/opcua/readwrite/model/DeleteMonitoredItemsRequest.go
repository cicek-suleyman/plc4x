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

// DeleteMonitoredItemsRequest is the corresponding interface of DeleteMonitoredItemsRequest
type DeleteMonitoredItemsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetNoOfMonitoredItemIds returns NoOfMonitoredItemIds (property field)
	GetNoOfMonitoredItemIds() int32
	// GetMonitoredItemIds returns MonitoredItemIds (property field)
	GetMonitoredItemIds() []uint32
}

// DeleteMonitoredItemsRequestExactly can be used when we want exactly this type and not a type which fulfills DeleteMonitoredItemsRequest.
// This is useful for switch cases.
type DeleteMonitoredItemsRequestExactly interface {
	DeleteMonitoredItemsRequest
	isDeleteMonitoredItemsRequest() bool
}

// _DeleteMonitoredItemsRequest is the data-structure of this message
type _DeleteMonitoredItemsRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader        ExtensionObjectDefinition
	SubscriptionId       uint32
	NoOfMonitoredItemIds int32
	MonitoredItemIds     []uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteMonitoredItemsRequest) GetIdentifier() string {
	return "781"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteMonitoredItemsRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DeleteMonitoredItemsRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteMonitoredItemsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_DeleteMonitoredItemsRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_DeleteMonitoredItemsRequest) GetNoOfMonitoredItemIds() int32 {
	return m.NoOfMonitoredItemIds
}

func (m *_DeleteMonitoredItemsRequest) GetMonitoredItemIds() []uint32 {
	return m.MonitoredItemIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteMonitoredItemsRequest factory function for _DeleteMonitoredItemsRequest
func NewDeleteMonitoredItemsRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, noOfMonitoredItemIds int32, monitoredItemIds []uint32) *_DeleteMonitoredItemsRequest {
	_result := &_DeleteMonitoredItemsRequest{
		RequestHeader:              requestHeader,
		SubscriptionId:             subscriptionId,
		NoOfMonitoredItemIds:       noOfMonitoredItemIds,
		MonitoredItemIds:           monitoredItemIds,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteMonitoredItemsRequest(structType any) DeleteMonitoredItemsRequest {
	if casted, ok := structType.(DeleteMonitoredItemsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteMonitoredItemsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteMonitoredItemsRequest) GetTypeName() string {
	return "DeleteMonitoredItemsRequest"
}

func (m *_DeleteMonitoredItemsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (noOfMonitoredItemIds)
	lengthInBits += 32

	// Array field
	if len(m.MonitoredItemIds) > 0 {
		lengthInBits += 32 * uint16(len(m.MonitoredItemIds))
	}

	return lengthInBits
}

func (m *_DeleteMonitoredItemsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeleteMonitoredItemsRequestParse(ctx context.Context, theBytes []byte, identifier string) (DeleteMonitoredItemsRequest, error) {
	return DeleteMonitoredItemsRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DeleteMonitoredItemsRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DeleteMonitoredItemsRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeleteMonitoredItemsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteMonitoredItemsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of DeleteMonitoredItemsRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (subscriptionId)
	_subscriptionId, _subscriptionIdErr := readBuffer.ReadUint32("subscriptionId", 32)
	if _subscriptionIdErr != nil {
		return nil, errors.Wrap(_subscriptionIdErr, "Error parsing 'subscriptionId' field of DeleteMonitoredItemsRequest")
	}
	subscriptionId := _subscriptionId

	// Simple Field (noOfMonitoredItemIds)
	_noOfMonitoredItemIds, _noOfMonitoredItemIdsErr := readBuffer.ReadInt32("noOfMonitoredItemIds", 32)
	if _noOfMonitoredItemIdsErr != nil {
		return nil, errors.Wrap(_noOfMonitoredItemIdsErr, "Error parsing 'noOfMonitoredItemIds' field of DeleteMonitoredItemsRequest")
	}
	noOfMonitoredItemIds := _noOfMonitoredItemIds

	// Array field (monitoredItemIds)
	if pullErr := readBuffer.PullContext("monitoredItemIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredItemIds")
	}
	// Count array
	monitoredItemIds := make([]uint32, utils.Max(noOfMonitoredItemIds, 0))
	// This happens when the size is set conditional to 0
	if len(monitoredItemIds) == 0 {
		monitoredItemIds = nil
	}
	{
		_numItems := uint16(utils.Max(noOfMonitoredItemIds, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint32("", 32)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'monitoredItemIds' field of DeleteMonitoredItemsRequest")
			}
			monitoredItemIds[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("monitoredItemIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredItemIds")
	}

	if closeErr := readBuffer.CloseContext("DeleteMonitoredItemsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteMonitoredItemsRequest")
	}

	// Create a partially initialized instance
	_child := &_DeleteMonitoredItemsRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		SubscriptionId:             subscriptionId,
		NoOfMonitoredItemIds:       noOfMonitoredItemIds,
		MonitoredItemIds:           monitoredItemIds,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DeleteMonitoredItemsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteMonitoredItemsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteMonitoredItemsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteMonitoredItemsRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (subscriptionId)
		subscriptionId := uint32(m.GetSubscriptionId())
		_subscriptionIdErr := writeBuffer.WriteUint32("subscriptionId", 32, (subscriptionId))
		if _subscriptionIdErr != nil {
			return errors.Wrap(_subscriptionIdErr, "Error serializing 'subscriptionId' field")
		}

		// Simple Field (noOfMonitoredItemIds)
		noOfMonitoredItemIds := int32(m.GetNoOfMonitoredItemIds())
		_noOfMonitoredItemIdsErr := writeBuffer.WriteInt32("noOfMonitoredItemIds", 32, (noOfMonitoredItemIds))
		if _noOfMonitoredItemIdsErr != nil {
			return errors.Wrap(_noOfMonitoredItemIdsErr, "Error serializing 'noOfMonitoredItemIds' field")
		}

		// Array Field (monitoredItemIds)
		if pushErr := writeBuffer.PushContext("monitoredItemIds", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for monitoredItemIds")
		}
		for _curItem, _element := range m.GetMonitoredItemIds() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint32("", 32, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'monitoredItemIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("monitoredItemIds", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for monitoredItemIds")
		}

		if popErr := writeBuffer.PopContext("DeleteMonitoredItemsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteMonitoredItemsRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteMonitoredItemsRequest) isDeleteMonitoredItemsRequest() bool {
	return true
}

func (m *_DeleteMonitoredItemsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
