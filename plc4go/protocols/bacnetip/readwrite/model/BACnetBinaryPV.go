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

// BACnetBinaryPV is an enum
type BACnetBinaryPV uint8

type IBACnetBinaryPV interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetBinaryPV_INACTIVE BACnetBinaryPV = 0
	BACnetBinaryPV_ACTIVE   BACnetBinaryPV = 1
)

var BACnetBinaryPVValues []BACnetBinaryPV

func init() {
	_ = errors.New
	BACnetBinaryPVValues = []BACnetBinaryPV{
		BACnetBinaryPV_INACTIVE,
		BACnetBinaryPV_ACTIVE,
	}
}

func BACnetBinaryPVByValue(value uint8) (enum BACnetBinaryPV, ok bool) {
	switch value {
	case 0:
		return BACnetBinaryPV_INACTIVE, true
	case 1:
		return BACnetBinaryPV_ACTIVE, true
	}
	return 0, false
}

func BACnetBinaryPVByName(value string) (enum BACnetBinaryPV, ok bool) {
	switch value {
	case "INACTIVE":
		return BACnetBinaryPV_INACTIVE, true
	case "ACTIVE":
		return BACnetBinaryPV_ACTIVE, true
	}
	return 0, false
}

func BACnetBinaryPVKnows(value uint8) bool {
	for _, typeValue := range BACnetBinaryPVValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetBinaryPV(structType any) BACnetBinaryPV {
	castFunc := func(typ any) BACnetBinaryPV {
		if sBACnetBinaryPV, ok := typ.(BACnetBinaryPV); ok {
			return sBACnetBinaryPV
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetBinaryPV) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetBinaryPV) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetBinaryPVParse(ctx context.Context, theBytes []byte) (BACnetBinaryPV, error) {
	return BACnetBinaryPVParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetBinaryPVParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetBinaryPV, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetBinaryPV", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetBinaryPV")
	}
	if enum, ok := BACnetBinaryPVByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetBinaryPV")
		return BACnetBinaryPV(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetBinaryPV) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetBinaryPV) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetBinaryPV", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetBinaryPV) PLC4XEnumName() string {
	switch e {
	case BACnetBinaryPV_INACTIVE:
		return "INACTIVE"
	case BACnetBinaryPV_ACTIVE:
		return "ACTIVE"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetBinaryPV) String() string {
	return e.PLC4XEnumName()
}
