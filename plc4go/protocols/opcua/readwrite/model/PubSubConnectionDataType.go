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

// PubSubConnectionDataType is the corresponding interface of PubSubConnectionDataType
type PubSubConnectionDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// GetPublisherId returns PublisherId (property field)
	GetPublisherId() Variant
	// GetTransportProfileUri returns TransportProfileUri (property field)
	GetTransportProfileUri() PascalString
	// GetAddress returns Address (property field)
	GetAddress() ExtensionObject
	// GetNoOfConnectionProperties returns NoOfConnectionProperties (property field)
	GetNoOfConnectionProperties() int32
	// GetConnectionProperties returns ConnectionProperties (property field)
	GetConnectionProperties() []ExtensionObjectDefinition
	// GetTransportSettings returns TransportSettings (property field)
	GetTransportSettings() ExtensionObject
	// GetNoOfWriterGroups returns NoOfWriterGroups (property field)
	GetNoOfWriterGroups() int32
	// GetWriterGroups returns WriterGroups (property field)
	GetWriterGroups() []PubSubGroupDataType
	// GetNoOfReaderGroups returns NoOfReaderGroups (property field)
	GetNoOfReaderGroups() int32
	// GetReaderGroups returns ReaderGroups (property field)
	GetReaderGroups() []PubSubGroupDataType
}

// PubSubConnectionDataTypeExactly can be used when we want exactly this type and not a type which fulfills PubSubConnectionDataType.
// This is useful for switch cases.
type PubSubConnectionDataTypeExactly interface {
	PubSubConnectionDataType
	isPubSubConnectionDataType() bool
}

// _PubSubConnectionDataType is the data-structure of this message
type _PubSubConnectionDataType struct {
	*_ExtensionObjectDefinition
	Name                     PascalString
	Enabled                  bool
	PublisherId              Variant
	TransportProfileUri      PascalString
	Address                  ExtensionObject
	NoOfConnectionProperties int32
	ConnectionProperties     []ExtensionObjectDefinition
	TransportSettings        ExtensionObject
	NoOfWriterGroups         int32
	WriterGroups             []PubSubGroupDataType
	NoOfReaderGroups         int32
	ReaderGroups             []PubSubGroupDataType
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubConnectionDataType) GetIdentifier() string {
	return "15619"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubConnectionDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PubSubConnectionDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubConnectionDataType) GetName() PascalString {
	return m.Name
}

func (m *_PubSubConnectionDataType) GetEnabled() bool {
	return m.Enabled
}

func (m *_PubSubConnectionDataType) GetPublisherId() Variant {
	return m.PublisherId
}

func (m *_PubSubConnectionDataType) GetTransportProfileUri() PascalString {
	return m.TransportProfileUri
}

func (m *_PubSubConnectionDataType) GetAddress() ExtensionObject {
	return m.Address
}

func (m *_PubSubConnectionDataType) GetNoOfConnectionProperties() int32 {
	return m.NoOfConnectionProperties
}

func (m *_PubSubConnectionDataType) GetConnectionProperties() []ExtensionObjectDefinition {
	return m.ConnectionProperties
}

func (m *_PubSubConnectionDataType) GetTransportSettings() ExtensionObject {
	return m.TransportSettings
}

func (m *_PubSubConnectionDataType) GetNoOfWriterGroups() int32 {
	return m.NoOfWriterGroups
}

func (m *_PubSubConnectionDataType) GetWriterGroups() []PubSubGroupDataType {
	return m.WriterGroups
}

func (m *_PubSubConnectionDataType) GetNoOfReaderGroups() int32 {
	return m.NoOfReaderGroups
}

