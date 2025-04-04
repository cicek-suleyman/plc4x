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

// AdsMultiRequestItemWrite is the corresponding interface of AdsMultiRequestItemWrite
type AdsMultiRequestItemWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsMultiRequestItem
	// GetItemIndexGroup returns ItemIndexGroup (property field)
	GetItemIndexGroup() uint32
	// GetItemIndexOffset returns ItemIndexOffset (property field)
	GetItemIndexOffset() uint32
	// GetItemWriteLength returns ItemWriteLength (property field)
	GetItemWriteLength() uint32
}

// AdsMultiRequestItemWriteExactly can be used when we want exactly this type and not a type which fulfills AdsMultiRequestItemWrite.
// This is useful for switch cases.
type AdsMultiRequestItemWriteExactly interface {
	AdsMultiRequestItemWrite
	isAdsMultiRequestItemWrite() bool
}

// _AdsMultiRequestItemWrite is the data-structure of this message
type _AdsMultiRequestItemWrite struct {
	*_AdsMultiRequestItem
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemWriteLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsMultiRequestItemWrite) GetIndexGroup() uint32 {
	return uint32(61569)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsMultiRequestItemWrite) InitializeParent(parent AdsMultiRequestItem) {}

func (m *_AdsMultiRequestItemWrite) GetParent() AdsMultiRequestItem {
	return m._AdsMultiRequestItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsMultiRequestItemWrite) GetItemIndexGroup() uint32 {
	return m.ItemIndexGroup
}

func (m *_AdsMultiRequestItemWrite) GetItemIndexOffset() uint32 {
	return m.ItemIndexOffset
}

func (m *_AdsMultiRequestItemWrite) GetItemWriteLength() uint32 {
	return m.ItemWriteLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsMultiRequestItemWrite factory function for _AdsMultiRequestItemWrite
func NewAdsMultiRequestItemWrite(itemIndexGroup uint32, itemIndexOffset uint32, itemWriteLength uint32) *_AdsMultiRequestItemWrite {
	_result := &_AdsMultiRequestItemWrite{
		ItemIndexGroup:       itemIndexGroup,
		ItemIndexOffset:      itemIndexOffset,
		ItemWriteLength:      itemWriteLength,
		_AdsMultiRequestItem: NewAdsMultiRequestItem(),
	}
	_result._AdsMultiRequestItem._AdsMultiRequestItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItemWrite(structType any) AdsMultiRequestItemWrite {
	if casted, ok := structType.(AdsMultiRequestItemWrite); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItemWrite); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItemWrite) GetTypeName() string {
	return "AdsMultiRequestItemWrite"
}

func (m *_AdsMultiRequestItemWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemWriteLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsMultiRequestItemWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsMultiRequestItemWriteParse(ctx context.Context, theBytes []byte, indexGroup uint32) (AdsMultiRequestItemWrite, error) {
	return AdsMultiRequestItemWriteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), indexGroup)
}

func AdsMultiRequestItemWriteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, indexGroup uint32) (AdsMultiRequestItemWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItemWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (itemIndexGroup)
	_itemIndexGroup, _itemIndexGroupErr := readBuffer.ReadUint32("itemIndexGroup", 32)
	if _itemIndexGroupErr != nil {
		return nil, errors.Wrap(_itemIndexGroupErr, "Error parsing 'itemIndexGroup' field of AdsMultiRequestItemWrite")
	}
	itemIndexGroup := _itemIndexGroup

	// Simple Field (itemIndexOffset)
	_itemIndexOffset, _itemIndexOffsetErr := readBuffer.ReadUint32("itemIndexOffset", 32)
	if _itemIndexOffsetErr != nil {
		return nil, errors.Wrap(_itemIndexOffsetErr, "Error parsing 'itemIndexOffset' field of AdsMultiRequestItemWrite")
	}
	itemIndexOffset := _itemIndexOffset

	// Simple Field (itemWriteLength)
	_itemWriteLength, _itemWriteLengthErr := readBuffer.ReadUint32("itemWriteLength", 32)
	if _itemWriteLengthErr != nil {
		return nil, errors.Wrap(_itemWriteLengthErr, "Error parsing 'itemWriteLength' field of AdsMultiRequestItemWrite")
	}
	itemWriteLength := _itemWriteLength

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItemWrite")
	}

	// Create a partially initialized instance
	_child := &_AdsMultiRequestItemWrite{
		_AdsMultiRequestItem: &_AdsMultiRequestItem{},
		ItemIndexGroup:       itemIndexGroup,
		ItemIndexOffset:      itemIndexOffset,
		ItemWriteLength:      itemWriteLength,
	}
	_child._AdsMultiRequestItem._AdsMultiRequestItemChildRequirements = _child
	return _child, nil
}

func (m *_AdsMultiRequestItemWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsMultiRequestItemWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItemWrite")
		}

		// Simple Field (itemIndexGroup)
		itemIndexGroup := uint32(m.GetItemIndexGroup())
		_itemIndexGroupErr := writeBuffer.WriteUint32("itemIndexGroup", 32, (itemIndexGroup))
		if _itemIndexGroupErr != nil {
			return errors.Wrap(_itemIndexGroupErr, "Error serializing 'itemIndexGroup' field")
		}

		// Simple Field (itemIndexOffset)
		itemIndexOffset := uint32(m.GetItemIndexOffset())
		_itemIndexOffsetErr := writeBuffer.WriteUint32("itemIndexOffset", 32, (itemIndexOffset))
		if _itemIndexOffsetErr != nil {
			return errors.Wrap(_itemIndexOffsetErr, "Error serializing 'itemIndexOffset' field")
		}

		// Simple Field (itemWriteLength)
		itemWriteLength := uint32(m.GetItemWriteLength())
		_itemWriteLengthErr := writeBuffer.WriteUint32("itemWriteLength", 32, (itemWriteLength))
		if _itemWriteLengthErr != nil {
			return errors.Wrap(_itemWriteLengthErr, "Error serializing 'itemWriteLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsMultiRequestItemWrite")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsMultiRequestItemWrite) isAdsMultiRequestItemWrite() bool {
	return true
}

func (m *_AdsMultiRequestItemWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
