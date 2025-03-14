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

// DefaultAmsPorts is an enum
type DefaultAmsPorts uint16

type IDefaultAmsPorts interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DefaultAmsPorts_CAM_CONTROLLER      DefaultAmsPorts = 900
	DefaultAmsPorts_RUNTIME_SYSTEM_01   DefaultAmsPorts = 851
	DefaultAmsPorts_RUNTIME_SYSTEM_02   DefaultAmsPorts = 852
	DefaultAmsPorts_RUNTIME_SYSTEM_03   DefaultAmsPorts = 853
	DefaultAmsPorts_RUNTIME_SYSTEM_04   DefaultAmsPorts = 854
	DefaultAmsPorts_RUNTIME_SYSTEM_05   DefaultAmsPorts = 855
	DefaultAmsPorts_RUNTIME_SYSTEM_06   DefaultAmsPorts = 856
	DefaultAmsPorts_RUNTIME_SYSTEM_07   DefaultAmsPorts = 857
	DefaultAmsPorts_RUNTIME_SYSTEM_08   DefaultAmsPorts = 858
	DefaultAmsPorts_RUNTIME_SYSTEM_09   DefaultAmsPorts = 859
	DefaultAmsPorts_RUNTIME_SYSTEM_10   DefaultAmsPorts = 860
	DefaultAmsPorts_RUNTIME_SYSTEM_11   DefaultAmsPorts = 861
	DefaultAmsPorts_RUNTIME_SYSTEM_12   DefaultAmsPorts = 862
	DefaultAmsPorts_RUNTIME_SYSTEM_13   DefaultAmsPorts = 863
	DefaultAmsPorts_RUNTIME_SYSTEM_14   DefaultAmsPorts = 864
	DefaultAmsPorts_RUNTIME_SYSTEM_15   DefaultAmsPorts = 865
	DefaultAmsPorts_RUNTIME_SYSTEM_16   DefaultAmsPorts = 866
	DefaultAmsPorts_RUNTIME_SYSTEM_17   DefaultAmsPorts = 867
	DefaultAmsPorts_RUNTIME_SYSTEM_18   DefaultAmsPorts = 868
	DefaultAmsPorts_RUNTIME_SYSTEM_19   DefaultAmsPorts = 869
	DefaultAmsPorts_RUNTIME_SYSTEM_20   DefaultAmsPorts = 870
	DefaultAmsPorts_RUNTIME_SYSTEM_21   DefaultAmsPorts = 871
	DefaultAmsPorts_RUNTIME_SYSTEM_22   DefaultAmsPorts = 872
	DefaultAmsPorts_RUNTIME_SYSTEM_23   DefaultAmsPorts = 873
	DefaultAmsPorts_RUNTIME_SYSTEM_24   DefaultAmsPorts = 874
	DefaultAmsPorts_RUNTIME_SYSTEM_25   DefaultAmsPorts = 875
	DefaultAmsPorts_RUNTIME_SYSTEM_26   DefaultAmsPorts = 876
	DefaultAmsPorts_RUNTIME_SYSTEM_27   DefaultAmsPorts = 877
	DefaultAmsPorts_RUNTIME_SYSTEM_28   DefaultAmsPorts = 878
	DefaultAmsPorts_RUNTIME_SYSTEM_29   DefaultAmsPorts = 879
	DefaultAmsPorts_RUNTIME_SYSTEM_30   DefaultAmsPorts = 880
	DefaultAmsPorts_RUNTIME_SYSTEM_31   DefaultAmsPorts = 881
	DefaultAmsPorts_RUNTIME_SYSTEM_32   DefaultAmsPorts = 882
	DefaultAmsPorts_RUNTIME_SYSTEM_33   DefaultAmsPorts = 883
	DefaultAmsPorts_RUNTIME_SYSTEM_34   DefaultAmsPorts = 884
	DefaultAmsPorts_RUNTIME_SYSTEM_35   DefaultAmsPorts = 885
	DefaultAmsPorts_RUNTIME_SYSTEM_36   DefaultAmsPorts = 886
	DefaultAmsPorts_RUNTIME_SYSTEM_37   DefaultAmsPorts = 887
	DefaultAmsPorts_RUNTIME_SYSTEM_38   DefaultAmsPorts = 888
	DefaultAmsPorts_RUNTIME_SYSTEM_39   DefaultAmsPorts = 889
	DefaultAmsPorts_RUNTIME_SYSTEM_40   DefaultAmsPorts = 890
	DefaultAmsPorts_RUNTIME_SYSTEM_41   DefaultAmsPorts = 891
	DefaultAmsPorts_RUNTIME_SYSTEM_42   DefaultAmsPorts = 892
	DefaultAmsPorts_RUNTIME_SYSTEM_43   DefaultAmsPorts = 893
	DefaultAmsPorts_RUNTIME_SYSTEM_44   DefaultAmsPorts = 894
	DefaultAmsPorts_RUNTIME_SYSTEM_45   DefaultAmsPorts = 895
	DefaultAmsPorts_RUNTIME_SYSTEM_46   DefaultAmsPorts = 896
	DefaultAmsPorts_RUNTIME_SYSTEM_47   DefaultAmsPorts = 897
	DefaultAmsPorts_RUNTIME_SYSTEM_48   DefaultAmsPorts = 898
	DefaultAmsPorts_RUNTIME_SYSTEM_49   DefaultAmsPorts = 899
	DefaultAmsPorts_NC                  DefaultAmsPorts = 500
	DefaultAmsPorts_RESERVED            DefaultAmsPorts = 400
	DefaultAmsPorts_IO                  DefaultAmsPorts = 300
	DefaultAmsPorts_REAL_TIME_CORE      DefaultAmsPorts = 200
	DefaultAmsPorts_EVENT_SYSTEM_LOGGER DefaultAmsPorts = 100
)

