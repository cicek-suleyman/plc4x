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

// BACnetMaintenance is an enum
type BACnetMaintenance uint8

type IBACnetMaintenance interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetMaintenance_NONE                     BACnetMaintenance = 0
	BACnetMaintenance_PERIODIC_TEST            BACnetMaintenance = 1
	BACnetMaintenance_NEED_SERVICE_OPERATIONAL BACnetMaintenance = 2
	BACnetMaintenance_NEED_SERVICE_INOPERATIVE BACnetMaintenance = 3
	BACnetMaintenance_VENDOR_PROPRIETARY_VALUE BACnetMaintenance = 0xFF
)

var BACnetMaintenanceValues []BACnetMaintenance

func init() {
	_ = errors.New
	BACnetMaintenanceValues = []BACnetMaintenance{
		BACnetMaintenance_NONE,
		BACnetMaintenance_PERIODIC_TEST,
		BACnetMaintenance_NEED_SERVICE_OPERATIONAL,
		BACnetMaintenance_NEED_SERVICE_INOPERATIVE,
		BACnetMaintenance_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetMaintenanceByValue(value uint8) (enum BACnetMaintenance, ok bool) {
	switch value {
	case 0:
		return BACnetMaintenance_NONE, true
	case 0xFF:
		return BACnetMaintenance_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetMaintenance_PERIODIC_TEST, true
	case 2:
		return BACnetMaintenance_NEED_SERVICE_OPERATIONAL, true
	case 3:
		return BACnetMaintenance_NEED_SERVICE_INOPERATIVE, true
	}
	return 0, false
}

func BACnetMaintenanceByName(value string) (enum BACnetMaintenance, ok bool) {
	switch value {
	case "NONE":
		return BACnetMaintenance_NONE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetMaintenance_VENDOR_PROPRIETARY_VALUE, true
	case "PERIODIC_TEST":
		return BACnetMaintenance_PERIODIC_TEST, true
	case "NEED_SERVICE_OPERATIONAL":
		return BACnetMaintenance_NEED_SERVICE_OPERATIONAL, true
	case "NEED_SERVICE_INOPERATIVE":
		return BACnetMaintenance_NEED_SERVICE_INOPERATIVE, true
	}
	return 0, false
}

func BACnetMaintenanceKnows(value uint8) bool {
	for _, typeValue := range BACnetMaintenanceValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetMaintenance(structType any) BACnetMaintenance {
	castFunc := func(typ any) BACnetMaintenance {
		if sBACnetMaintenance, ok := typ.(BACnetMaintenance); ok {
			return sBACnetMaintenance
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetMaintenance) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetMaintenance) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetMaintenanceParse(ctx context.Context, theBytes []byte) (BACnetMaintenance, error) {
	return BACnetMaintenanceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetMaintenanceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetMaintenance, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetMaintenance", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetMaintenance")
	}
	if enum, ok := BACnetMaintenanceByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetMaintenance")
		return BACnetMaintenance(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetMaintenance) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetMaintenance) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetMaintenance", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetMaintenance) PLC4XEnumName() string {
	switch e {
	case BACnetMaintenance_NONE:
		return "NONE"
	case BACnetMaintenance_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetMaintenance_PERIODIC_TEST:
		return "PERIODIC_TEST"
	case BACnetMaintenance_NEED_SERVICE_OPERATIONAL:
		return "NEED_SERVICE_OPERATIONAL"
	case BACnetMaintenance_NEED_SERVICE_INOPERATIVE:
		return "NEED_SERVICE_INOPERATIVE"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetMaintenance) String() string {
	return e.PLC4XEnumName()
}
