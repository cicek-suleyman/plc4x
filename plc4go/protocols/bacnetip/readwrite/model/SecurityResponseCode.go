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

// SecurityResponseCode is an enum
type SecurityResponseCode uint8

type ISecurityResponseCode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	SecurityResponseCode_SUCCESS                        SecurityResponseCode = 0x00
	SecurityResponseCode_ACCESS_DENIED                  SecurityResponseCode = 0x01
	SecurityResponseCode_BAD_DESTINATION_ADDRESS        SecurityResponseCode = 0x02
	SecurityResponseCode_BAD_DESTINATION_DEVICE_ID      SecurityResponseCode = 0x03
	SecurityResponseCode_BAD_SIGNATURE                  SecurityResponseCode = 0x04
	SecurityResponseCode_BAD_SOURCE_ADDRESS             SecurityResponseCode = 0x05
	SecurityResponseCode_BAD_TIMESTAMP                  SecurityResponseCode = 0x06
	SecurityResponseCode_CANNOT_USE_KEY                 SecurityResponseCode = 0x07
	SecurityResponseCode_CANNOT_VERIFY_MESSAGE_ID       SecurityResponseCode = 0x08
	SecurityResponseCode_CORRECT_KEY_REVISION           SecurityResponseCode = 0x09
	SecurityResponseCode_DESTINATION_DEVICE_ID_REQUIRED SecurityResponseCode = 0x0A
	SecurityResponseCode_DUPLICATE_MESSAGE              SecurityResponseCode = 0x0B
	SecurityResponseCode_ENCRYPTION_NOT_CONFIGURED      SecurityResponseCode = 0x0C
	SecurityResponseCode_ENCRYPTION_REQUIRED            SecurityResponseCode = 0x0D
	SecurityResponseCode_INCORRECT_KEY                  SecurityResponseCode = 0x0E
	SecurityResponseCode_INVALID_KEY_DATA               SecurityResponseCode = 0x0F
	SecurityResponseCode_KEY_UPDATE_IN_PROGRESS         SecurityResponseCode = 0x10
	SecurityResponseCode_MALFORMED_MESSAGE              SecurityResponseCode = 0x11
	SecurityResponseCode_NOT_KEY_SERVER                 SecurityResponseCode = 0x12
	SecurityResponseCode_SECURITY_NOT_CONFIGURED        SecurityResponseCode = 0x13
	SecurityResponseCode_SOURCE_SECURITY_REQUIRED       SecurityResponseCode = 0x14
	SecurityResponseCode_TOO_MANY_KEYS                  SecurityResponseCode = 0x15
	SecurityResponseCode_UNKNOWN_AUTHENTICATION_TYPE    SecurityResponseCode = 0x16
	SecurityResponseCode_UNKNOWN_KEY                    SecurityResponseCode = 0x17
	SecurityResponseCode_UNKNOWN_KEY_REVISION           SecurityResponseCode = 0x18
	SecurityResponseCode_UNKNOWN_SOURCE_MESSAGE         SecurityResponseCode = 0x19
)

var SecurityResponseCodeValues []SecurityResponseCode

func init() {
	_ = errors.New
	SecurityResponseCodeValues = []SecurityResponseCode{
		SecurityResponseCode_SUCCESS,
		SecurityResponseCode_ACCESS_DENIED,
		SecurityResponseCode_BAD_DESTINATION_ADDRESS,
		SecurityResponseCode_BAD_DESTINATION_DEVICE_ID,
		SecurityResponseCode_BAD_SIGNATURE,
		SecurityResponseCode_BAD_SOURCE_ADDRESS,
		SecurityResponseCode_BAD_TIMESTAMP,
		SecurityResponseCode_CANNOT_USE_KEY,
		SecurityResponseCode_CANNOT_VERIFY_MESSAGE_ID,
		SecurityResponseCode_CORRECT_KEY_REVISION,
		SecurityResponseCode_DESTINATION_DEVICE_ID_REQUIRED,
		SecurityResponseCode_DUPLICATE_MESSAGE,
		SecurityResponseCode_ENCRYPTION_NOT_CONFIGURED,
		SecurityResponseCode_ENCRYPTION_REQUIRED,
		SecurityResponseCode_INCORRECT_KEY,
		SecurityResponseCode_INVALID_KEY_DATA,
		SecurityResponseCode_KEY_UPDATE_IN_PROGRESS,
		SecurityResponseCode_MALFORMED_MESSAGE,
		SecurityResponseCode_NOT_KEY_SERVER,
		SecurityResponseCode_SECURITY_NOT_CONFIGURED,
		SecurityResponseCode_SOURCE_SECURITY_REQUIRED,
		SecurityResponseCode_TOO_MANY_KEYS,
		SecurityResponseCode_UNKNOWN_AUTHENTICATION_TYPE,
		SecurityResponseCode_UNKNOWN_KEY,
		SecurityResponseCode_UNKNOWN_KEY_REVISION,
		SecurityResponseCode_UNKNOWN_SOURCE_MESSAGE,
	}
}

