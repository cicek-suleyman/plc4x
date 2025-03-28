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

// OpcuaNodeIdServicesVariableThree is an enum
type OpcuaNodeIdServicesVariableThree int32

type IOpcuaNodeIdServicesVariableThree interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableThree_ThreeDVectorType_VectorUnit                     OpcuaNodeIdServicesVariableThree = 17717
	OpcuaNodeIdServicesVariableThree_ThreeDVectorType_X                              OpcuaNodeIdServicesVariableThree = 18769
	OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Y                              OpcuaNodeIdServicesVariableThree = 18770
	OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Z                              OpcuaNodeIdServicesVariableThree = 18771
	OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_LengthUnit       OpcuaNodeIdServicesVariableThree = 18775
	OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_X                OpcuaNodeIdServicesVariableThree = 18776
	OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Y                OpcuaNodeIdServicesVariableThree = 18777
	OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Z                OpcuaNodeIdServicesVariableThree = 18778
	OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_AngleUnit                 OpcuaNodeIdServicesVariableThree = 18782
	OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_A                         OpcuaNodeIdServicesVariableThree = 18783
	OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_B                         OpcuaNodeIdServicesVariableThree = 18784
	OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_C                         OpcuaNodeIdServicesVariableThree = 18785
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation                     OpcuaNodeIdServicesVariableThree = 18792
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Constant                        OpcuaNodeIdServicesVariableThree = 18793
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_BaseFrame                       OpcuaNodeIdServicesVariableThree = 18794
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_FixedBase                       OpcuaNodeIdServicesVariableThree = 18795
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates            OpcuaNodeIdServicesVariableThree = 18796
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_LengthUnit OpcuaNodeIdServicesVariableThree = 18797
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_X          OpcuaNodeIdServicesVariableThree = 18798
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Y          OpcuaNodeIdServicesVariableThree = 18799
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Z          OpcuaNodeIdServicesVariableThree = 18800
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_AngleUnit           OpcuaNodeIdServicesVariableThree = 19073
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_A                   OpcuaNodeIdServicesVariableThree = 19074
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_B                   OpcuaNodeIdServicesVariableThree = 19075
	OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_C                   OpcuaNodeIdServicesVariableThree = 19076
)

var OpcuaNodeIdServicesVariableThreeValues []OpcuaNodeIdServicesVariableThree

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableThreeValues = []OpcuaNodeIdServicesVariableThree{
		OpcuaNodeIdServicesVariableThree_ThreeDVectorType_VectorUnit,
		OpcuaNodeIdServicesVariableThree_ThreeDVectorType_X,
		OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Y,
		OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Z,
		OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_LengthUnit,
		OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_X,
		OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Y,
		OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Z,
		OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_AngleUnit,
		OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_A,
		OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_B,
		OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_C,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Constant,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_BaseFrame,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_FixedBase,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_LengthUnit,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_X,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Y,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Z,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_AngleUnit,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_A,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_B,
		OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_C,
	}
}