var DefaultAmsPortsValues []DefaultAmsPorts

func init() {
	_ = errors.New
	DefaultAmsPortsValues = []DefaultAmsPorts{
		DefaultAmsPorts_CAM_CONTROLLER,
		DefaultAmsPorts_RUNTIME_SYSTEM_01,
		DefaultAmsPorts_RUNTIME_SYSTEM_02,
		DefaultAmsPorts_RUNTIME_SYSTEM_03,
		DefaultAmsPorts_RUNTIME_SYSTEM_04,
		DefaultAmsPorts_RUNTIME_SYSTEM_05,
		DefaultAmsPorts_RUNTIME_SYSTEM_06,
		DefaultAmsPorts_RUNTIME_SYSTEM_07,
		DefaultAmsPorts_RUNTIME_SYSTEM_08,
		DefaultAmsPorts_RUNTIME_SYSTEM_09,
		DefaultAmsPorts_RUNTIME_SYSTEM_10,
		DefaultAmsPorts_RUNTIME_SYSTEM_11,
		DefaultAmsPorts_RUNTIME_SYSTEM_12,
		DefaultAmsPorts_RUNTIME_SYSTEM_13,
		DefaultAmsPorts_RUNTIME_SYSTEM_14,
		DefaultAmsPorts_RUNTIME_SYSTEM_15,
		DefaultAmsPorts_RUNTIME_SYSTEM_16,
		DefaultAmsPorts_RUNTIME_SYSTEM_17,
		DefaultAmsPorts_RUNTIME_SYSTEM_18,
		DefaultAmsPorts_RUNTIME_SYSTEM_19,
		DefaultAmsPorts_RUNTIME_SYSTEM_20,
		DefaultAmsPorts_RUNTIME_SYSTEM_21,
		DefaultAmsPorts_RUNTIME_SYSTEM_22,
		DefaultAmsPorts_RUNTIME_SYSTEM_23,
		DefaultAmsPorts_RUNTIME_SYSTEM_24,
		DefaultAmsPorts_RUNTIME_SYSTEM_25,
		DefaultAmsPorts_RUNTIME_SYSTEM_26,
		DefaultAmsPorts_RUNTIME_SYSTEM_27,
		DefaultAmsPorts_RUNTIME_SYSTEM_28,
		DefaultAmsPorts_RUNTIME_SYSTEM_29,
		DefaultAmsPorts_RUNTIME_SYSTEM_30,
		DefaultAmsPorts_RUNTIME_SYSTEM_31,
		DefaultAmsPorts_RUNTIME_SYSTEM_32,
		DefaultAmsPorts_RUNTIME_SYSTEM_33,
		DefaultAmsPorts_RUNTIME_SYSTEM_34,
		DefaultAmsPorts_RUNTIME_SYSTEM_35,
		DefaultAmsPorts_RUNTIME_SYSTEM_36,
		DefaultAmsPorts_RUNTIME_SYSTEM_37,
		DefaultAmsPorts_RUNTIME_SYSTEM_38,
		DefaultAmsPorts_RUNTIME_SYSTEM_39,
		DefaultAmsPorts_RUNTIME_SYSTEM_40,
		DefaultAmsPorts_RUNTIME_SYSTEM_41,
		DefaultAmsPorts_RUNTIME_SYSTEM_42,
		DefaultAmsPorts_RUNTIME_SYSTEM_43,
		DefaultAmsPorts_RUNTIME_SYSTEM_44,
		DefaultAmsPorts_RUNTIME_SYSTEM_45,
		DefaultAmsPorts_RUNTIME_SYSTEM_46,
		DefaultAmsPorts_RUNTIME_SYSTEM_47,
		DefaultAmsPorts_RUNTIME_SYSTEM_48,
		DefaultAmsPorts_RUNTIME_SYSTEM_49,
		DefaultAmsPorts_NC,
		DefaultAmsPorts_RESERVED,
		DefaultAmsPorts_IO,
		DefaultAmsPorts_REAL_TIME_CORE,
		DefaultAmsPorts_EVENT_SYSTEM_LOGGER,
	}
}

