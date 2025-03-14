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

// NodeTypeDescription is the corresponding interface of NodeTypeDescription
type NodeTypeDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetTypeDefinitionNode returns TypeDefinitionNode (property field)
	GetTypeDefinitionNode() ExpandedNodeId
	// GetIncludeSubTypes returns IncludeSubTypes (property field)
	GetIncludeSubTypes() bool
	// GetNoOfDataToReturn returns NoOfDataToReturn (property field)
	GetNoOfDataToReturn() int32
	// GetDataToReturn returns DataToReturn (property field)
	GetDataToReturn() []ExtensionObjectDefinition
}

// NodeTypeDescriptionExactly can be used when we want exactly this type and not a type which fulfills NodeTypeDescription.
// This is useful for switch cases.
type NodeTypeDescriptionExactly interface {
	NodeTypeDescription
	isNodeTypeDescription() bool
}

// _NodeTypeDescription is the data-structure of this message
type _NodeTypeDescription struct {
	*_ExtensionObjectDefinition
	TypeDefinitionNode ExpandedNodeId
	IncludeSubTypes    bool
	NoOfDataToReturn   int32
	DataToReturn       []ExtensionObjectDefinition
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeTypeDescription) GetIdentifier() string {
	return "575"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeTypeDescription) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_NodeTypeDescription) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeTypeDescription) GetTypeDefinitionNode() ExpandedNodeId {
	return m.TypeDefinitionNode
}

func (m *_NodeTypeDescription) GetIncludeSubTypes() bool {
	return m.IncludeSubTypes
}

func (m *_NodeTypeDescription) GetNoOfDataToReturn() int32 {
	return m.NoOfDataToReturn
}

func (m *_NodeTypeDescription) GetDataToReturn() []ExtensionObjectDefinition {
	return m.DataToReturn
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeTypeDescription factory function for _NodeTypeDescription
func NewNodeTypeDescription(typeDefinitionNode ExpandedNodeId, includeSubTypes bool, noOfDataToReturn int32, dataToReturn []ExtensionObjectDefinition) *_NodeTypeDescription {
	_result := &_NodeTypeDescription{
		TypeDefinitionNode:         typeDefinitionNode,
		IncludeSubTypes:            includeSubTypes,
		NoOfDataToReturn:           noOfDataToReturn,
		DataToReturn:               dataToReturn,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeTypeDescription(structType any) NodeTypeDescription {
	if casted, ok := structType.(NodeTypeDescription); ok {
		return casted
	}
	if casted, ok := structType.(*NodeTypeDescription); ok {
		return *casted
	}
	return nil
}

func (m *_NodeTypeDescription) GetTypeName() string {
	return "NodeTypeDescription"
}

func (m *_NodeTypeDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (typeDefinitionNode)
	lengthInBits += m.TypeDefinitionNode.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (includeSubTypes)
	lengthInBits += 1

	// Simple field (noOfDataToReturn)
	lengthInBits += 32

	// Array field
	if len(m.DataToReturn) > 0 {
		for _curItem, element := range m.DataToReturn {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataToReturn), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NodeTypeDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeTypeDescriptionParse(ctx context.Context, theBytes []byte, identifier string) (NodeTypeDescription, error) {
	return NodeTypeDescriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func NodeTypeDescriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (NodeTypeDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NodeTypeDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeTypeDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (typeDefinitionNode)
	if pullErr := readBuffer.PullContext("typeDefinitionNode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for typeDefinitionNode")
	}
	_typeDefinitionNode, _typeDefinitionNodeErr := ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	if _typeDefinitionNodeErr != nil {
		return nil, errors.Wrap(_typeDefinitionNodeErr, "Error parsing 'typeDefinitionNode' field of NodeTypeDescription")
	}
	typeDefinitionNode := _typeDefinitionNode.(ExpandedNodeId)
	if closeErr := readBuffer.CloseContext("typeDefinitionNode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for typeDefinitionNode")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of NodeTypeDescription")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (includeSubTypes)
	_includeSubTypes, _includeSubTypesErr := readBuffer.ReadBit("includeSubTypes")
	if _includeSubTypesErr != nil {
		return nil, errors.Wrap(_includeSubTypesErr, "Error parsing 'includeSubTypes' field of NodeTypeDescription")
	}
	includeSubTypes := _includeSubTypes

	// Simple Field (noOfDataToReturn)
	_noOfDataToReturn, _noOfDataToReturnErr := readBuffer.ReadInt32("noOfDataToReturn", 32)
	if _noOfDataToReturnErr != nil {
		return nil, errors.Wrap(_noOfDataToReturnErr, "Error parsing 'noOfDataToReturn' field of NodeTypeDescription")
	}
	noOfDataToReturn := _noOfDataToReturn

	// Array field (dataToReturn)
	if pullErr := readBuffer.PullContext("dataToReturn", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataToReturn")
	}
	// Count array
	dataToReturn := make([]ExtensionObjectDefinition, utils.Max(noOfDataToReturn, 0))
	// This happens when the size is set conditional to 0
	if len(dataToReturn) == 0 {
		dataToReturn = nil
	}
	{
		_numItems := uint16(utils.Max(noOfDataToReturn, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "572")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'dataToReturn' field of NodeTypeDescription")
			}
			dataToReturn[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("dataToReturn", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataToReturn")
	}

	if closeErr := readBuffer.CloseContext("NodeTypeDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeTypeDescription")
	}

	// Create a partially initialized instance
	_child := &_NodeTypeDescription{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		TypeDefinitionNode:         typeDefinitionNode,
		IncludeSubTypes:            includeSubTypes,
		NoOfDataToReturn:           noOfDataToReturn,
		DataToReturn:               dataToReturn,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NodeTypeDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeTypeDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeTypeDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeTypeDescription")
		}

		// Simple Field (typeDefinitionNode)
		if pushErr := writeBuffer.PushContext("typeDefinitionNode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for typeDefinitionNode")
		}
		_typeDefinitionNodeErr := writeBuffer.WriteSerializable(ctx, m.GetTypeDefinitionNode())
		if popErr := writeBuffer.PopContext("typeDefinitionNode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for typeDefinitionNode")
		}
		if _typeDefinitionNodeErr != nil {
			return errors.Wrap(_typeDefinitionNodeErr, "Error serializing 'typeDefinitionNode' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (includeSubTypes)
		includeSubTypes := bool(m.GetIncludeSubTypes())
		_includeSubTypesErr := writeBuffer.WriteBit("includeSubTypes", (includeSubTypes))
		if _includeSubTypesErr != nil {
			return errors.Wrap(_includeSubTypesErr, "Error serializing 'includeSubTypes' field")
		}

		// Simple Field (noOfDataToReturn)
		noOfDataToReturn := int32(m.GetNoOfDataToReturn())
		_noOfDataToReturnErr := writeBuffer.WriteInt32("noOfDataToReturn", 32, (noOfDataToReturn))
		if _noOfDataToReturnErr != nil {
			return errors.Wrap(_noOfDataToReturnErr, "Error serializing 'noOfDataToReturn' field")
		}

		// Array Field (dataToReturn)
		if pushErr := writeBuffer.PushContext("dataToReturn", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataToReturn")
		}
		for _curItem, _element := range m.GetDataToReturn() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDataToReturn()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'dataToReturn' field")
			}
		}
		if popErr := writeBuffer.PopContext("dataToReturn", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataToReturn")
		}

		if popErr := writeBuffer.PopContext("NodeTypeDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeTypeDescription")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeTypeDescription) isNodeTypeDescription() bool {
	return true
}

func (m *_NodeTypeDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