func (m *_PubSubConnectionDataType) GetReaderGroups() []PubSubGroupDataType {
	return m.ReaderGroups
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPubSubConnectionDataType factory function for _PubSubConnectionDataType
func NewPubSubConnectionDataType(name PascalString, enabled bool, publisherId Variant, transportProfileUri PascalString, address ExtensionObject, noOfConnectionProperties int32, connectionProperties []ExtensionObjectDefinition, transportSettings ExtensionObject, noOfWriterGroups int32, writerGroups []PubSubGroupDataType, noOfReaderGroups int32, readerGroups []PubSubGroupDataType) *_PubSubConnectionDataType {
	_result := &_PubSubConnectionDataType{
		Name:                       name,
		Enabled:                    enabled,
		PublisherId:                publisherId,
		TransportProfileUri:        transportProfileUri,
		Address:                    address,
		NoOfConnectionProperties:   noOfConnectionProperties,
		ConnectionProperties:       connectionProperties,
		TransportSettings:          transportSettings,
		NoOfWriterGroups:           noOfWriterGroups,
		WriterGroups:               writerGroups,
		NoOfReaderGroups:           noOfReaderGroups,
		ReaderGroups:               readerGroups,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPubSubConnectionDataType(structType any) PubSubConnectionDataType {
	if casted, ok := structType.(PubSubConnectionDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubConnectionDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubConnectionDataType) GetTypeName() string {
	return "PubSubConnectionDataType"
}

func (m *_PubSubConnectionDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	// Simple field (publisherId)
	lengthInBits += m.PublisherId.GetLengthInBits(ctx)

	// Simple field (transportProfileUri)
	lengthInBits += m.TransportProfileUri.GetLengthInBits(ctx)

	// Simple field (address)
	lengthInBits += m.Address.GetLengthInBits(ctx)

	// Simple field (noOfConnectionProperties)
	lengthInBits += 32

	// Array field
	if len(m.ConnectionProperties) > 0 {
		for _curItem, element := range m.ConnectionProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ConnectionProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (transportSettings)
	lengthInBits += m.TransportSettings.GetLengthInBits(ctx)

	// Simple field (noOfWriterGroups)
	lengthInBits += 32

	// Array field
	if len(m.WriterGroups) > 0 {
		for _curItem, element := range m.WriterGroups {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.WriterGroups), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfReaderGroups)
	lengthInBits += 32

	// Array field
	if len(m.ReaderGroups) > 0 {
		for _curItem, element := range m.ReaderGroups {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReaderGroups), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PubSubConnectionDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PubSubConnectionDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (PubSubConnectionDataType, error) {
	return PubSubConnectionDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PubSubConnectionDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PubSubConnectionDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PubSubConnectionDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubConnectionDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of PubSubConnectionDataType")
	}
	name := _name.(PascalString)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of PubSubConnectionDataType")
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

	// Simple Field (enabled)
	_enabled, _enabledErr := readBuffer.ReadBit("enabled")
	if _enabledErr != nil {
		return nil, errors.Wrap(_enabledErr, "Error parsing 'enabled' field of PubSubConnectionDataType")
	}
	enabled := _enabled

	// Simple Field (publisherId)
	if pullErr := readBuffer.PullContext("publisherId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for publisherId")
	}
	_publisherId, _publisherIdErr := VariantParseWithBuffer(ctx, readBuffer)
	if _publisherIdErr != nil {
		return nil, errors.Wrap(_publisherIdErr, "Error parsing 'publisherId' field of PubSubConnectionDataType")
	}
	publisherId := _publisherId.(Variant)
	if closeErr := readBuffer.CloseContext("publisherId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for publisherId")
	}

	// Simple Field (transportProfileUri)
	if pullErr := readBuffer.PullContext("transportProfileUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportProfileUri")
	}
	_transportProfileUri, _transportProfileUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _transportProfileUriErr != nil {
		return nil, errors.Wrap(_transportProfileUriErr, "Error parsing 'transportProfileUri' field of PubSubConnectionDataType")
	}
	transportProfileUri := _transportProfileUri.(PascalString)
	if closeErr := readBuffer.CloseContext("transportProfileUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportProfileUri")
	}

	// Simple Field (address)
	if pullErr := readBuffer.PullContext("address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for address")
	}
	_address, _addressErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of PubSubConnectionDataType")
	}
	address := _address.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for address")
	}

	// Simple Field (noOfConnectionProperties)
	_noOfConnectionProperties, _noOfConnectionPropertiesErr := readBuffer.ReadInt32("noOfConnectionProperties", 32)
	if _noOfConnectionPropertiesErr != nil {
		return nil, errors.Wrap(_noOfConnectionPropertiesErr, "Error parsing 'noOfConnectionProperties' field of PubSubConnectionDataType")
	}
	noOfConnectionProperties := _noOfConnectionProperties

	// Array field (connectionProperties)
	if pullErr := readBuffer.PullContext("connectionProperties", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for connectionProperties")
	}
	// Count array
	connectionProperties := make([]ExtensionObjectDefinition, utils.Max(noOfConnectionProperties, 0))
	// This happens when the size is set conditional to 0
	if len(connectionProperties) == 0 {
		connectionProperties = nil
	}
	{
		_numItems := uint16(utils.Max(noOfConnectionProperties, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "14535")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'connectionProperties' field of PubSubConnectionDataType")
			}
			connectionProperties[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("connectionProperties", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for connectionProperties")
	}

	// Simple Field (transportSettings)
	if pullErr := readBuffer.PullContext("transportSettings"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSettings")
	}
	_transportSettings, _transportSettingsErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _transportSettingsErr != nil {
		return nil, errors.Wrap(_transportSettingsErr, "Error parsing 'transportSettings' field of PubSubConnectionDataType")
	}
	transportSettings := _transportSettings.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("transportSettings"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSettings")
	}

	// Simple Field (noOfWriterGroups)
	_noOfWriterGroups, _noOfWriterGroupsErr := readBuffer.ReadInt32("noOfWriterGroups", 32)
	if _noOfWriterGroupsErr != nil {
		return nil, errors.Wrap(_noOfWriterGroupsErr, "Error parsing 'noOfWriterGroups' field of PubSubConnectionDataType")
	}
	noOfWriterGroups := _noOfWriterGroups

	// Array field (writerGroups)
	if pullErr := readBuffer.PullContext("writerGroups", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for writerGroups")
	}
	// Count array
	writerGroups := make([]PubSubGroupDataType, utils.Max(noOfWriterGroups, 0))
	// This happens when the size is set conditional to 0
	if len(writerGroups) == 0 {
		writerGroups = nil
	}
	{
		_numItems := uint16(utils.Max(noOfWriterGroups, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "15609")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'writerGroups' field of PubSubConnectionDataType")
			}
			writerGroups[_curItem] = _item.(PubSubGroupDataType)
		}
	}
	if closeErr := readBuffer.CloseContext("writerGroups", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for writerGroups")
	}

	// Simple Field (noOfReaderGroups)
	_noOfReaderGroups, _noOfReaderGroupsErr := readBuffer.ReadInt32("noOfReaderGroups", 32)
	if _noOfReaderGroupsErr != nil {
		return nil, errors.Wrap(_noOfReaderGroupsErr, "Error parsing 'noOfReaderGroups' field of PubSubConnectionDataType")
	}
	noOfReaderGroups := _noOfReaderGroups

	// Array field (readerGroups)
	if pullErr := readBuffer.PullContext("readerGroups", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for readerGroups")
	}
	// Count array
	readerGroups := make([]PubSubGroupDataType, utils.Max(noOfReaderGroups, 0))
	// This happens when the size is set conditional to 0
	if len(readerGroups) == 0 {
		readerGroups = nil
	}
	{
		_numItems := uint16(utils.Max(noOfReaderGroups, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "15609")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'readerGroups' field of PubSubConnectionDataType")
			}
			readerGroups[_curItem] = _item.(PubSubGroupDataType)
		}
	}
	if closeErr := readBuffer.CloseContext("readerGroups", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for readerGroups")
	}

	if closeErr := readBuffer.CloseContext("PubSubConnectionDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubConnectionDataType")
	}

	// Create a partially initialized instance
	_child := &_PubSubConnectionDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Name:                       name,
		Enabled:                    enabled,
		PublisherId:                publisherId,
		TransportProfileUri:        transportProfileUri,
		Address:                    address,
		NoOfConnectionProperties:   noOfConnectionProperties,
		ConnectionProperties:       connectionProperties,
		TransportSettings:          transportSettings,
		NoOfWriterGroups:           noOfWriterGroups,
		WriterGroups:               writerGroups,
		NoOfReaderGroups:           noOfReaderGroups,
		ReaderGroups:               readerGroups,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PubSubConnectionDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubConnectionDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubConnectionDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubConnectionDataType")
		}

		// Simple Field (name)
		if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for name")
		}
		_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
		if popErr := writeBuffer.PopContext("name"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for name")
		}
		if _nameErr != nil {
			return errors.Wrap(_nameErr, "Error serializing 'name' field")
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

		// Simple Field (enabled)
		enabled := bool(m.GetEnabled())
		_enabledErr := writeBuffer.WriteBit("enabled", (enabled))
		if _enabledErr != nil {
			return errors.Wrap(_enabledErr, "Error serializing 'enabled' field")
		}

		// Simple Field (publisherId)
		if pushErr := writeBuffer.PushContext("publisherId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for publisherId")
		}
		_publisherIdErr := writeBuffer.WriteSerializable(ctx, m.GetPublisherId())
		if popErr := writeBuffer.PopContext("publisherId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for publisherId")
		}
		if _publisherIdErr != nil {
			return errors.Wrap(_publisherIdErr, "Error serializing 'publisherId' field")
		}

		// Simple Field (transportProfileUri)
		if pushErr := writeBuffer.PushContext("transportProfileUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportProfileUri")
		}
		_transportProfileUriErr := writeBuffer.WriteSerializable(ctx, m.GetTransportProfileUri())
		if popErr := writeBuffer.PopContext("transportProfileUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportProfileUri")
		}
		if _transportProfileUriErr != nil {
			return errors.Wrap(_transportProfileUriErr, "Error serializing 'transportProfileUri' field")
		}

		// Simple Field (address)
		if pushErr := writeBuffer.PushContext("address"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for address")
		}
		_addressErr := writeBuffer.WriteSerializable(ctx, m.GetAddress())
		if popErr := writeBuffer.PopContext("address"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for address")
		}
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (noOfConnectionProperties)
		noOfConnectionProperties := int32(m.GetNoOfConnectionProperties())
		_noOfConnectionPropertiesErr := writeBuffer.WriteInt32("noOfConnectionProperties", 32, (noOfConnectionProperties))
		if _noOfConnectionPropertiesErr != nil {
			return errors.Wrap(_noOfConnectionPropertiesErr, "Error serializing 'noOfConnectionProperties' field")
		}

		// Array Field (connectionProperties)
		if pushErr := writeBuffer.PushContext("connectionProperties", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for connectionProperties")
		}
		for _curItem, _element := range m.GetConnectionProperties() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetConnectionProperties()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'connectionProperties' field")
			}
		}
		if popErr := writeBuffer.PopContext("connectionProperties", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for connectionProperties")
		}

		// Simple Field (transportSettings)
		if pushErr := writeBuffer.PushContext("transportSettings"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportSettings")
		}
		_transportSettingsErr := writeBuffer.WriteSerializable(ctx, m.GetTransportSettings())
		if popErr := writeBuffer.PopContext("transportSettings"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportSettings")
		}
		if _transportSettingsErr != nil {
			return errors.Wrap(_transportSettingsErr, "Error serializing 'transportSettings' field")
		}

		// Simple Field (noOfWriterGroups)
		noOfWriterGroups := int32(m.GetNoOfWriterGroups())
		_noOfWriterGroupsErr := writeBuffer.WriteInt32("noOfWriterGroups", 32, (noOfWriterGroups))
		if _noOfWriterGroupsErr != nil {
			return errors.Wrap(_noOfWriterGroupsErr, "Error serializing 'noOfWriterGroups' field")
		}

		// Array Field (writerGroups)
		if pushErr := writeBuffer.PushContext("writerGroups", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for writerGroups")
		}
		for _curItem, _element := range m.GetWriterGroups() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetWriterGroups()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'writerGroups' field")
			}
		}
		if popErr := writeBuffer.PopContext("writerGroups", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for writerGroups")
		}

		// Simple Field (noOfReaderGroups)
		noOfReaderGroups := int32(m.GetNoOfReaderGroups())
		_noOfReaderGroupsErr := writeBuffer.WriteInt32("noOfReaderGroups", 32, (noOfReaderGroups))
		if _noOfReaderGroupsErr != nil {
			return errors.Wrap(_noOfReaderGroupsErr, "Error serializing 'noOfReaderGroups' field")
		}

		// Array Field (readerGroups)
		if pushErr := writeBuffer.PushContext("readerGroups", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for readerGroups")
		}
		for _curItem, _element := range m.GetReaderGroups() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetReaderGroups()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'readerGroups' field")
			}
		}
		if popErr := writeBuffer.PopContext("readerGroups", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for readerGroups")
		}

		if popErr := writeBuffer.PopContext("PubSubConnectionDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubConnectionDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubConnectionDataType) isPubSubConnectionDataType() bool {
	return true
}

func (m *_PubSubConnectionDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