func DefaultAmsPortsByValue(value uint16) (enum DefaultAmsPorts, ok bool) {
	switch value {
	case 100:
		return DefaultAmsPorts_EVENT_SYSTEM_LOGGER, true
	case 200:
		return DefaultAmsPorts_REAL_TIME_CORE, true
	case 300:
		return DefaultAmsPorts_IO, true
	case 400:
		return DefaultAmsPorts_RESERVED, true
	case 500:
		return DefaultAmsPorts_NC, true
	case 851:
		return DefaultAmsPorts_RUNTIME_SYSTEM_01, true
	case 852:
		return DefaultAmsPorts_RUNTIME_SYSTEM_02, true
	case 853:
		return DefaultAmsPorts_RUNTIME_SYSTEM_03, true
	case 854:
		return DefaultAmsPorts_RUNTIME_SYSTEM_04, true
	case 855:
		return DefaultAmsPorts_RUNTIME_SYSTEM_05, true
	case 856:
		return DefaultAmsPorts_RUNTIME_SYSTEM_06, true
	case 857:
		return DefaultAmsPorts_RUNTIME_SYSTEM_07, true
	case 858:
		return DefaultAmsPorts_RUNTIME_SYSTEM_08, true
	case 859:
		return DefaultAmsPorts_RUNTIME_SYSTEM_09, true
	case 860:
		return DefaultAmsPorts_RUNTIME_SYSTEM_10, true
	case 861:
		return DefaultAmsPorts_RUNTIME_SYSTEM_11, true
	case 862:
		return DefaultAmsPorts_RUNTIME_SYSTEM_12, true
	case 863:
		return DefaultAmsPorts_RUNTIME_SYSTEM_13, true
	case 864:
		return DefaultAmsPorts_RUNTIME_SYSTEM_14, true
	case 865:
		return DefaultAmsPorts_RUNTIME_SYSTEM_15, true
	case 866:
		return DefaultAmsPorts_RUNTIME_SYSTEM_16, true
	case 867:
		return DefaultAmsPorts_RUNTIME_SYSTEM_17, true
	case 868:
		return DefaultAmsPorts_RUNTIME_SYSTEM_18, true
	case 869:
		return DefaultAmsPorts_RUNTIME_SYSTEM_19, true
	case 870:
		return DefaultAmsPorts_RUNTIME_SYSTEM_20, true
	case 871:
		return DefaultAmsPorts_RUNTIME_SYSTEM_21, true
	case 872:
		return DefaultAmsPorts_RUNTIME_SYSTEM_22, true
	case 873:
		return DefaultAmsPorts_RUNTIME_SYSTEM_23, true
	case 874:
		return DefaultAmsPorts_RUNTIME_SYSTEM_24, true
	case 875:
		return DefaultAmsPorts_RUNTIME_SYSTEM_25, true
	case 876:
		return DefaultAmsPorts_RUNTIME_SYSTEM_26, true
	case 877:
		return DefaultAmsPorts_RUNTIME_SYSTEM_27, true
	case 878:
		return DefaultAmsPorts_RUNTIME_SYSTEM_28, true
	case 879:
		return DefaultAmsPorts_RUNTIME_SYSTEM_29, true
	case 880:
		return DefaultAmsPorts_RUNTIME_SYSTEM_30, true
	case 881:
		return DefaultAmsPorts_RUNTIME_SYSTEM_31, true
	case 882:
		return DefaultAmsPorts_RUNTIME_SYSTEM_32, true
	case 883:
		return DefaultAmsPorts_RUNTIME_SYSTEM_33, true
	case 884:
		return DefaultAmsPorts_RUNTIME_SYSTEM_34, true
	case 885:
		return DefaultAmsPorts_RUNTIME_SYSTEM_35, true
	case 886:
		return DefaultAmsPorts_RUNTIME_SYSTEM_36, true
	case 887:
		return DefaultAmsPorts_RUNTIME_SYSTEM_37, true
	case 888:
		return DefaultAmsPorts_RUNTIME_SYSTEM_38, true
	case 889:
		return DefaultAmsPorts_RUNTIME_SYSTEM_39, true
	case 890:
		return DefaultAmsPorts_RUNTIME_SYSTEM_40, true
	case 891:
		return DefaultAmsPorts_RUNTIME_SYSTEM_41, true
	case 892:
		return DefaultAmsPorts_RUNTIME_SYSTEM_42, true
	case 893:
		return DefaultAmsPorts_RUNTIME_SYSTEM_43, true
	case 894:
		return DefaultAmsPorts_RUNTIME_SYSTEM_44, true
	case 895:
		return DefaultAmsPorts_RUNTIME_SYSTEM_45, true
	case 896:
		return DefaultAmsPorts_RUNTIME_SYSTEM_46, true
	case 897:
		return DefaultAmsPorts_RUNTIME_SYSTEM_47, true
	case 898:
		return DefaultAmsPorts_RUNTIME_SYSTEM_48, true
	case 899:
		return DefaultAmsPorts_RUNTIME_SYSTEM_49, true
	case 900:
		return DefaultAmsPorts_CAM_CONTROLLER, true
	}
	return 0, false
}