func SecurityResponseCodeByValue(value uint8) (enum SecurityResponseCode, ok bool) {
	switch value {
	case 0x00:
		return SecurityResponseCode_SUCCESS, true
	case 0x01:
		return SecurityResponseCode_ACCESS_DENIED, true
	case 0x02:
		return SecurityResponseCode_BAD_DESTINATION_ADDRESS, true
	case 0x03:
		return SecurityResponseCode_BAD_DESTINATION_DEVICE_ID, true
	case 0x04:
		return SecurityResponseCode_BAD_SIGNATURE, true
	case 0x05:
		return SecurityResponseCode_BAD_SOURCE_ADDRESS, true
	case 0x06:
		return SecurityResponseCode_BAD_TIMESTAMP, true
	case 0x07:
		return SecurityResponseCode_CANNOT_USE_KEY, true
	case 0x08:
		return SecurityResponseCode_CANNOT_VERIFY_MESSAGE_ID, true
	case 0x09:
		return SecurityResponseCode_CORRECT_KEY_REVISION, true
	case 0x0A:
		return SecurityResponseCode_DESTINATION_DEVICE_ID_REQUIRED, true
	case 0x0B:
		return SecurityResponseCode_DUPLICATE_MESSAGE, true
	case 0x0C:
		return SecurityResponseCode_ENCRYPTION_NOT_CONFIGURED, true
	case 0x0D:
		return SecurityResponseCode_ENCRYPTION_REQUIRED, true
	case 0x0E:
		return SecurityResponseCode_INCORRECT_KEY, true
	case 0x0F:
		return SecurityResponseCode_INVALID_KEY_DATA, true
	case 0x10:
		return SecurityResponseCode_KEY_UPDATE_IN_PROGRESS, true
	case 0x11:
		return SecurityResponseCode_MALFORMED_MESSAGE, true
	case 0x12:
		return SecurityResponseCode_NOT_KEY_SERVER, true
	case 0x13:
		return SecurityResponseCode_SECURITY_NOT_CONFIGURED, true
	case 0x14:
		return SecurityResponseCode_SOURCE_SECURITY_REQUIRED, true
	case 0x15:
		return SecurityResponseCode_TOO_MANY_KEYS, true
	case 0x16:
		return SecurityResponseCode_UNKNOWN_AUTHENTICATION_TYPE, true
	case 0x17:
		return SecurityResponseCode_UNKNOWN_KEY, true
	case 0x18:
		return SecurityResponseCode_UNKNOWN_KEY_REVISION, true
	case 0x19:
		return SecurityResponseCode_UNKNOWN_SOURCE_MESSAGE, true
	}
	return 0, false
}

func SecurityResponseCodeByName(value string) (enum SecurityResponseCode, ok bool) {
	switch value {
	case "SUCCESS":
		return SecurityResponseCode_SUCCESS, true
	case "ACCESS_DENIED":
		return SecurityResponseCode_ACCESS_DENIED, true
	case "BAD_DESTINATION_ADDRESS":
		return SecurityResponseCode_BAD_DESTINATION_ADDRESS, true
	case "BAD_DESTINATION_DEVICE_ID":
		return SecurityResponseCode_BAD_DESTINATION_DEVICE_ID, true
	case "BAD_SIGNATURE":
		return SecurityResponseCode_BAD_SIGNATURE, true
	case "BAD_SOURCE_ADDRESS":
		return SecurityResponseCode_BAD_SOURCE_ADDRESS, true
	case "BAD_TIMESTAMP":
		return SecurityResponseCode_BAD_TIMESTAMP, true
	case "CANNOT_USE_KEY":
		return SecurityResponseCode_CANNOT_USE_KEY, true
	case "CANNOT_VERIFY_MESSAGE_ID":
		return SecurityResponseCode_CANNOT_VERIFY_MESSAGE_ID, true
	case "CORRECT_KEY_REVISION":
		return SecurityResponseCode_CORRECT_KEY_REVISION, true
	case "DESTINATION_DEVICE_ID_REQUIRED":
		return SecurityResponseCode_DESTINATION_DEVICE_ID_REQUIRED, true
	case "DUPLICATE_MESSAGE":
		return SecurityResponseCode_DUPLICATE_MESSAGE, true
	case "ENCRYPTION_NOT_CONFIGURED":
		return SecurityResponseCode_ENCRYPTION_NOT_CONFIGURED, true
	case "ENCRYPTION_REQUIRED":
		return SecurityResponseCode_ENCRYPTION_REQUIRED, true
	case "INCORRECT_KEY":
		return SecurityResponseCode_INCORRECT_KEY, true
	case "INVALID_KEY_DATA":
		return SecurityResponseCode_INVALID_KEY_DATA, true
	case "KEY_UPDATE_IN_PROGRESS":
		return SecurityResponseCode_KEY_UPDATE_IN_PROGRESS, true
	case "MALFORMED_MESSAGE":
		return SecurityResponseCode_MALFORMED_MESSAGE, true
	case "NOT_KEY_SERVER":
		return SecurityResponseCode_NOT_KEY_SERVER, true
	case "SECURITY_NOT_CONFIGURED":
		return SecurityResponseCode_SECURITY_NOT_CONFIGURED, true
	case "SOURCE_SECURITY_REQUIRED":
		return SecurityResponseCode_SOURCE_SECURITY_REQUIRED, true
	case "TOO_MANY_KEYS":
		return SecurityResponseCode_TOO_MANY_KEYS, true
	case "UNKNOWN_AUTHENTICATION_TYPE":
		return SecurityResponseCode_UNKNOWN_AUTHENTICATION_TYPE, true
	case "UNKNOWN_KEY":
		return SecurityResponseCode_UNKNOWN_KEY, true
	case "UNKNOWN_KEY_REVISION":
		return SecurityResponseCode_UNKNOWN_KEY_REVISION, true
	case "UNKNOWN_SOURCE_MESSAGE":
		return SecurityResponseCode_UNKNOWN_SOURCE_MESSAGE, true
	}
	return 0, false
}

