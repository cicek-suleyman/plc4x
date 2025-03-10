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

// COTPProtocolClass is an enum
type COTPProtocolClass uint8

type ICOTPProtocolClass interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	COTPProtocolClass_CLASS_0 COTPProtocolClass = 0x00
	COTPProtocolClass_CLASS_1 COTPProtocolClass = 0x10
	COTPProtocolClass_CLASS_2 COTPProtocolClass = 0x20
	COTPProtocolClass_CLASS_3 COTPProtocolClass = 0x30
	COTPProtocolClass_CLASS_4 COTPProtocolClass = 0x40
)

var COTPProtocolClassValues []COTPProtocolClass

func init() {
	_ = errors.New
	COTPProtocolClassValues = []COTPProtocolClass{
		COTPProtocolClass_CLASS_0,
		COTPProtocolClass_CLASS_1,
		COTPProtocolClass_CLASS_2,
		COTPProtocolClass_CLASS_3,
		COTPProtocolClass_CLASS_4,
	}
}

func COTPProtocolClassByValue(value uint8) (enum COTPProtocolClass, ok bool) {
	switch value {
	case 0x00:
		return COTPProtocolClass_CLASS_0, true
	case 0x10:
		return COTPProtocolClass_CLASS_1, true
	case 0x20:
		return COTPProtocolClass_CLASS_2, true
	case 0x30:
		return COTPProtocolClass_CLASS_3, true
	case 0x40:
		return COTPProtocolClass_CLASS_4, true
	}
	return 0, false
}

func COTPProtocolClassByName(value string) (enum COTPProtocolClass, ok bool) {
	switch value {
	case "CLASS_0":
		return COTPProtocolClass_CLASS_0, true
	case "CLASS_1":
		return COTPProtocolClass_CLASS_1, true
	case "CLASS_2":
		return COTPProtocolClass_CLASS_2, true
	case "CLASS_3":
		return COTPProtocolClass_CLASS_3, true
	case "CLASS_4":
		return COTPProtocolClass_CLASS_4, true
	}
	return 0, false
}

func COTPProtocolClassKnows(value uint8) bool {
	for _, typeValue := range COTPProtocolClassValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCOTPProtocolClass(structType any) COTPProtocolClass {
	castFunc := func(typ any) COTPProtocolClass {
		if sCOTPProtocolClass, ok := typ.(COTPProtocolClass); ok {
			return sCOTPProtocolClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m COTPProtocolClass) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m COTPProtocolClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPProtocolClassParse(ctx context.Context, theBytes []byte) (COTPProtocolClass, error) {
	return COTPProtocolClassParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func COTPProtocolClassParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (COTPProtocolClass, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("COTPProtocolClass", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading COTPProtocolClass")
	}
	if enum, ok := COTPProtocolClassByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for COTPProtocolClass")
		return COTPProtocolClass(val), nil
	} else {
		return enum, nil
	}
}

func (e COTPProtocolClass) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e COTPProtocolClass) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("COTPProtocolClass", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e COTPProtocolClass) PLC4XEnumName() string {
	switch e {
	case COTPProtocolClass_CLASS_0:
		return "CLASS_0"
	case COTPProtocolClass_CLASS_1:
		return "CLASS_1"
	case COTPProtocolClass_CLASS_2:
		return "CLASS_2"
	case COTPProtocolClass_CLASS_3:
		return "CLASS_3"
	case COTPProtocolClass_CLASS_4:
		return "CLASS_4"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e COTPProtocolClass) String() string {
	return e.PLC4XEnumName()
}