func DefaultAmsPortsByName(value string) (enum DefaultAmsPorts, ok bool) {
	switch value {
	case "EVENT_SYSTEM_LOGGER":
		return DefaultAmsPorts_EVENT_SYSTEM_LOGGER, true
	case "REAL_TIME_CORE":
		return DefaultAmsPorts_REAL_TIME_CORE, true
	case "IO":
		return DefaultAmsPorts_IO, true
	case "RESERVED":
		return DefaultAmsPorts_RESERVED, true
	case "NC":
		return DefaultAmsPorts_NC, true
	case "RUNTIME_SYSTEM_01":
		return DefaultAmsPorts_RUNTIME_SYSTEM_01, true
	case "RUNTIME_SYSTEM_02":
		return DefaultAmsPorts_RUNTIME_SYSTEM_02, true
	case "RUNTIME_SYSTEM_03":
		return DefaultAmsPorts_RUNTIME_SYSTEM_03, true
	case "RUNTIME_SYSTEM_04":
		return DefaultAmsPorts_RUNTIME_SYSTEM_04, true
	case "RUNTIME_SYSTEM_05":
		return DefaultAmsPorts_RUNTIME_SYSTEM_05, true
	case "RUNTIME_SYSTEM_06":
		return DefaultAmsPorts_RUNTIME_SYSTEM_06, true
	case "RUNTIME_SYSTEM_07":
		return DefaultAmsPorts_RUNTIME_SYSTEM_07, true
	case "RUNTIME_SYSTEM_08":
		return DefaultAmsPorts_RUNTIME_SYSTEM_08, true
	case "RUNTIME_SYSTEM_09":
		return DefaultAmsPorts_RUNTIME_SYSTEM_09, true
	case "RUNTIME_SYSTEM_10":
		return DefaultAmsPorts_RUNTIME_SYSTEM_10, true
	case "RUNTIME_SYSTEM_11":
		return DefaultAmsPorts_RUNTIME_SYSTEM_11, true
	case "RUNTIME_SYSTEM_12":
		return DefaultAmsPorts_RUNTIME_SYSTEM_12, true
	case "RUNTIME_SYSTEM_13":
		return DefaultAmsPorts_RUNTIME_SYSTEM_13, true
	case "RUNTIME_SYSTEM_14":
		return DefaultAmsPorts_RUNTIME_SYSTEM_14, true
	case "RUNTIME_SYSTEM_15":
		return DefaultAmsPorts_RUNTIME_SYSTEM_15, true
	case "RUNTIME_SYSTEM_16":
		return DefaultAmsPorts_RUNTIME_SYSTEM_16, true
	case "RUNTIME_SYSTEM_17":
		return DefaultAmsPorts_RUNTIME_SYSTEM_17, true
	case "RUNTIME_SYSTEM_18":
		return DefaultAmsPorts_RUNTIME_SYSTEM_18, true
	case "RUNTIME_SYSTEM_19":
		return DefaultAmsPorts_RUNTIME_SYSTEM_19, true
	case "RUNTIME_SYSTEM_20":
		return DefaultAmsPorts_RUNTIME_SYSTEM_20, true
	case "RUNTIME_SYSTEM_21":
		return DefaultAmsPorts_RUNTIME_SYSTEM_21, true
	case "RUNTIME_SYSTEM_22":
		return DefaultAmsPorts_RUNTIME_SYSTEM_22, true
	case "RUNTIME_SYSTEM_23":
		return DefaultAmsPorts_RUNTIME_SYSTEM_23, true
	case "RUNTIME_SYSTEM_24":
		return DefaultAmsPorts_RUNTIME_SYSTEM_24, true
	case "RUNTIME_SYSTEM_25":
		return DefaultAmsPorts_RUNTIME_SYSTEM_25, true
	case "RUNTIME_SYSTEM_26":
		return DefaultAmsPorts_RUNTIME_SYSTEM_26, true
	case "RUNTIME_SYSTEM_27":
		return DefaultAmsPorts_RUNTIME_SYSTEM_27, true
	case "RUNTIME_SYSTEM_28":
		return DefaultAmsPorts_RUNTIME_SYSTEM_28, true
	case "RUNTIME_SYSTEM_29":
		return DefaultAmsPorts_RUNTIME_SYSTEM_29, true
	case "RUNTIME_SYSTEM_30":
		return DefaultAmsPorts_RUNTIME_SYSTEM_30, true
	case "RUNTIME_SYSTEM_31":
		return DefaultAmsPorts_RUNTIME_SYSTEM_31, true
	case "RUNTIME_SYSTEM_32":
		return DefaultAmsPorts_RUNTIME_SYSTEM_32, true
	case "RUNTIME_SYSTEM_33":
		return DefaultAmsPorts_RUNTIME_SYSTEM_33, true
	case "RUNTIME_SYSTEM_34":
		return DefaultAmsPorts_RUNTIME_SYSTEM_34, true
	case "RUNTIME_SYSTEM_35":
		return DefaultAmsPorts_RUNTIME_SYSTEM_35, true
	case "RUNTIME_SYSTEM_36":
		return DefaultAmsPorts_RUNTIME_SYSTEM_36, true
	case "RUNTIME_SYSTEM_37":
		return DefaultAmsPorts_RUNTIME_SYSTEM_37, true
	case "RUNTIME_SYSTEM_38":
		return DefaultAmsPorts_RUNTIME_SYSTEM_38, true
	case "RUNTIME_SYSTEM_39":
		return DefaultAmsPorts_RUNTIME_SYSTEM_39, true
	case "RUNTIME_SYSTEM_40":
		return DefaultAmsPorts_RUNTIME_SYSTEM_40, true
	case "RUNTIME_SYSTEM_41":
		return DefaultAmsPorts_RUNTIME_SYSTEM_41, true
	case "RUNTIME_SYSTEM_42":
		return DefaultAmsPorts_RUNTIME_SYSTEM_42, true
	case "RUNTIME_SYSTEM_43":
		return DefaultAmsPorts_RUNTIME_SYSTEM_43, true
	case "RUNTIME_SYSTEM_44":
		return DefaultAmsPorts_RUNTIME_SYSTEM_44, true
	case "RUNTIME_SYSTEM_45":
		return DefaultAmsPorts_RUNTIME_SYSTEM_45, true
	case "RUNTIME_SYSTEM_46":
		return DefaultAmsPorts_RUNTIME_SYSTEM_46, true
	case "RUNTIME_SYSTEM_47":
		return DefaultAmsPorts_RUNTIME_SYSTEM_47, true
	case "RUNTIME_SYSTEM_48":
		return DefaultAmsPorts_RUNTIME_SYSTEM_48, true
	case "RUNTIME_SYSTEM_49":
		return DefaultAmsPorts_RUNTIME_SYSTEM_49, true
	case "CAM_CONTROLLER":
		return DefaultAmsPorts_CAM_CONTROLLER, true
	}
	return 0, false
}

