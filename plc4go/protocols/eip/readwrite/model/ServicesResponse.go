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

// ServicesResponse is the corresponding interface of ServicesResponse
type ServicesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TypeId
	// GetEncapsulationProtocol returns EncapsulationProtocol (property field)
	GetEncapsulationProtocol() uint16
	// GetSupportsCIPEncapsulation returns SupportsCIPEncapsulation (property field)
	GetSupportsCIPEncapsulation() bool
	// GetSupportsUDP returns SupportsUDP (property field)
	GetSupportsUDP() bool
	// GetData returns Data (property field)
	GetData() []byte
}

// ServicesResponseExactly can be used when we want exactly this type and not a type which fulfills ServicesResponse.
// This is useful for switch cases.
type ServicesResponseExactly interface {
	ServicesResponse
	isServicesResponse() bool
}

// _ServicesResponse is the data-structure of this message
type _ServicesResponse struct {
	*_TypeId
	EncapsulationProtocol    uint16
	SupportsCIPEncapsulation bool
	SupportsUDP              bool
	Data                     []byte
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServicesResponse) GetId() uint16 {
	return 0x0100
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServicesResponse) InitializeParent(parent TypeId) {}

func (m *_ServicesResponse) GetParent() TypeId {
	return m._TypeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServicesResponse) GetEncapsulationProtocol() uint16 {
	return m.EncapsulationProtocol
}

func (m *_ServicesResponse) GetSupportsCIPEncapsulation() bool {
	return m.SupportsCIPEncapsulation
}

func (m *_ServicesResponse) GetSupportsUDP() bool {
	return m.SupportsUDP
}

func (m *_ServicesResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewServicesResponse factory function for _ServicesResponse
func NewServicesResponse(encapsulationProtocol uint16, supportsCIPEncapsulation bool, supportsUDP bool, data []byte) *_ServicesResponse {
	_result := &_ServicesResponse{
		EncapsulationProtocol:    encapsulationProtocol,
		SupportsCIPEncapsulation: supportsCIPEncapsulation,
		SupportsUDP:              supportsUDP,
		Data:                     data,
		_TypeId:                  NewTypeId(),
	}
	_result._TypeId._TypeIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastServicesResponse(structType any) ServicesResponse {
	if casted, ok := structType.(ServicesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ServicesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ServicesResponse) GetTypeName() string {
	return "ServicesResponse"
}

func (m *_ServicesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (serviceLen)
	lengthInBits += 16

	// Simple field (encapsulationProtocol)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (supportsCIPEncapsulation)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 12

	// Simple field (supportsUDP)
	lengthInBits += 1

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ServicesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ServicesResponseParse(ctx context.Context, theBytes []byte) (ServicesResponse, error) {
	return ServicesResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ServicesResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ServicesResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ServicesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServicesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (serviceLen) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	serviceLen, _serviceLenErr := readBuffer.ReadUint16("serviceLen", 16)
	_ = serviceLen
	if _serviceLenErr != nil {
		return nil, errors.Wrap(_serviceLenErr, "Error parsing 'serviceLen' field of ServicesResponse")
	}

	// Simple Field (encapsulationProtocol)
	_encapsulationProtocol, _encapsulationProtocolErr := readBuffer.ReadUint16("encapsulationProtocol", 16)
	if _encapsulationProtocolErr != nil {
		return nil, errors.Wrap(_encapsulationProtocolErr, "Error parsing 'encapsulationProtocol' field of ServicesResponse")
	}
	encapsulationProtocol := _encapsulationProtocol

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ServicesResponse")
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

	// Simple Field (supportsCIPEncapsulation)
	_supportsCIPEncapsulation, _supportsCIPEncapsulationErr := readBuffer.ReadBit("supportsCIPEncapsulation")
	if _supportsCIPEncapsulationErr != nil {
		return nil, errors.Wrap(_supportsCIPEncapsulationErr, "Error parsing 'supportsCIPEncapsulation' field of ServicesResponse")
	}
	supportsCIPEncapsulation := _supportsCIPEncapsulation

	var reservedField1 *uint16
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 12)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ServicesResponse")
		}
		if reserved != uint16(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint16(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Simple Field (supportsUDP)
	_supportsUDP, _supportsUDPErr := readBuffer.ReadBit("supportsUDP")
	if _supportsUDPErr != nil {
		return nil, errors.Wrap(_supportsUDPErr, "Error parsing 'supportsUDP' field of ServicesResponse")
	}
	supportsUDP := _supportsUDP
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(serviceLen) - uint16(uint16(4)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ServicesResponse")
	}

	if closeErr := readBuffer.CloseContext("ServicesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServicesResponse")
	}

	// Create a partially initialized instance
	_child := &_ServicesResponse{
		_TypeId:                  &_TypeId{},
		EncapsulationProtocol:    encapsulationProtocol,
		SupportsCIPEncapsulation: supportsCIPEncapsulation,
		SupportsUDP:              supportsUDP,
		Data:                     data,
		reservedField0:           reservedField0,
		reservedField1:           reservedField1,
	}
	_child._TypeId._TypeIdChildRequirements = _child
	return _child, nil
}

func (m *_ServicesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServicesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServicesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServicesResponse")
		}

		// Implicit Field (serviceLen) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		serviceLen := uint16(uint16(uint16(m.GetLengthInBytes(ctx))) - uint16(uint16(4)))
		_serviceLenErr := writeBuffer.WriteUint16("serviceLen", 16, (serviceLen))
		if _serviceLenErr != nil {
			return errors.Wrap(_serviceLenErr, "Error serializing 'serviceLen' field")
		}

		// Simple Field (encapsulationProtocol)
		encapsulationProtocol := uint16(m.GetEncapsulationProtocol())
		_encapsulationProtocolErr := writeBuffer.WriteUint16("encapsulationProtocol", 16, (encapsulationProtocol))
		if _encapsulationProtocolErr != nil {
			return errors.Wrap(_encapsulationProtocolErr, "Error serializing 'encapsulationProtocol' field")
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
			_err := writeBuffer.WriteUint8("reserved", 2, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (supportsCIPEncapsulation)
		supportsCIPEncapsulation := bool(m.GetSupportsCIPEncapsulation())
		_supportsCIPEncapsulationErr := writeBuffer.WriteBit("supportsCIPEncapsulation", (supportsCIPEncapsulation))
		if _supportsCIPEncapsulationErr != nil {
			return errors.Wrap(_supportsCIPEncapsulationErr, "Error serializing 'supportsCIPEncapsulation' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint16 = uint16(0x00)
			if m.reservedField1 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint16(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteUint16("reserved", 12, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (supportsUDP)
		supportsUDP := bool(m.GetSupportsUDP())
		_supportsUDPErr := writeBuffer.WriteBit("supportsUDP", (supportsUDP))
		if _supportsUDPErr != nil {
			return errors.Wrap(_supportsUDPErr, "Error serializing 'supportsUDP' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ServicesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServicesResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServicesResponse) isServicesResponse() bool {
	return true
}

func (m *_ServicesResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
