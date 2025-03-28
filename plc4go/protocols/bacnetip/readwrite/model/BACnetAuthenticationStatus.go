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

// BACnetAuthenticationStatus is an enum
type BACnetAuthenticationStatus uint8

type IBACnetAuthenticationStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetAuthenticationStatus_NOT_READY                         BACnetAuthenticationStatus = 0
	BACnetAuthenticationStatus_READY                             BACnetAuthenticationStatus = 1
	BACnetAuthenticationStatus_DISABLED                          BACnetAuthenticationStatus = 2
	BACnetAuthenticationStatus_WAITING_FOR_AUTHENTICATION_FACTOR BACnetAuthenticationStatus = 3
	BACnetAuthenticationStatus_WAITING_FOR_ACCOMPANIMENT         BACnetAuthenticationStatus = 4
	BACnetAuthenticationStatus_WAITING_FOR_VERIFICATION          BACnetAuthenticationStatus = 5
	BACnetAuthenticationStatus_IN_PROGRESS                       BACnetAuthenticationStatus = 6
)

var BACnetAuthenticationStatusValues []BACnetAuthenticationStatus

func init() {
	_ = errors.New
	BACnetAuthenticationStatusValues = []BACnetAuthenticationStatus{
		BACnetAuthenticationStatus_NOT_READY,
		BACnetAuthenticationStatus_READY,
		BACnetAuthenticationStatus_DISABLED,
		BACnetAuthenticationStatus_WAITING_FOR_AUTHENTICATION_FACTOR,
		BACnetAuthenticationStatus_WAITING_FOR_ACCOMPANIMENT,
		BACnetAuthenticationStatus_WAITING_FOR_VERIFICATION,
		BACnetAuthenticationStatus_IN_PROGRESS,
	}
}

func BACnetAuthenticationStatusByValue(value uint8) (enum BACnetAuthenticationStatus, ok bool) {
	switch value {
	case 0:
		return BACnetAuthenticationStatus_NOT_READY, true
	case 1:
		return BACnetAuthenticationStatus_READY, true
	case 2:
		return BACnetAuthenticationStatus_DISABLED, true
	case 3:
		return BACnetAuthenticationStatus_WAITING_FOR_AUTHENTICATION_FACTOR, true
	case 4:
		return BACnetAuthenticationStatus_WAITING_FOR_ACCOMPANIMENT, true
	case 5:
		return BACnetAuthenticationStatus_WAITING_FOR_VERIFICATION, true
	case 6:
		return BACnetAuthenticationStatus_IN_PROGRESS, true
	}
	return 0, false
}

func BACnetAuthenticationStatusByName(value string) (enum BACnetAuthenticationStatus, ok bool) {
	switch value {
	case "NOT_READY":
		return BACnetAuthenticationStatus_NOT_READY, true
	case "READY":
		return BACnetAuthenticationStatus_READY, true
	case "DISABLED":
		return BACnetAuthenticationStatus_DISABLED, true
	case "WAITING_FOR_AUTHENTICATION_FACTOR":
		return BACnetAuthenticationStatus_WAITING_FOR_AUTHENTICATION_FACTOR, true
	case "WAITING_FOR_ACCOMPANIMENT":
		return BACnetAuthenticationStatus_WAITING_FOR_ACCOMPANIMENT, true
	case "WAITING_FOR_VERIFICATION":
		return BACnetAuthenticationStatus_WAITING_FOR_VERIFICATION, true
	case "IN_PROGRESS":
		return BACnetAuthenticationStatus_IN_PROGRESS, true
	}
	return 0, false
}

func BACnetAuthenticationStatusKnows(value uint8) bool {
	for _, typeValue := range BACnetAuthenticationStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAuthenticationStatus(structType any) BACnetAuthenticationStatus {
	castFunc := func(typ any) BACnetAuthenticationStatus {
		if sBACnetAuthenticationStatus, ok := typ.(BACnetAuthenticationStatus); ok {
			return sBACnetAuthenticationStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAuthenticationStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetAuthenticationStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationStatusParse(ctx context.Context, theBytes []byte) (BACnetAuthenticationStatus, error) {
	return BACnetAuthenticationStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthenticationStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetAuthenticationStatus", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAuthenticationStatus")
	}
	if enum, ok := BACnetAuthenticationStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetAuthenticationStatus")
		return BACnetAuthenticationStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAuthenticationStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAuthenticationStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetAuthenticationStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAuthenticationStatus) PLC4XEnumName() string {
	switch e {
	case BACnetAuthenticationStatus_NOT_READY:
		return "NOT_READY"
	case BACnetAuthenticationStatus_READY:
		return "READY"
	case BACnetAuthenticationStatus_DISABLED:
		return "DISABLED"
	case BACnetAuthenticationStatus_WAITING_FOR_AUTHENTICATION_FACTOR:
		return "WAITING_FOR_AUTHENTICATION_FACTOR"
	case BACnetAuthenticationStatus_WAITING_FOR_ACCOMPANIMENT:
		return "WAITING_FOR_ACCOMPANIMENT"
	case BACnetAuthenticationStatus_WAITING_FOR_VERIFICATION:
		return "WAITING_FOR_VERIFICATION"
	case BACnetAuthenticationStatus_IN_PROGRESS:
		return "IN_PROGRESS"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetAuthenticationStatus) String() string {
	return e.PLC4XEnumName()
}
