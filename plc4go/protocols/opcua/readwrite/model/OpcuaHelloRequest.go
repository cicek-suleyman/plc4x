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

// OpcuaHelloRequest is the corresponding interface of OpcuaHelloRequest
type OpcuaHelloRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetChunk returns Chunk (property field)
	GetChunk() string
	// GetVersion returns Version (property field)
	GetVersion() int32
	// GetReceiveBufferSize returns ReceiveBufferSize (property field)
	GetReceiveBufferSize() int32
	// GetSendBufferSize returns SendBufferSize (property field)
	GetSendBufferSize() int32
	// GetMaxMessageSize returns MaxMessageSize (property field)
	GetMaxMessageSize() int32
	// GetMaxChunkCount returns MaxChunkCount (property field)
	GetMaxChunkCount() int32
	// GetEndpoint returns Endpoint (property field)
	GetEndpoint() PascalString
}

// OpcuaHelloRequestExactly can be used when we want exactly this type and not a type which fulfills OpcuaHelloRequest.
// This is useful for switch cases.
type OpcuaHelloRequestExactly interface {
	OpcuaHelloRequest
	isOpcuaHelloRequest() bool
}

// _OpcuaHelloRequest is the data-structure of this message
type _OpcuaHelloRequest struct {
	*_MessagePDU
	Chunk             string
	Version           int32
	ReceiveBufferSize int32
	SendBufferSize    int32
	MaxMessageSize    int32
	MaxChunkCount     int32
	Endpoint          PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaHelloRequest) GetMessageType() string {
	return "HEL"
}

func (m *_OpcuaHelloRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaHelloRequest) InitializeParent(parent MessagePDU) {}

func (m *_OpcuaHelloRequest) GetParent() MessagePDU {
	return m._MessagePDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaHelloRequest) GetChunk() string {
	return m.Chunk
}

func (m *_OpcuaHelloRequest) GetVersion() int32 {
	return m.Version
}

func (m *_OpcuaHelloRequest) GetReceiveBufferSize() int32 {
	return m.ReceiveBufferSize
}

func (m *_OpcuaHelloRequest) GetSendBufferSize() int32 {
	return m.SendBufferSize
}

func (m *_OpcuaHelloRequest) GetMaxMessageSize() int32 {
	return m.MaxMessageSize
}

func (m *_OpcuaHelloRequest) GetMaxChunkCount() int32 {
	return m.MaxChunkCount
}

func (m *_OpcuaHelloRequest) GetEndpoint() PascalString {
	return m.Endpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaHelloRequest factory function for _OpcuaHelloRequest
func NewOpcuaHelloRequest(chunk string, version int32, receiveBufferSize int32, sendBufferSize int32, maxMessageSize int32, maxChunkCount int32, endpoint PascalString) *_OpcuaHelloRequest {
	_result := &_OpcuaHelloRequest{
		Chunk:             chunk,
		Version:           version,
		ReceiveBufferSize: receiveBufferSize,
		SendBufferSize:    sendBufferSize,
		MaxMessageSize:    maxMessageSize,
		MaxChunkCount:     maxChunkCount,
		Endpoint:          endpoint,
		_MessagePDU:       NewMessagePDU(),
	}
	_result._MessagePDU._MessagePDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaHelloRequest(structType any) OpcuaHelloRequest {
	if casted, ok := structType.(OpcuaHelloRequest); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaHelloRequest); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaHelloRequest) GetTypeName() string {
	return "OpcuaHelloRequest"
}

func (m *_OpcuaHelloRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (chunk)
	lengthInBits += 8

	// Implicit Field (messageSize)
	lengthInBits += 32

	// Simple field (version)
	lengthInBits += 32

	// Simple field (receiveBufferSize)
	lengthInBits += 32

	// Simple field (sendBufferSize)
	lengthInBits += 32

	// Simple field (maxMessageSize)
	lengthInBits += 32

	// Simple field (maxChunkCount)
	lengthInBits += 32

	// Simple field (endpoint)
	lengthInBits += m.Endpoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaHelloRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaHelloRequestParse(ctx context.Context, theBytes []byte, response bool) (OpcuaHelloRequest, error) {
	return OpcuaHelloRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func OpcuaHelloRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (OpcuaHelloRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("OpcuaHelloRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaHelloRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (chunk)
	_chunk, _chunkErr := readBuffer.ReadString("chunk", uint32(8), "UTF-8")
	if _chunkErr != nil {
		return nil, errors.Wrap(_chunkErr, "Error parsing 'chunk' field of OpcuaHelloRequest")
	}
	chunk := _chunk

	// Implicit Field (messageSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	messageSize, _messageSizeErr := readBuffer.ReadInt32("messageSize", 32)
	_ = messageSize
	if _messageSizeErr != nil {
		return nil, errors.Wrap(_messageSizeErr, "Error parsing 'messageSize' field of OpcuaHelloRequest")
	}

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadInt32("version", 32)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field of OpcuaHelloRequest")
	}
	version := _version

	// Simple Field (receiveBufferSize)
	_receiveBufferSize, _receiveBufferSizeErr := readBuffer.ReadInt32("receiveBufferSize", 32)
	if _receiveBufferSizeErr != nil {
		return nil, errors.Wrap(_receiveBufferSizeErr, "Error parsing 'receiveBufferSize' field of OpcuaHelloRequest")
	}
	receiveBufferSize := _receiveBufferSize

	// Simple Field (sendBufferSize)
	_sendBufferSize, _sendBufferSizeErr := readBuffer.ReadInt32("sendBufferSize", 32)
	if _sendBufferSizeErr != nil {
		return nil, errors.Wrap(_sendBufferSizeErr, "Error parsing 'sendBufferSize' field of OpcuaHelloRequest")
	}
	sendBufferSize := _sendBufferSize

	// Simple Field (maxMessageSize)
	_maxMessageSize, _maxMessageSizeErr := readBuffer.ReadInt32("maxMessageSize", 32)
	if _maxMessageSizeErr != nil {
		return nil, errors.Wrap(_maxMessageSizeErr, "Error parsing 'maxMessageSize' field of OpcuaHelloRequest")
	}
	maxMessageSize := _maxMessageSize

	// Simple Field (maxChunkCount)
	_maxChunkCount, _maxChunkCountErr := readBuffer.ReadInt32("maxChunkCount", 32)
	if _maxChunkCountErr != nil {
		return nil, errors.Wrap(_maxChunkCountErr, "Error parsing 'maxChunkCount' field of OpcuaHelloRequest")
	}
	maxChunkCount := _maxChunkCount

	// Simple Field (endpoint)
	if pullErr := readBuffer.PullContext("endpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endpoint")
	}
	_endpoint, _endpointErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _endpointErr != nil {
		return nil, errors.Wrap(_endpointErr, "Error parsing 'endpoint' field of OpcuaHelloRequest")
	}
	endpoint := _endpoint.(PascalString)
	if closeErr := readBuffer.CloseContext("endpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endpoint")
	}

	if closeErr := readBuffer.CloseContext("OpcuaHelloRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaHelloRequest")
	}

	// Create a partially initialized instance
	_child := &_OpcuaHelloRequest{
		_MessagePDU:       &_MessagePDU{},
		Chunk:             chunk,
		Version:           version,
		ReceiveBufferSize: receiveBufferSize,
		SendBufferSize:    sendBufferSize,
		MaxMessageSize:    maxMessageSize,
		MaxChunkCount:     maxChunkCount,
		Endpoint:          endpoint,
	}
	_child._MessagePDU._MessagePDUChildRequirements = _child
	return _child, nil
}

