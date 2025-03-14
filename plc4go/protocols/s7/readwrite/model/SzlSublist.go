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

// SzlSublist is an enum
type SzlSublist uint8

type ISzlSublist interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	SzlSublist_NONE                                                          SzlSublist = 0x00
	SzlSublist_MODULE_IDENTIFICATION                                         SzlSublist = 0x11
	SzlSublist_CPU_FEATURES                                                  SzlSublist = 0x12
	SzlSublist_USER_MEMORY_AREA                                              SzlSublist = 0x13
	SzlSublist_SYSTEM_AREAS                                                  SzlSublist = 0x14
	SzlSublist_BLOCK_TYPES                                                   SzlSublist = 0x15
	SzlSublist_STATUS_MODULE_LEDS                                            SzlSublist = 0x19
	SzlSublist_COMPONENT_IDENTIFICATION                                      SzlSublist = 0x1C
	SzlSublist_INTERRUPT_STATUS                                              SzlSublist = 0x22
	SzlSublist_ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS           SzlSublist = 0x25
	SzlSublist_COMMUNICATION_STATUS_DATA                                     SzlSublist = 0x32
	SzlSublist_H_CPU_GROUP_INFORMATION                                       SzlSublist = 0x71
	SzlSublist_STATUS_SINGLE_MODULE_LED                                      SzlSublist = 0x74
	SzlSublist_SWITCHED_DP_SLAVES_H_SYSTEM                                   SzlSublist = 0x75
	SzlSublist_DP_MASTER_SYSTEM_INFORMATION                                  SzlSublist = 0x90
	SzlSublist_MODULE_STATUS_INFORMATION                                     SzlSublist = 0x91
	SzlSublist_RACK_OR_STATION_STATUS_INFORMATION                            SzlSublist = 0x92
	SzlSublist_RACK_OR_STATION_STATUS_INFORMATION_2                          SzlSublist = 0x94
	SzlSublist_ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION SzlSublist = 0x95
	SzlSublist_MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP         SzlSublist = 0x96
	SzlSublist_TOOL_CHANGER_INFORMATION_PROFINET                             SzlSublist = 0x9C
	SzlSublist_DIAGNOSTIC_BUFFER                                             SzlSublist = 0xA0
	SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR0                             SzlSublist = 0xB1
	SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_GI                          SzlSublist = 0xB2
	SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_LA                          SzlSublist = 0xB3
	SzlSublist_DIAGNOSTIC_DATA_DP_SLAVE                                      SzlSublist = 0xB4
)

var SzlSublistValues []SzlSublist

func init() {
	_ = errors.New
	SzlSublistValues = []SzlSublist{
		SzlSublist_NONE,
		SzlSublist_MODULE_IDENTIFICATION,
		SzlSublist_CPU_FEATURES,
		SzlSublist_USER_MEMORY_AREA,
		SzlSublist_SYSTEM_AREAS,
		SzlSublist_BLOCK_TYPES,
		SzlSublist_STATUS_MODULE_LEDS,
		SzlSublist_COMPONENT_IDENTIFICATION,
		SzlSublist_INTERRUPT_STATUS,
		SzlSublist_ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS,
		SzlSublist_COMMUNICATION_STATUS_DATA,
		SzlSublist_H_CPU_GROUP_INFORMATION,
		SzlSublist_STATUS_SINGLE_MODULE_LED,
		SzlSublist_SWITCHED_DP_SLAVES_H_SYSTEM,
		SzlSublist_DP_MASTER_SYSTEM_INFORMATION,
		SzlSublist_MODULE_STATUS_INFORMATION,
		SzlSublist_RACK_OR_STATION_STATUS_INFORMATION,
		SzlSublist_RACK_OR_STATION_STATUS_INFORMATION_2,
		SzlSublist_ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION,
		SzlSublist_MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP,
		SzlSublist_TOOL_CHANGER_INFORMATION_PROFINET,
		SzlSublist_DIAGNOSTIC_BUFFER,
		SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR0,
		SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_GI,
		SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_LA,
		SzlSublist_DIAGNOSTIC_DATA_DP_SLAVE,
	}
}