func OpcuaNodeIdServicesVariableThreeByValue(value int32) (enum OpcuaNodeIdServicesVariableThree, ok bool) {
	switch value {
	case 17717:
		return OpcuaNodeIdServicesVariableThree_ThreeDVectorType_VectorUnit, true
	case 18769:
		return OpcuaNodeIdServicesVariableThree_ThreeDVectorType_X, true
	case 18770:
		return OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Y, true
	case 18771:
		return OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Z, true
	case 18775:
		return OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_LengthUnit, true
	case 18776:
		return OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_X, true
	case 18777:
		return OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Y, true
	case 18778:
		return OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Z, true
	case 18782:
		return OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_AngleUnit, true
	case 18783:
		return OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_A, true
	case 18784:
		return OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_B, true
	case 18785:
		return OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_C, true
	case 18792:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation, true
	case 18793:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Constant, true
	case 18794:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_BaseFrame, true
	case 18795:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_FixedBase, true
	case 18796:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates, true
	case 18797:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_LengthUnit, true
	case 18798:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_X, true
	case 18799:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Y, true
	case 18800:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Z, true
	case 19073:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_AngleUnit, true
	case 19074:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_A, true
	case 19075:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_B, true
	case 19076:
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_C, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableThreeByName(value string) (enum OpcuaNodeIdServicesVariableThree, ok bool) {
	switch value {
	case "ThreeDVectorType_VectorUnit":
		return OpcuaNodeIdServicesVariableThree_ThreeDVectorType_VectorUnit, true
	case "ThreeDVectorType_X":
		return OpcuaNodeIdServicesVariableThree_ThreeDVectorType_X, true
	case "ThreeDVectorType_Y":
		return OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Y, true
	case "ThreeDVectorType_Z":
		return OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Z, true
	case "ThreeDCartesianCoordinatesType_LengthUnit":
		return OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_LengthUnit, true
	case "ThreeDCartesianCoordinatesType_X":
		return OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_X, true
	case "ThreeDCartesianCoordinatesType_Y":
		return OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Y, true
	case "ThreeDCartesianCoordinatesType_Z":
		return OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Z, true
	case "ThreeDOrientationType_AngleUnit":
		return OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_AngleUnit, true
	case "ThreeDOrientationType_A":
		return OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_A, true
	case "ThreeDOrientationType_B":
		return OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_B, true
	case "ThreeDOrientationType_C":
		return OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_C, true
	case "ThreeDFrameType_Orientation":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation, true
	case "ThreeDFrameType_Constant":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Constant, true
	case "ThreeDFrameType_BaseFrame":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_BaseFrame, true
	case "ThreeDFrameType_FixedBase":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_FixedBase, true
	case "ThreeDFrameType_CartesianCoordinates":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates, true
	case "ThreeDFrameType_CartesianCoordinates_LengthUnit":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_LengthUnit, true
	case "ThreeDFrameType_CartesianCoordinates_X":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_X, true
	case "ThreeDFrameType_CartesianCoordinates_Y":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Y, true
	case "ThreeDFrameType_CartesianCoordinates_Z":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Z, true
	case "ThreeDFrameType_Orientation_AngleUnit":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_AngleUnit, true
	case "ThreeDFrameType_Orientation_A":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_A, true
	case "ThreeDFrameType_Orientation_B":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_B, true
	case "ThreeDFrameType_Orientation_C":
		return OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_C, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableThreeKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableThreeValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableThree(structType any) OpcuaNodeIdServicesVariableThree {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableThree {
		if sOpcuaNodeIdServicesVariableThree, ok := typ.(OpcuaNodeIdServicesVariableThree); ok {
			return sOpcuaNodeIdServicesVariableThree
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableThree) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableThree) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableThreeParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableThree, error) {
	return OpcuaNodeIdServicesVariableThreeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableThreeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableThree, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableThree", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableThree")
	}
	if enum, ok := OpcuaNodeIdServicesVariableThreeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableThree")
		return OpcuaNodeIdServicesVariableThree(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableThree) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableThree) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableThree", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableThree) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableThree_ThreeDVectorType_VectorUnit:
		return "ThreeDVectorType_VectorUnit"
	case OpcuaNodeIdServicesVariableThree_ThreeDVectorType_X:
		return "ThreeDVectorType_X"
	case OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Y:
		return "ThreeDVectorType_Y"
	case OpcuaNodeIdServicesVariableThree_ThreeDVectorType_Z:
		return "ThreeDVectorType_Z"
	case OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_LengthUnit:
		return "ThreeDCartesianCoordinatesType_LengthUnit"
	case OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_X:
		return "ThreeDCartesianCoordinatesType_X"
	case OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Y:
		return "ThreeDCartesianCoordinatesType_Y"
	case OpcuaNodeIdServicesVariableThree_ThreeDCartesianCoordinatesType_Z:
		return "ThreeDCartesianCoordinatesType_Z"
	case OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_AngleUnit:
		return "ThreeDOrientationType_AngleUnit"
	case OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_A:
		return "ThreeDOrientationType_A"
	case OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_B:
		return "ThreeDOrientationType_B"
	case OpcuaNodeIdServicesVariableThree_ThreeDOrientationType_C:
		return "ThreeDOrientationType_C"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation:
		return "ThreeDFrameType_Orientation"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Constant:
		return "ThreeDFrameType_Constant"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_BaseFrame:
		return "ThreeDFrameType_BaseFrame"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_FixedBase:
		return "ThreeDFrameType_FixedBase"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates:
		return "ThreeDFrameType_CartesianCoordinates"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_LengthUnit:
		return "ThreeDFrameType_CartesianCoordinates_LengthUnit"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_X:
		return "ThreeDFrameType_CartesianCoordinates_X"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Y:
		return "ThreeDFrameType_CartesianCoordinates_Y"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_CartesianCoordinates_Z:
		return "ThreeDFrameType_CartesianCoordinates_Z"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_AngleUnit:
		return "ThreeDFrameType_Orientation_AngleUnit"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_A:
		return "ThreeDFrameType_Orientation_A"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_B:
		return "ThreeDFrameType_Orientation_B"
	case OpcuaNodeIdServicesVariableThree_ThreeDFrameType_Orientation_C:
		return "ThreeDFrameType_Orientation_C"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableThree) String() string {
	return e.PLC4XEnumName()
}
