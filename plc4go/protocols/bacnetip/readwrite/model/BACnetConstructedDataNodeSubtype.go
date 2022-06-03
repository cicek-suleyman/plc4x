/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataNodeSubtype is the data-structure of this message
type BACnetConstructedDataNodeSubtype struct {
	*BACnetConstructedData
	NodeSubType *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataNodeSubtype is the corresponding interface of BACnetConstructedDataNodeSubtype
type IBACnetConstructedDataNodeSubtype interface {
	IBACnetConstructedData
	// GetNodeSubType returns NodeSubType (property field)
	GetNodeSubType() *BACnetApplicationTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataNodeSubtype) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataNodeSubtype) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NODE_SUBTYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataNodeSubtype) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataNodeSubtype) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataNodeSubtype) GetNodeSubType() *BACnetApplicationTagCharacterString {
	return m.NodeSubType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNodeSubtype factory function for BACnetConstructedDataNodeSubtype
func NewBACnetConstructedDataNodeSubtype(nodeSubType *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataNodeSubtype {
	_result := &BACnetConstructedDataNodeSubtype{
		NodeSubType:           nodeSubType,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataNodeSubtype(structType interface{}) *BACnetConstructedDataNodeSubtype {
	if casted, ok := structType.(BACnetConstructedDataNodeSubtype); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNodeSubtype); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataNodeSubtype(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataNodeSubtype(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataNodeSubtype) GetTypeName() string {
	return "BACnetConstructedDataNodeSubtype"
}

func (m *BACnetConstructedDataNodeSubtype) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataNodeSubtype) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nodeSubType)
	lengthInBits += m.NodeSubType.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataNodeSubtype) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNodeSubtypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataNodeSubtype, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNodeSubtype"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeSubType)
	if pullErr := readBuffer.PullContext("nodeSubType"); pullErr != nil {
		return nil, pullErr
	}
	_nodeSubType, _nodeSubTypeErr := BACnetApplicationTagParse(readBuffer)
	if _nodeSubTypeErr != nil {
		return nil, errors.Wrap(_nodeSubTypeErr, "Error parsing 'nodeSubType' field")
	}
	nodeSubType := CastBACnetApplicationTagCharacterString(_nodeSubType)
	if closeErr := readBuffer.CloseContext("nodeSubType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNodeSubtype"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataNodeSubtype{
		NodeSubType:           CastBACnetApplicationTagCharacterString(nodeSubType),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataNodeSubtype) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNodeSubtype"); pushErr != nil {
			return pushErr
		}

		// Simple Field (nodeSubType)
		if pushErr := writeBuffer.PushContext("nodeSubType"); pushErr != nil {
			return pushErr
		}
		_nodeSubTypeErr := m.NodeSubType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("nodeSubType"); popErr != nil {
			return popErr
		}
		if _nodeSubTypeErr != nil {
			return errors.Wrap(_nodeSubTypeErr, "Error serializing 'nodeSubType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNodeSubtype"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataNodeSubtype) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