func SecurityResponseCodeKnows(value uint8) bool {
	for _, typeValue := range SecurityResponseCodeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastSecurityResponseCode(structType any) SecurityResponseCode {
	castFunc := func(typ any) SecurityResponseCode {
		if sSecurityResponseCode, ok := typ.(SecurityResponseCode); ok {
			return sSecurityResponseCode
		}
		return 0
	}
	return castFunc(structType)
}

func (m SecurityResponseCode) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m SecurityResponseCode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityResponseCodeParse(ctx context.Context, theBytes []byte) (SecurityResponseCode, error) {
	return SecurityResponseCodeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityResponseCodeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityResponseCode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("SecurityResponseCode", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading SecurityResponseCode")
	}
	if enum, ok := SecurityResponseCodeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for SecurityResponseCode")
		return SecurityResponseCode(val), nil
	} else {
		return enum, nil
	}
}

func (e SecurityResponseCode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e SecurityResponseCode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("SecurityResponseCode", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e SecurityResponseCode) PLC4XEnumName() string {
	switch e {
	case SecurityResponseCode_SUCCESS:
		return "SUCCESS"
	case SecurityResponseCode_ACCESS_DENIED:
		return "ACCESS_DENIED"
	case SecurityResponseCode_BAD_DESTINATION_ADDRESS:
		return "BAD_DESTINATION_ADDRESS"
	case SecurityResponseCode_BAD_DESTINATION_DEVICE_ID:
		return "BAD_DESTINATION_DEVICE_ID"
	case SecurityResponseCode_BAD_SIGNATURE:
		return "BAD_SIGNATURE"
	case SecurityResponseCode_BAD_SOURCE_ADDRESS:
		return "BAD_SOURCE_ADDRESS"
	case SecurityResponseCode_BAD_TIMESTAMP:
		return "BAD_TIMESTAMP"
	case SecurityResponseCode_CANNOT_USE_KEY:
		return "CANNOT_USE_KEY"
	case SecurityResponseCode_CANNOT_VERIFY_MESSAGE_ID:
		return "CANNOT_VERIFY_MESSAGE_ID"
	case SecurityResponseCode_CORRECT_KEY_REVISION:
		return "CORRECT_KEY_REVISION"
	case SecurityResponseCode_DESTINATION_DEVICE_ID_REQUIRED:
		return "DESTINATION_DEVICE_ID_REQUIRED"
	case SecurityResponseCode_DUPLICATE_MESSAGE:
		return "DUPLICATE_MESSAGE"
	case SecurityResponseCode_ENCRYPTION_NOT_CONFIGURED:
		return "ENCRYPTION_NOT_CONFIGURED"
	case SecurityResponseCode_ENCRYPTION_REQUIRED:
		return "ENCRYPTION_REQUIRED"
	case SecurityResponseCode_INCORRECT_KEY:
		return "INCORRECT_KEY"
	case SecurityResponseCode_INVALID_KEY_DATA:
		return "INVALID_KEY_DATA"
	case SecurityResponseCode_KEY_UPDATE_IN_PROGRESS:
		return "KEY_UPDATE_IN_PROGRESS"
	case SecurityResponseCode_MALFORMED_MESSAGE:
		return "MALFORMED_MESSAGE"
	case SecurityResponseCode_NOT_KEY_SERVER:
		return "NOT_KEY_SERVER"
	case SecurityResponseCode_SECURITY_NOT_CONFIGURED:
		return "SECURITY_NOT_CONFIGURED"
	case SecurityResponseCode_SOURCE_SECURITY_REQUIRED:
		return "SOURCE_SECURITY_REQUIRED"
	case SecurityResponseCode_TOO_MANY_KEYS:
		return "TOO_MANY_KEYS"
	case SecurityResponseCode_UNKNOWN_AUTHENTICATION_TYPE:
		return "UNKNOWN_AUTHENTICATION_TYPE"
	case SecurityResponseCode_UNKNOWN_KEY:
		return "UNKNOWN_KEY"
	case SecurityResponseCode_UNKNOWN_KEY_REVISION:
		return "UNKNOWN_KEY_REVISION"
	case SecurityResponseCode_UNKNOWN_SOURCE_MESSAGE:
		return "UNKNOWN_SOURCE_MESSAGE"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e SecurityResponseCode) String() string {
	return e.PLC4XEnumName()
}