func (m *_OpcuaHelloRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaHelloRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaHelloRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaHelloRequest")
		}

		// Simple Field (chunk)
		chunk := string(m.GetChunk())
		_chunkErr := writeBuffer.WriteString("chunk", uint32(8), "UTF-8", (chunk))
		if _chunkErr != nil {
			return errors.Wrap(_chunkErr, "Error serializing 'chunk' field")
		}

		// Implicit Field (messageSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		messageSize := int32(int32(m.GetLengthInBytes(ctx)))
		_messageSizeErr := writeBuffer.WriteInt32("messageSize", 32, (messageSize))
		if _messageSizeErr != nil {
			return errors.Wrap(_messageSizeErr, "Error serializing 'messageSize' field")
		}

		// Simple Field (version)
		version := int32(m.GetVersion())
		_versionErr := writeBuffer.WriteInt32("version", 32, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		// Simple Field (receiveBufferSize)
		receiveBufferSize := int32(m.GetReceiveBufferSize())
		_receiveBufferSizeErr := writeBuffer.WriteInt32("receiveBufferSize", 32, (receiveBufferSize))
		if _receiveBufferSizeErr != nil {
			return errors.Wrap(_receiveBufferSizeErr, "Error serializing 'receiveBufferSize' field")
		}

		// Simple Field (sendBufferSize)
		sendBufferSize := int32(m.GetSendBufferSize())
		_sendBufferSizeErr := writeBuffer.WriteInt32("sendBufferSize", 32, (sendBufferSize))
		if _sendBufferSizeErr != nil {
			return errors.Wrap(_sendBufferSizeErr, "Error serializing 'sendBufferSize' field")
		}

		// Simple Field (maxMessageSize)
		maxMessageSize := int32(m.GetMaxMessageSize())
		_maxMessageSizeErr := writeBuffer.WriteInt32("maxMessageSize", 32, (maxMessageSize))
		if _maxMessageSizeErr != nil {
			return errors.Wrap(_maxMessageSizeErr, "Error serializing 'maxMessageSize' field")
		}

		// Simple Field (maxChunkCount)
		maxChunkCount := int32(m.GetMaxChunkCount())
		_maxChunkCountErr := writeBuffer.WriteInt32("maxChunkCount", 32, (maxChunkCount))
		if _maxChunkCountErr != nil {
			return errors.Wrap(_maxChunkCountErr, "Error serializing 'maxChunkCount' field")
		}

		// Simple Field (endpoint)
		if pushErr := writeBuffer.PushContext("endpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for endpoint")
		}
		_endpointErr := writeBuffer.WriteSerializable(ctx, m.GetEndpoint())
		if popErr := writeBuffer.PopContext("endpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for endpoint")
		}
		if _endpointErr != nil {
			return errors.Wrap(_endpointErr, "Error serializing 'endpoint' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaHelloRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaHelloRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaHelloRequest) isOpcuaHelloRequest() bool {
	return true
}

func (m *_OpcuaHelloRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