func DefaultAmsPortsKnows(value uint16) bool {
	for _, typeValue := range DefaultAmsPortsValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDefaultAmsPorts(structType any) DefaultAmsPorts {
	castFunc := func(typ any) DefaultAmsPorts {
		if sDefaultAmsPorts, ok := typ.(DefaultAmsPorts); ok {
			return sDefaultAmsPorts
		}
		return 0
	}
	return castFunc(structType)
}

func (m DefaultAmsPorts) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m DefaultAmsPorts) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DefaultAmsPortsParse(ctx context.Context, theBytes []byte) (DefaultAmsPorts, error) {
	return DefaultAmsPortsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DefaultAmsPortsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DefaultAmsPorts, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("DefaultAmsPorts", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DefaultAmsPorts")
	}
	if enum, ok := DefaultAmsPortsByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for DefaultAmsPorts")
		return DefaultAmsPorts(val), nil
	} else {
		return enum, nil
	}
}

func (e DefaultAmsPorts) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DefaultAmsPorts) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("DefaultAmsPorts", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DefaultAmsPorts) PLC4XEnumName() string {
	switch e {
	case DefaultAmsPorts_EVENT_SYSTEM_LOGGER:
		return "EVENT_SYSTEM_LOGGER"
	case DefaultAmsPorts_REAL_TIME_CORE:
		return "REAL_TIME_CORE"
	case DefaultAmsPorts_IO:
		return "IO"
	case DefaultAmsPorts_RESERVED:
		return "RESERVED"
	case DefaultAmsPorts_NC:
		return "NC"
	case DefaultAmsPorts_RUNTIME_SYSTEM_01:
		return "RUNTIME_SYSTEM_01"
	case DefaultAmsPorts_RUNTIME_SYSTEM_02:
		return "RUNTIME_SYSTEM_02"
	case DefaultAmsPorts_RUNTIME_SYSTEM_03:
		return "RUNTIME_SYSTEM_03"
	case DefaultAmsPorts_RUNTIME_SYSTEM_04:
		return "RUNTIME_SYSTEM_04"
	case DefaultAmsPorts_RUNTIME_SYSTEM_05:
		return "RUNTIME_SYSTEM_05"
	case DefaultAmsPorts_RUNTIME_SYSTEM_06:
		return "RUNTIME_SYSTEM_06"
	case DefaultAmsPorts_RUNTIME_SYSTEM_07:
		return "RUNTIME_SYSTEM_07"
	case DefaultAmsPorts_RUNTIME_SYSTEM_08:
		return "RUNTIME_SYSTEM_08"
	case DefaultAmsPorts_RUNTIME_SYSTEM_09:
		return "RUNTIME_SYSTEM_09"
	case DefaultAmsPorts_RUNTIME_SYSTEM_10:
		return "RUNTIME_SYSTEM_10"
	case DefaultAmsPorts_RUNTIME_SYSTEM_11:
		return "RUNTIME_SYSTEM_11"
	case DefaultAmsPorts_RUNTIME_SYSTEM_12:
		return "RUNTIME_SYSTEM_12"
	case DefaultAmsPorts_RUNTIME_SYSTEM_13:
		return "RUNTIME_SYSTEM_13"
	case DefaultAmsPorts_RUNTIME_SYSTEM_14:
		return "RUNTIME_SYSTEM_14"
	case DefaultAmsPorts_RUNTIME_SYSTEM_15:
		return "RUNTIME_SYSTEM_15"
	case DefaultAmsPorts_RUNTIME_SYSTEM_16:
		return "RUNTIME_SYSTEM_16"
	case DefaultAmsPorts_RUNTIME_SYSTEM_17:
		return "RUNTIME_SYSTEM_17"
	case DefaultAmsPorts_RUNTIME_SYSTEM_18:
		return "RUNTIME_SYSTEM_18"
	case DefaultAmsPorts_RUNTIME_SYSTEM_19:
		return "RUNTIME_SYSTEM_19"
	case DefaultAmsPorts_RUNTIME_SYSTEM_20:
		return "RUNTIME_SYSTEM_20"
	case DefaultAmsPorts_RUNTIME_SYSTEM_21:
		return "RUNTIME_SYSTEM_21"
	case DefaultAmsPorts_RUNTIME_SYSTEM_22:
		return "RUNTIME_SYSTEM_22"
	case DefaultAmsPorts_RUNTIME_SYSTEM_23:
		return "RUNTIME_SYSTEM_23"
	case DefaultAmsPorts_RUNTIME_SYSTEM_24:
		return "RUNTIME_SYSTEM_24"
	case DefaultAmsPorts_RUNTIME_SYSTEM_25:
		return "RUNTIME_SYSTEM_25"
	case DefaultAmsPorts_RUNTIME_SYSTEM_26:
		return "RUNTIME_SYSTEM_26"
	case DefaultAmsPorts_RUNTIME_SYSTEM_27:
		return "RUNTIME_SYSTEM_27"
	case DefaultAmsPorts_RUNTIME_SYSTEM_28:
		return "RUNTIME_SYSTEM_28"
	case DefaultAmsPorts_RUNTIME_SYSTEM_29:
		return "RUNTIME_SYSTEM_29"
	case DefaultAmsPorts_RUNTIME_SYSTEM_30:
		return "RUNTIME_SYSTEM_30"
	case DefaultAmsPorts_RUNTIME_SYSTEM_31:
		return "RUNTIME_SYSTEM_31"
	case DefaultAmsPorts_RUNTIME_SYSTEM_32:
		return "RUNTIME_SYSTEM_32"
	case DefaultAmsPorts_RUNTIME_SYSTEM_33:
		return "RUNTIME_SYSTEM_33"
	case DefaultAmsPorts_RUNTIME_SYSTEM_34:
		return "RUNTIME_SYSTEM_34"
	case DefaultAmsPorts_RUNTIME_SYSTEM_35:
		return "RUNTIME_SYSTEM_35"
	case DefaultAmsPorts_RUNTIME_SYSTEM_36:
		return "RUNTIME_SYSTEM_36"
	case DefaultAmsPorts_RUNTIME_SYSTEM_37:
		return "RUNTIME_SYSTEM_37"
	case DefaultAmsPorts_RUNTIME_SYSTEM_38:
		return "RUNTIME_SYSTEM_38"
	case DefaultAmsPorts_RUNTIME_SYSTEM_39:
		return "RUNTIME_SYSTEM_39"
	case DefaultAmsPorts_RUNTIME_SYSTEM_40:
		return "RUNTIME_SYSTEM_40"
	case DefaultAmsPorts_RUNTIME_SYSTEM_41:
		return "RUNTIME_SYSTEM_41"
	case DefaultAmsPorts_RUNTIME_SYSTEM_42:
		return "RUNTIME_SYSTEM_42"
	case DefaultAmsPorts_RUNTIME_SYSTEM_43:
		return "RUNTIME_SYSTEM_43"
	case DefaultAmsPorts_RUNTIME_SYSTEM_44:
		return "RUNTIME_SYSTEM_44"
	case DefaultAmsPorts_RUNTIME_SYSTEM_45:
		return "RUNTIME_SYSTEM_45"
	case DefaultAmsPorts_RUNTIME_SYSTEM_46:
		return "RUNTIME_SYSTEM_46"
	case DefaultAmsPorts_RUNTIME_SYSTEM_47:
		return "RUNTIME_SYSTEM_47"
	case DefaultAmsPorts_RUNTIME_SYSTEM_48:
		return "RUNTIME_SYSTEM_48"
	case DefaultAmsPorts_RUNTIME_SYSTEM_49:
		return "RUNTIME_SYSTEM_49"
	case DefaultAmsPorts_CAM_CONTROLLER:
		return "CAM_CONTROLLER"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e DefaultAmsPorts) String() string {
	return e.PLC4XEnumName()
}
