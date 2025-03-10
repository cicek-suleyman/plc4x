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

// ZoneStatusTemp is an enum
type ZoneStatusTemp uint8

type IZoneStatusTemp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ZoneStatusTemp_ZONE_SEALED   ZoneStatusTemp = 0x0
	ZoneStatusTemp_ZONE_UNSEALED ZoneStatusTemp = 0x1
	ZoneStatusTemp_ZONE_OPEN     ZoneStatusTemp = 0x2
	ZoneStatusTemp_ZONE_SHORT    ZoneStatusTemp = 0x3
)

var ZoneStatusTempValues []ZoneStatusTemp

func init() {
	_ = errors.New
	ZoneStatusTempValues = []ZoneStatusTemp{
		ZoneStatusTemp_ZONE_SEALED,
		ZoneStatusTemp_ZONE_UNSEALED,
		ZoneStatusTemp_ZONE_OPEN,
		ZoneStatusTemp_ZONE_SHORT,
	}
}

func ZoneStatusTempByValue(value uint8) (enum ZoneStatusTemp, ok bool) {
	switch value {
	case 0x0:
		return ZoneStatusTemp_ZONE_SEALED, true
	case 0x1:
		return ZoneStatusTemp_ZONE_UNSEALED, true
	case 0x2:
		return ZoneStatusTemp_ZONE_OPEN, true
	case 0x3:
		return ZoneStatusTemp_ZONE_SHORT, true
	}
	return 0, false
}

func ZoneStatusTempByName(value string) (enum ZoneStatusTemp, ok bool) {
	switch value {
	case "ZONE_SEALED":
		return ZoneStatusTemp_ZONE_SEALED, true
	case "ZONE_UNSEALED":
		return ZoneStatusTemp_ZONE_UNSEALED, true
	case "ZONE_OPEN":
		return ZoneStatusTemp_ZONE_OPEN, true
	case "ZONE_SHORT":
		return ZoneStatusTemp_ZONE_SHORT, true
	}
	return 0, false
}

func ZoneStatusTempKnows(value uint8) bool {
	for _, typeValue := range ZoneStatusTempValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastZoneStatusTemp(structType any) ZoneStatusTemp {
	castFunc := func(typ any) ZoneStatusTemp {
		if sZoneStatusTemp, ok := typ.(ZoneStatusTemp); ok {
			return sZoneStatusTemp
		}
		return 0
	}
	return castFunc(structType)
}

func (m ZoneStatusTemp) GetLengthInBits(ctx context.Context) uint16 {
	return 2
}

func (m ZoneStatusTemp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ZoneStatusTempParse(ctx context.Context, theBytes []byte) (ZoneStatusTemp, error) {
	return ZoneStatusTempParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ZoneStatusTempParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ZoneStatusTemp, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("ZoneStatusTemp", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ZoneStatusTemp")
	}
	if enum, ok := ZoneStatusTempByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ZoneStatusTemp")
		return ZoneStatusTemp(val), nil
	} else {
		return enum, nil
	}
}

func (e ZoneStatusTemp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ZoneStatusTemp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("ZoneStatusTemp", 2, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ZoneStatusTemp) PLC4XEnumName() string {
	switch e {
	case ZoneStatusTemp_ZONE_SEALED:
		return "ZONE_SEALED"
	case ZoneStatusTemp_ZONE_UNSEALED:
		return "ZONE_UNSEALED"
	case ZoneStatusTemp_ZONE_OPEN:
		return "ZONE_OPEN"
	case ZoneStatusTemp_ZONE_SHORT:
		return "ZONE_SHORT"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e ZoneStatusTemp) String() string {
	return e.PLC4XEnumName()
}
