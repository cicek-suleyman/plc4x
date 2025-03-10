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

// ServerState is an enum
type ServerState uint32

type IServerState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ServerState_serverStateRunning            ServerState = 0
	ServerState_serverStateFailed             ServerState = 1
	ServerState_serverStateNoConfiguration    ServerState = 2
	ServerState_serverStateSuspended          ServerState = 3
	ServerState_serverStateShutdown           ServerState = 4
	ServerState_serverStateTest               ServerState = 5
	ServerState_serverStateCommunicationFault ServerState = 6
	ServerState_serverStateUnknown            ServerState = 7
)

var ServerStateValues []ServerState

func init() {
	_ = errors.New
	ServerStateValues = []ServerState{
		ServerState_serverStateRunning,
		ServerState_serverStateFailed,
		ServerState_serverStateNoConfiguration,
		ServerState_serverStateSuspended,
		ServerState_serverStateShutdown,
		ServerState_serverStateTest,
		ServerState_serverStateCommunicationFault,
		ServerState_serverStateUnknown,
	}
}

func ServerStateByValue(value uint32) (enum ServerState, ok bool) {
	switch value {
	case 0:
		return ServerState_serverStateRunning, true
	case 1:
		return ServerState_serverStateFailed, true
	case 2:
		return ServerState_serverStateNoConfiguration, true
	case 3:
		return ServerState_serverStateSuspended, true
	case 4:
		return ServerState_serverStateShutdown, true
	case 5:
		return ServerState_serverStateTest, true
	case 6:
		return ServerState_serverStateCommunicationFault, true
	case 7:
		return ServerState_serverStateUnknown, true
	}
	return 0, false
}

func ServerStateByName(value string) (enum ServerState, ok bool) {
	switch value {
	case "serverStateRunning":
		return ServerState_serverStateRunning, true
	case "serverStateFailed":
		return ServerState_serverStateFailed, true
	case "serverStateNoConfiguration":
		return ServerState_serverStateNoConfiguration, true
	case "serverStateSuspended":
		return ServerState_serverStateSuspended, true
	case "serverStateShutdown":
		return ServerState_serverStateShutdown, true
	case "serverStateTest":
		return ServerState_serverStateTest, true
	case "serverStateCommunicationFault":
		return ServerState_serverStateCommunicationFault, true
	case "serverStateUnknown":
		return ServerState_serverStateUnknown, true
	}
	return 0, false
}

func ServerStateKnows(value uint32) bool {
	for _, typeValue := range ServerStateValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastServerState(structType any) ServerState {
	castFunc := func(typ any) ServerState {
		if sServerState, ok := typ.(ServerState); ok {
			return sServerState
		}
		return 0
	}
	return castFunc(structType)
}

func (m ServerState) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m ServerState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ServerStateParse(ctx context.Context, theBytes []byte) (ServerState, error) {
	return ServerStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ServerStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ServerState, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("ServerState", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ServerState")
	}
	if enum, ok := ServerStateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ServerState")
		return ServerState(val), nil
	} else {
		return enum, nil
	}
}

func (e ServerState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ServerState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("ServerState", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ServerState) PLC4XEnumName() string {
	switch e {
	case ServerState_serverStateRunning:
		return "serverStateRunning"
	case ServerState_serverStateFailed:
		return "serverStateFailed"
	case ServerState_serverStateNoConfiguration:
		return "serverStateNoConfiguration"
	case ServerState_serverStateSuspended:
		return "serverStateSuspended"
	case ServerState_serverStateShutdown:
		return "serverStateShutdown"
	case ServerState_serverStateTest:
		return "serverStateTest"
	case ServerState_serverStateCommunicationFault:
		return "serverStateCommunicationFault"
	case ServerState_serverStateUnknown:
		return "serverStateUnknown"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e ServerState) String() string {
	return e.PLC4XEnumName()
}
