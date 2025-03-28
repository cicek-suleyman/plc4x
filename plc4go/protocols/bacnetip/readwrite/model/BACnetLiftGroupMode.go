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

// BACnetLiftGroupMode is an enum
type BACnetLiftGroupMode uint8

type IBACnetLiftGroupMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetLiftGroupMode_UNKNOWN         BACnetLiftGroupMode = 0
	BACnetLiftGroupMode_NORMAL          BACnetLiftGroupMode = 1
	BACnetLiftGroupMode_DOWN_PEAK       BACnetLiftGroupMode = 2
	BACnetLiftGroupMode_TWO_WAY         BACnetLiftGroupMode = 3
	BACnetLiftGroupMode_FOUR_WAY        BACnetLiftGroupMode = 4
	BACnetLiftGroupMode_EMERGENCY_POWER BACnetLiftGroupMode = 5
	BACnetLiftGroupMode_UP_PEAK         BACnetLiftGroupMode = 6
)

var BACnetLiftGroupModeValues []BACnetLiftGroupMode

func init() {
	_ = errors.New
	BACnetLiftGroupModeValues = []BACnetLiftGroupMode{
		BACnetLiftGroupMode_UNKNOWN,
		BACnetLiftGroupMode_NORMAL,
		BACnetLiftGroupMode_DOWN_PEAK,
		BACnetLiftGroupMode_TWO_WAY,
		BACnetLiftGroupMode_FOUR_WAY,
		BACnetLiftGroupMode_EMERGENCY_POWER,
		BACnetLiftGroupMode_UP_PEAK,
	}
}

func BACnetLiftGroupModeByValue(value uint8) (enum BACnetLiftGroupMode, ok bool) {
	switch value {
	case 0:
		return BACnetLiftGroupMode_UNKNOWN, true
	case 1:
		return BACnetLiftGroupMode_NORMAL, true
	case 2:
		return BACnetLiftGroupMode_DOWN_PEAK, true
	case 3:
		return BACnetLiftGroupMode_TWO_WAY, true
	case 4:
		return BACnetLiftGroupMode_FOUR_WAY, true
	case 5:
		return BACnetLiftGroupMode_EMERGENCY_POWER, true
	case 6:
		return BACnetLiftGroupMode_UP_PEAK, true
	}
	return 0, false
}

func BACnetLiftGroupModeByName(value string) (enum BACnetLiftGroupMode, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetLiftGroupMode_UNKNOWN, true
	case "NORMAL":
		return BACnetLiftGroupMode_NORMAL, true
	case "DOWN_PEAK":
		return BACnetLiftGroupMode_DOWN_PEAK, true
	case "TWO_WAY":
		return BACnetLiftGroupMode_TWO_WAY, true
	case "FOUR_WAY":
		return BACnetLiftGroupMode_FOUR_WAY, true
	case "EMERGENCY_POWER":
		return BACnetLiftGroupMode_EMERGENCY_POWER, true
	case "UP_PEAK":
		return BACnetLiftGroupMode_UP_PEAK, true
	}
	return 0, false
}

func BACnetLiftGroupModeKnows(value uint8) bool {
	for _, typeValue := range BACnetLiftGroupModeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLiftGroupMode(structType any) BACnetLiftGroupMode {
	castFunc := func(typ any) BACnetLiftGroupMode {
		if sBACnetLiftGroupMode, ok := typ.(BACnetLiftGroupMode); ok {
			return sBACnetLiftGroupMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLiftGroupMode) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetLiftGroupMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftGroupModeParse(ctx context.Context, theBytes []byte) (BACnetLiftGroupMode, error) {
	return BACnetLiftGroupModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLiftGroupModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftGroupMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetLiftGroupMode", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLiftGroupMode")
	}
	if enum, ok := BACnetLiftGroupModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetLiftGroupMode")
		return BACnetLiftGroupMode(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLiftGroupMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLiftGroupMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetLiftGroupMode", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLiftGroupMode) PLC4XEnumName() string {
	switch e {
	case BACnetLiftGroupMode_UNKNOWN:
		return "UNKNOWN"
	case BACnetLiftGroupMode_NORMAL:
		return "NORMAL"
	case BACnetLiftGroupMode_DOWN_PEAK:
		return "DOWN_PEAK"
	case BACnetLiftGroupMode_TWO_WAY:
		return "TWO_WAY"
	case BACnetLiftGroupMode_FOUR_WAY:
		return "FOUR_WAY"
	case BACnetLiftGroupMode_EMERGENCY_POWER:
		return "EMERGENCY_POWER"
	case BACnetLiftGroupMode_UP_PEAK:
		return "UP_PEAK"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetLiftGroupMode) String() string {
	return e.PLC4XEnumName()
}