func SzlSublistByValue(value uint8) (enum SzlSublist, ok bool) {
	switch value {
	case 0x00:
		return SzlSublist_NONE, true
	case 0x11:
		return SzlSublist_MODULE_IDENTIFICATION, true
	case 0x12:
		return SzlSublist_CPU_FEATURES, true
	case 0x13:
		return SzlSublist_USER_MEMORY_AREA, true
	case 0x14:
		return SzlSublist_SYSTEM_AREAS, true
	case 0x15:
		return SzlSublist_BLOCK_TYPES, true
	case 0x19:
		return SzlSublist_STATUS_MODULE_LEDS, true
	case 0x1C:
		return SzlSublist_COMPONENT_IDENTIFICATION, true
	case 0x22:
		return SzlSublist_INTERRUPT_STATUS, true
	case 0x25:
		return SzlSublist_ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS, true
	case 0x32:
		return SzlSublist_COMMUNICATION_STATUS_DATA, true
	case 0x71:
		return SzlSublist_H_CPU_GROUP_INFORMATION, true
	case 0x74:
		return SzlSublist_STATUS_SINGLE_MODULE_LED, true
	case 0x75:
		return SzlSublist_SWITCHED_DP_SLAVES_H_SYSTEM, true
	case 0x90:
		return SzlSublist_DP_MASTER_SYSTEM_INFORMATION, true
	case 0x91:
		return SzlSublist_MODULE_STATUS_INFORMATION, true
	case 0x92:
		return SzlSublist_RACK_OR_STATION_STATUS_INFORMATION, true
	case 0x94:
		return SzlSublist_RACK_OR_STATION_STATUS_INFORMATION_2, true
	case 0x95:
		return SzlSublist_ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION, true
	case 0x96:
		return SzlSublist_MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP, true
	case 0x9C:
		return SzlSublist_TOOL_CHANGER_INFORMATION_PROFINET, true
	case 0xA0:
		return SzlSublist_DIAGNOSTIC_BUFFER, true
	case 0xB1:
		return SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR0, true
	case 0xB2:
		return SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_GI, true
	case 0xB3:
		return SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_LA, true
	case 0xB4:
		return SzlSublist_DIAGNOSTIC_DATA_DP_SLAVE, true
	}
	return 0, false
}

func SzlSublistByName(value string) (enum SzlSublist, ok bool) {
	switch value {
	case "NONE":
		return SzlSublist_NONE, true
	case "MODULE_IDENTIFICATION":
		return SzlSublist_MODULE_IDENTIFICATION, true
	case "CPU_FEATURES":
		return SzlSublist_CPU_FEATURES, true
	case "USER_MEMORY_AREA":
		return SzlSublist_USER_MEMORY_AREA, true
	case "SYSTEM_AREAS":
		return SzlSublist_SYSTEM_AREAS, true
	case "BLOCK_TYPES":
		return SzlSublist_BLOCK_TYPES, true
	case "STATUS_MODULE_LEDS":
		return SzlSublist_STATUS_MODULE_LEDS, true
	case "COMPONENT_IDENTIFICATION":
		return SzlSublist_COMPONENT_IDENTIFICATION, true
	case "INTERRUPT_STATUS":
		return SzlSublist_INTERRUPT_STATUS, true
	case "ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS":
		return SzlSublist_ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS, true
	case "COMMUNICATION_STATUS_DATA":
		return SzlSublist_COMMUNICATION_STATUS_DATA, true
	case "H_CPU_GROUP_INFORMATION":
		return SzlSublist_H_CPU_GROUP_INFORMATION, true
	case "STATUS_SINGLE_MODULE_LED":
		return SzlSublist_STATUS_SINGLE_MODULE_LED, true
	case "SWITCHED_DP_SLAVES_H_SYSTEM":
		return SzlSublist_SWITCHED_DP_SLAVES_H_SYSTEM, true
	case "DP_MASTER_SYSTEM_INFORMATION":
		return SzlSublist_DP_MASTER_SYSTEM_INFORMATION, true
	case "MODULE_STATUS_INFORMATION":
		return SzlSublist_MODULE_STATUS_INFORMATION, true
	case "RACK_OR_STATION_STATUS_INFORMATION":
		return SzlSublist_RACK_OR_STATION_STATUS_INFORMATION, true
	case "RACK_OR_STATION_STATUS_INFORMATION_2":
		return SzlSublist_RACK_OR_STATION_STATUS_INFORMATION_2, true
	case "ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION":
		return SzlSublist_ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION, true
	case "MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP":
		return SzlSublist_MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP, true
	case "TOOL_CHANGER_INFORMATION_PROFINET":
		return SzlSublist_TOOL_CHANGER_INFORMATION_PROFINET, true
	case "DIAGNOSTIC_BUFFER":
		return SzlSublist_DIAGNOSTIC_BUFFER, true
	case "MODULE_DIAGNOSTIC_INFORMATION_DR0":
		return SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR0, true
	case "MODULE_DIAGNOSTIC_INFORMATION_DR1_GI":
		return SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_GI, true
	case "MODULE_DIAGNOSTIC_INFORMATION_DR1_LA":
		return SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_LA, true
	case "DIAGNOSTIC_DATA_DP_SLAVE":
		return SzlSublist_DIAGNOSTIC_DATA_DP_SLAVE, true
	}
	return 0, false
}

