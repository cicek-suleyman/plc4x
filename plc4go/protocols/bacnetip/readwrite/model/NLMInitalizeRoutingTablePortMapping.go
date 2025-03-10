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

// NLMInitalizeRoutingTablePortMapping is the corresponding interface of NLMInitalizeRoutingTablePortMapping
type NLMInitalizeRoutingTablePortMapping interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// GetPortId returns PortId (property field)
	GetPortId() uint8
	// GetPortInfoLength returns PortInfoLength (property field)
	GetPortInfoLength() uint8
	// GetPortInfo returns PortInfo (property field)
	GetPortInfo() []byte
}

// NLMInitalizeRoutingTablePortMappingExactly can be used when we want exactly this type and not a type which fulfills NLMInitalizeRoutingTablePortMapping.
// This is useful for switch cases.
type NLMInitalizeRoutingTablePortMappingExactly interface {
	NLMInitalizeRoutingTablePortMapping
	isNLMInitalizeRoutingTablePortMapping() bool
}

// _NLMInitalizeRoutingTablePortMapping is the data-structure of this message
type _NLMInitalizeRoutingTablePortMapping struct {
	DestinationNetworkAddress uint16
	PortId                    uint8
	PortInfoLength            uint8
	PortInfo                  []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMInitalizeRoutingTablePortMapping) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

func (m *_NLMInitalizeRoutingTablePortMapping) GetPortId() uint8 {
	return m.PortId
}

func (m *_NLMInitalizeRoutingTablePortMapping) GetPortInfoLength() uint8 {
	return m.PortInfoLength
}

func (m *_NLMInitalizeRoutingTablePortMapping) GetPortInfo() []byte {
	return m.PortInfo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMInitalizeRoutingTablePortMapping factory function for _NLMInitalizeRoutingTablePortMapping
func NewNLMInitalizeRoutingTablePortMapping(destinationNetworkAddress uint16, portId uint8, portInfoLength uint8, portInfo []byte) *_NLMInitalizeRoutingTablePortMapping {
	return &_NLMInitalizeRoutingTablePortMapping{DestinationNetworkAddress: destinationNetworkAddress, PortId: portId, PortInfoLength: portInfoLength, PortInfo: portInfo}
}

// Deprecated: use the interface for direct cast
func CastNLMInitalizeRoutingTablePortMapping(structType any) NLMInitalizeRoutingTablePortMapping {
	if casted, ok := structType.(NLMInitalizeRoutingTablePortMapping); ok {
		return casted
	}
	if casted, ok := structType.(*NLMInitalizeRoutingTablePortMapping); ok {
		return *casted
	}
	return nil
}

func (m *_NLMInitalizeRoutingTablePortMapping) GetTypeName() string {
	return "NLMInitalizeRoutingTablePortMapping"
}

func (m *_NLMInitalizeRoutingTablePortMapping) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	// Simple field (portId)
	lengthInBits += 8

	// Simple field (portInfoLength)
	lengthInBits += 8

	// Array field
	if len(m.PortInfo) > 0 {
		lengthInBits += 8 * uint16(len(m.PortInfo))
	}

	return lengthInBits
}

func (m *_NLMInitalizeRoutingTablePortMapping) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMInitalizeRoutingTablePortMappingParse(ctx context.Context, theBytes []byte) (NLMInitalizeRoutingTablePortMapping, error) {
	return NLMInitalizeRoutingTablePortMappingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NLMInitalizeRoutingTablePortMappingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NLMInitalizeRoutingTablePortMapping, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMInitalizeRoutingTablePortMapping"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMInitalizeRoutingTablePortMapping")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field of NLMInitalizeRoutingTablePortMapping")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	// Simple Field (portId)
	_portId, _portIdErr := readBuffer.ReadUint8("portId", 8)
	if _portIdErr != nil {
		return nil, errors.Wrap(_portIdErr, "Error parsing 'portId' field of NLMInitalizeRoutingTablePortMapping")
	}
	portId := _portId

	// Simple Field (portInfoLength)
	_portInfoLength, _portInfoLengthErr := readBuffer.ReadUint8("portInfoLength", 8)
	if _portInfoLengthErr != nil {
		return nil, errors.Wrap(_portInfoLengthErr, "Error parsing 'portInfoLength' field of NLMInitalizeRoutingTablePortMapping")
	}
	portInfoLength := _portInfoLength
	// Byte Array field (portInfo)
	numberOfBytesportInfo := int(portInfoLength)
	portInfo, _readArrayErr := readBuffer.ReadByteArray("portInfo", numberOfBytesportInfo)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'portInfo' field of NLMInitalizeRoutingTablePortMapping")
	}

	if closeErr := readBuffer.CloseContext("NLMInitalizeRoutingTablePortMapping"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMInitalizeRoutingTablePortMapping")
	}

	// Create the instance
	return &_NLMInitalizeRoutingTablePortMapping{
		DestinationNetworkAddress: destinationNetworkAddress,
		PortId:                    portId,
		PortInfoLength:            portInfoLength,
		PortInfo:                  portInfo,
	}, nil
}

func (m *_NLMInitalizeRoutingTablePortMapping) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMInitalizeRoutingTablePortMapping) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NLMInitalizeRoutingTablePortMapping"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLMInitalizeRoutingTablePortMapping")
	}

	// Simple Field (destinationNetworkAddress)
	destinationNetworkAddress := uint16(m.GetDestinationNetworkAddress())
	_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
	if _destinationNetworkAddressErr != nil {
		return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
	}

	// Simple Field (portId)
	portId := uint8(m.GetPortId())
	_portIdErr := writeBuffer.WriteUint8("portId", 8, (portId))
	if _portIdErr != nil {
		return errors.Wrap(_portIdErr, "Error serializing 'portId' field")
	}

	// Simple Field (portInfoLength)
	portInfoLength := uint8(m.GetPortInfoLength())
	_portInfoLengthErr := writeBuffer.WriteUint8("portInfoLength", 8, (portInfoLength))
	if _portInfoLengthErr != nil {
		return errors.Wrap(_portInfoLengthErr, "Error serializing 'portInfoLength' field")
	}

	// Array Field (portInfo)
	// Byte Array field (portInfo)
	if err := writeBuffer.WriteByteArray("portInfo", m.GetPortInfo()); err != nil {
		return errors.Wrap(err, "Error serializing 'portInfo' field")
	}

	if popErr := writeBuffer.PopContext("NLMInitalizeRoutingTablePortMapping"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLMInitalizeRoutingTablePortMapping")
	}
	return nil
}

func (m *_NLMInitalizeRoutingTablePortMapping) isNLMInitalizeRoutingTablePortMapping() bool {
	return true
}

func (m *_NLMInitalizeRoutingTablePortMapping) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
