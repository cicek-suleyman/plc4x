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

// OpcuaNodeIdServicesVariableAnalog is an enum
type OpcuaNodeIdServicesVariableAnalog int32

type IOpcuaNodeIdServicesVariableAnalog interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EURange               OpcuaNodeIdServicesVariableAnalog = 2369
	OpcuaNodeIdServicesVariableAnalog_AnalogItemType_InstrumentRange       OpcuaNodeIdServicesVariableAnalog = 2370
	OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EngineeringUnits      OpcuaNodeIdServicesVariableAnalog = 2371
	OpcuaNodeIdServicesVariableAnalog_AnalogItemType_Definition            OpcuaNodeIdServicesVariableAnalog = 3774
	OpcuaNodeIdServicesVariableAnalog_AnalogItemType_ValuePrecision        OpcuaNodeIdServicesVariableAnalog = 3775
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_Definition            OpcuaNodeIdServicesVariableAnalog = 17498
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_ValuePrecision        OpcuaNodeIdServicesVariableAnalog = 17499
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_InstrumentRange       OpcuaNodeIdServicesVariableAnalog = 17500
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EURange               OpcuaNodeIdServicesVariableAnalog = 17501
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EngineeringUnits      OpcuaNodeIdServicesVariableAnalog = 17502
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_Definition       OpcuaNodeIdServicesVariableAnalog = 17571
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_ValuePrecision   OpcuaNodeIdServicesVariableAnalog = 17572
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_InstrumentRange  OpcuaNodeIdServicesVariableAnalog = 17573
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EURange          OpcuaNodeIdServicesVariableAnalog = 17574
	OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EngineeringUnits OpcuaNodeIdServicesVariableAnalog = 17575
)

var OpcuaNodeIdServicesVariableAnalogValues []OpcuaNodeIdServicesVariableAnalog

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableAnalogValues = []OpcuaNodeIdServicesVariableAnalog{
		OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EURange,
		OpcuaNodeIdServicesVariableAnalog_AnalogItemType_InstrumentRange,
		OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EngineeringUnits,
		OpcuaNodeIdServicesVariableAnalog_AnalogItemType_Definition,
		OpcuaNodeIdServicesVariableAnalog_AnalogItemType_ValuePrecision,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_Definition,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_ValuePrecision,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_InstrumentRange,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EURange,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EngineeringUnits,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_Definition,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_ValuePrecision,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_InstrumentRange,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EURange,
		OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EngineeringUnits,
	}
}

func OpcuaNodeIdServicesVariableAnalogByValue(value int32) (enum OpcuaNodeIdServicesVariableAnalog, ok bool) {
	switch value {
	case 17498:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_Definition, true
	case 17499:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_ValuePrecision, true
	case 17500:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_InstrumentRange, true
	case 17501:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EURange, true
	case 17502:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EngineeringUnits, true
	case 17571:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_Definition, true
	case 17572:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_ValuePrecision, true
	case 17573:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_InstrumentRange, true
	case 17574:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EURange, true
	case 17575:
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EngineeringUnits, true
	case 2369:
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EURange, true
	case 2370:
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_InstrumentRange, true
	case 2371:
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EngineeringUnits, true
	case 3774:
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_Definition, true
	case 3775:
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_ValuePrecision, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAnalogByName(value string) (enum OpcuaNodeIdServicesVariableAnalog, ok bool) {
	switch value {
	case "AnalogUnitType_Definition":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_Definition, true
	case "AnalogUnitType_ValuePrecision":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_ValuePrecision, true
	case "AnalogUnitType_InstrumentRange":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_InstrumentRange, true
	case "AnalogUnitType_EURange":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EURange, true
	case "AnalogUnitType_EngineeringUnits":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EngineeringUnits, true
	case "AnalogUnitRangeType_Definition":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_Definition, true
	case "AnalogUnitRangeType_ValuePrecision":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_ValuePrecision, true
	case "AnalogUnitRangeType_InstrumentRange":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_InstrumentRange, true
	case "AnalogUnitRangeType_EURange":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EURange, true
	case "AnalogUnitRangeType_EngineeringUnits":
		return OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EngineeringUnits, true
	case "AnalogItemType_EURange":
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EURange, true
	case "AnalogItemType_InstrumentRange":
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_InstrumentRange, true
	case "AnalogItemType_EngineeringUnits":
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EngineeringUnits, true
	case "AnalogItemType_Definition":
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_Definition, true
	case "AnalogItemType_ValuePrecision":
		return OpcuaNodeIdServicesVariableAnalog_AnalogItemType_ValuePrecision, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAnalogKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableAnalogValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableAnalog(structType any) OpcuaNodeIdServicesVariableAnalog {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableAnalog {
		if sOpcuaNodeIdServicesVariableAnalog, ok := typ.(OpcuaNodeIdServicesVariableAnalog); ok {
			return sOpcuaNodeIdServicesVariableAnalog
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableAnalog) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableAnalog) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableAnalogParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableAnalog, error) {
	return OpcuaNodeIdServicesVariableAnalogParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableAnalogParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableAnalog, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableAnalog", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableAnalog")
	}
	if enum, ok := OpcuaNodeIdServicesVariableAnalogByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableAnalog")
		return OpcuaNodeIdServicesVariableAnalog(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableAnalog) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableAnalog) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableAnalog", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableAnalog) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_Definition:
		return "AnalogUnitType_Definition"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_ValuePrecision:
		return "AnalogUnitType_ValuePrecision"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_InstrumentRange:
		return "AnalogUnitType_InstrumentRange"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EURange:
		return "AnalogUnitType_EURange"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitType_EngineeringUnits:
		return "AnalogUnitType_EngineeringUnits"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_Definition:
		return "AnalogUnitRangeType_Definition"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_ValuePrecision:
		return "AnalogUnitRangeType_ValuePrecision"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_InstrumentRange:
		return "AnalogUnitRangeType_InstrumentRange"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EURange:
		return "AnalogUnitRangeType_EURange"
	case OpcuaNodeIdServicesVariableAnalog_AnalogUnitRangeType_EngineeringUnits:
		return "AnalogUnitRangeType_EngineeringUnits"
	case OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EURange:
		return "AnalogItemType_EURange"
	case OpcuaNodeIdServicesVariableAnalog_AnalogItemType_InstrumentRange:
		return "AnalogItemType_InstrumentRange"
	case OpcuaNodeIdServicesVariableAnalog_AnalogItemType_EngineeringUnits:
		return "AnalogItemType_EngineeringUnits"
	case OpcuaNodeIdServicesVariableAnalog_AnalogItemType_Definition:
		return "AnalogItemType_Definition"
	case OpcuaNodeIdServicesVariableAnalog_AnalogItemType_ValuePrecision:
		return "AnalogItemType_ValuePrecision"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableAnalog) String() string {
	return e.PLC4XEnumName()
}