func SzlSublistKnows(value uint8) bool {
	for _, typeValue := range SzlSublistValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastSzlSublist(structType any) SzlSublist {
	castFunc := func(typ any) SzlSublist {
		if sSzlSublist, ok := typ.(SzlSublist); ok {
			return sSzlSublist
		}
		return 0
	}
	return castFunc(structType)
}

func (m SzlSublist) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m SzlSublist) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SzlSublistParse(ctx context.Context, theBytes []byte) (SzlSublist, error) {
	return SzlSublistParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SzlSublistParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SzlSublist, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("SzlSublist", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading SzlSublist")
	}
	if enum, ok := SzlSublistByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for SzlSublist")
		return SzlSublist(val), nil
	} else {
		return enum, nil
	}
}

func (e SzlSublist) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e SzlSublist) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("SzlSublist", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e SzlSublist) PLC4XEnumName() string {
	switch e {
	case SzlSublist_NONE:
		return "NONE"
	case SzlSublist_MODULE_IDENTIFICATION:
		return "MODULE_IDENTIFICATION"
	case SzlSublist_CPU_FEATURES:
		return "CPU_FEATURES"
	case SzlSublist_USER_MEMORY_AREA:
		return "USER_MEMORY_AREA"
	case SzlSublist_SYSTEM_AREAS:
		return "SYSTEM_AREAS"
	case SzlSublist_BLOCK_TYPES:
		return "BLOCK_TYPES"
	case SzlSublist_STATUS_MODULE_LEDS:
		return "STATUS_MODULE_LEDS"
	case SzlSublist_COMPONENT_IDENTIFICATION:
		return "COMPONENT_IDENTIFICATION"
	case SzlSublist_INTERRUPT_STATUS:
		return "INTERRUPT_STATUS"
	case SzlSublist_ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS:
		return "ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS"
	case SzlSublist_COMMUNICATION_STATUS_DATA:
		return "COMMUNICATION_STATUS_DATA"
	case SzlSublist_H_CPU_GROUP_INFORMATION:
		return "H_CPU_GROUP_INFORMATION"
	case SzlSublist_STATUS_SINGLE_MODULE_LED:
		return "STATUS_SINGLE_MODULE_LED"
	case SzlSublist_SWITCHED_DP_SLAVES_H_SYSTEM:
		return "SWITCHED_DP_SLAVES_H_SYSTEM"
	case SzlSublist_DP_MASTER_SYSTEM_INFORMATION:
		return "DP_MASTER_SYSTEM_INFORMATION"
	case SzlSublist_MODULE_STATUS_INFORMATION:
		return "MODULE_STATUS_INFORMATION"
	case SzlSublist_RACK_OR_STATION_STATUS_INFORMATION:
		return "RACK_OR_STATION_STATUS_INFORMATION"
	case SzlSublist_RACK_OR_STATION_STATUS_INFORMATION_2:
		return "RACK_OR_STATION_STATUS_INFORMATION_2"
	case SzlSublist_ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION:
		return "ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION"
	case SzlSublist_MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP:
		return "MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP"
	case SzlSublist_TOOL_CHANGER_INFORMATION_PROFINET:
		return "TOOL_CHANGER_INFORMATION_PROFINET"
	case SzlSublist_DIAGNOSTIC_BUFFER:
		return "DIAGNOSTIC_BUFFER"
	case SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR0:
		return "MODULE_DIAGNOSTIC_INFORMATION_DR0"
	case SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_GI:
		return "MODULE_DIAGNOSTIC_INFORMATION_DR1_GI"
	case SzlSublist_MODULE_DIAGNOSTIC_INFORMATION_DR1_LA:
		return "MODULE_DIAGNOSTIC_INFORMATION_DR1_LA"
	case SzlSublist_DIAGNOSTIC_DATA_DP_SLAVE:
		return "DIAGNOSTIC_DATA_DP_SLAVE"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e SzlSublist) String() string {
	return e.PLC4XEnumName()
}
