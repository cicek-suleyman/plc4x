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

// OpcuaNodeIdServicesVariableAdd is an enum
type OpcuaNodeIdServicesVariableAdd int32

type IOpcuaNodeIdServicesVariableAdd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableAdd_AddCommentMethodType_InputArguments                     OpcuaNodeIdServicesVariableAdd = 3864
	OpcuaNodeIdServicesVariableAdd_AddCertificateMethodType_InputArguments                 OpcuaNodeIdServicesVariableAdd = 12519
	OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_InputArguments          OpcuaNodeIdServicesVariableAdd = 14502
	OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_OutputArguments         OpcuaNodeIdServicesVariableAdd = 14503
	OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_InputArguments             OpcuaNodeIdServicesVariableAdd = 14505
	OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_OutputArguments            OpcuaNodeIdServicesVariableAdd = 14506
	OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_InputArguments               OpcuaNodeIdServicesVariableAdd = 15467
	OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_OutputArguments              OpcuaNodeIdServicesVariableAdd = 15468
	OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_InputArguments              OpcuaNodeIdServicesVariableAdd = 15497
	OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_OutputArguments             OpcuaNodeIdServicesVariableAdd = 15498
	OpcuaNodeIdServicesVariableAdd_AddIdentityMethodType_InputArguments                    OpcuaNodeIdServicesVariableAdd = 15637
	OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_InputArguments                        OpcuaNodeIdServicesVariableAdd = 16003
	OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_OutputArguments                       OpcuaNodeIdServicesVariableAdd = 16004
	OpcuaNodeIdServicesVariableAdd_AddApplicationMethodType_InputArguments                 OpcuaNodeIdServicesVariableAdd = 16185
	OpcuaNodeIdServicesVariableAdd_AddEndpointMethodType_InputArguments                    OpcuaNodeIdServicesVariableAdd = 16189
	OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_InputArguments                  OpcuaNodeIdServicesVariableAdd = 16718
	OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_OutputArguments                 OpcuaNodeIdServicesVariableAdd = 16719
	OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_InputArguments  OpcuaNodeIdServicesVariableAdd = 17031
	OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_OutputArguments OpcuaNodeIdServicesVariableAdd = 17032
	OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_InputArguments     OpcuaNodeIdServicesVariableAdd = 17043
	OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_OutputArguments    OpcuaNodeIdServicesVariableAdd = 17066
	OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_InputArguments               OpcuaNodeIdServicesVariableAdd = 17068
	OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_OutputArguments              OpcuaNodeIdServicesVariableAdd = 17069
	OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_InputArguments           OpcuaNodeIdServicesVariableAdd = 23822
	OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_OutputArguments          OpcuaNodeIdServicesVariableAdd = 23823
	OpcuaNodeIdServicesVariableAdd_AddUserMethodType_InputArguments                        OpcuaNodeIdServicesVariableAdd = 24283
	OpcuaNodeIdServicesVariableAdd_AddPriorityMappingEntryMethodType_InputArguments        OpcuaNodeIdServicesVariableAdd = 25234
	OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_InputArguments         OpcuaNodeIdServicesVariableAdd = 25289
	OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_OutputArguments        OpcuaNodeIdServicesVariableAdd = 25290
	OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_InputArguments                  OpcuaNodeIdServicesVariableAdd = 25377
	OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_OutputArguments                 OpcuaNodeIdServicesVariableAdd = 25378
	OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_InputArguments            OpcuaNodeIdServicesVariableAdd = 25382
	OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_OutputArguments           OpcuaNodeIdServicesVariableAdd = 25383
)

var OpcuaNodeIdServicesVariableAddValues []OpcuaNodeIdServicesVariableAdd

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableAddValues = []OpcuaNodeIdServicesVariableAdd{
		OpcuaNodeIdServicesVariableAdd_AddCommentMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddCertificateMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddIdentityMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddApplicationMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddEndpointMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddUserMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPriorityMappingEntryMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_InputArguments,
		OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableAddByValue(value int32) (enum OpcuaNodeIdServicesVariableAdd, ok bool) {
	switch value {
	case 12519:
		return OpcuaNodeIdServicesVariableAdd_AddCertificateMethodType_InputArguments, true
	case 14502:
		return OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_InputArguments, true
	case 14503:
		return OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_OutputArguments, true
	case 14505:
		return OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_InputArguments, true
	case 14506:
		return OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_OutputArguments, true
	case 15467:
		return OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_InputArguments, true
	case 15468:
		return OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_OutputArguments, true
	case 15497:
		return OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_InputArguments, true
	case 15498:
		return OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_OutputArguments, true
	case 15637:
		return OpcuaNodeIdServicesVariableAdd_AddIdentityMethodType_InputArguments, true
	case 16003:
		return OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_InputArguments, true
	case 16004:
		return OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_OutputArguments, true
	case 16185:
		return OpcuaNodeIdServicesVariableAdd_AddApplicationMethodType_InputArguments, true
	case 16189:
		return OpcuaNodeIdServicesVariableAdd_AddEndpointMethodType_InputArguments, true
	case 16718:
		return OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_InputArguments, true
	case 16719:
		return OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_OutputArguments, true
	case 17031:
		return OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_InputArguments, true
	case 17032:
		return OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_OutputArguments, true
	case 17043:
		return OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_InputArguments, true
	case 17066:
		return OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_OutputArguments, true
	case 17068:
		return OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_InputArguments, true
	case 17069:
		return OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_OutputArguments, true
	case 23822:
		return OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_InputArguments, true
	case 23823:
		return OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_OutputArguments, true
	case 24283:
		return OpcuaNodeIdServicesVariableAdd_AddUserMethodType_InputArguments, true
	case 25234:
		return OpcuaNodeIdServicesVariableAdd_AddPriorityMappingEntryMethodType_InputArguments, true
	case 25289:
		return OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_InputArguments, true
	case 25290:
		return OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_OutputArguments, true
	case 25377:
		return OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_InputArguments, true
	case 25378:
		return OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_OutputArguments, true
	case 25382:
		return OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_InputArguments, true
	case 25383:
		return OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_OutputArguments, true
	case 3864:
		return OpcuaNodeIdServicesVariableAdd_AddCommentMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAddByName(value string) (enum OpcuaNodeIdServicesVariableAdd, ok bool) {
	switch value {
	case "AddCertificateMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddCertificateMethodType_InputArguments, true
	case "AddPublishedDataItemsMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_InputArguments, true
	case "AddPublishedDataItemsMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_OutputArguments, true
	case "AddPublishedEventsMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_InputArguments, true
	case "AddPublishedEventsMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_OutputArguments, true
	case "AddSecurityGroupMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_InputArguments, true
	case "AddSecurityGroupMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_OutputArguments, true
	case "AddExtensionFieldMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_InputArguments, true
	case "AddExtensionFieldMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_OutputArguments, true
	case "AddIdentityMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddIdentityMethodType_InputArguments, true
	case "AddRoleMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_InputArguments, true
	case "AddRoleMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_OutputArguments, true
	case "AddApplicationMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddApplicationMethodType_InputArguments, true
	case "AddEndpointMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddEndpointMethodType_InputArguments, true
	case "AddConnectionMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_InputArguments, true
	case "AddConnectionMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_OutputArguments, true
	case "AddPublishedDataItemsTemplateMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_InputArguments, true
	case "AddPublishedDataItemsTemplateMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_OutputArguments, true
	case "AddPublishedEventsTemplateMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_InputArguments, true
	case "AddPublishedEventsTemplateMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_OutputArguments, true
	case "AddDataSetFolderMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_InputArguments, true
	case "AddDataSetFolderMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_OutputArguments, true
	case "AddSubscribedDataSetMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_InputArguments, true
	case "AddSubscribedDataSetMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_OutputArguments, true
	case "AddUserMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddUserMethodType_InputArguments, true
	case "AddPriorityMappingEntryMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPriorityMappingEntryMethodType_InputArguments, true
	case "AddSecurityGroupFolderMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_InputArguments, true
	case "AddSecurityGroupFolderMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_OutputArguments, true
	case "AddPushTargetMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_InputArguments, true
	case "AddPushTargetMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_OutputArguments, true
	case "AddPushTargetFolderMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_InputArguments, true
	case "AddPushTargetFolderMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_OutputArguments, true
	case "AddCommentMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableAdd_AddCommentMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAddKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableAddValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableAdd(structType any) OpcuaNodeIdServicesVariableAdd {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableAdd {
		if sOpcuaNodeIdServicesVariableAdd, ok := typ.(OpcuaNodeIdServicesVariableAdd); ok {
			return sOpcuaNodeIdServicesVariableAdd
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableAdd) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableAdd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableAddParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableAdd, error) {
	return OpcuaNodeIdServicesVariableAddParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableAddParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableAdd, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableAdd", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableAdd")
	}
	if enum, ok := OpcuaNodeIdServicesVariableAddByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableAdd")
		return OpcuaNodeIdServicesVariableAdd(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableAdd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableAdd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableAdd", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableAdd) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableAdd_AddCertificateMethodType_InputArguments:
		return "AddCertificateMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_InputArguments:
		return "AddPublishedDataItemsMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsMethodType_OutputArguments:
		return "AddPublishedDataItemsMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_InputArguments:
		return "AddPublishedEventsMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPublishedEventsMethodType_OutputArguments:
		return "AddPublishedEventsMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_InputArguments:
		return "AddSecurityGroupMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddSecurityGroupMethodType_OutputArguments:
		return "AddSecurityGroupMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_InputArguments:
		return "AddExtensionFieldMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddExtensionFieldMethodType_OutputArguments:
		return "AddExtensionFieldMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddIdentityMethodType_InputArguments:
		return "AddIdentityMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_InputArguments:
		return "AddRoleMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddRoleMethodType_OutputArguments:
		return "AddRoleMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddApplicationMethodType_InputArguments:
		return "AddApplicationMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddEndpointMethodType_InputArguments:
		return "AddEndpointMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_InputArguments:
		return "AddConnectionMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddConnectionMethodType_OutputArguments:
		return "AddConnectionMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_InputArguments:
		return "AddPublishedDataItemsTemplateMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPublishedDataItemsTemplateMethodType_OutputArguments:
		return "AddPublishedDataItemsTemplateMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_InputArguments:
		return "AddPublishedEventsTemplateMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPublishedEventsTemplateMethodType_OutputArguments:
		return "AddPublishedEventsTemplateMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_InputArguments:
		return "AddDataSetFolderMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddDataSetFolderMethodType_OutputArguments:
		return "AddDataSetFolderMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_InputArguments:
		return "AddSubscribedDataSetMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddSubscribedDataSetMethodType_OutputArguments:
		return "AddSubscribedDataSetMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddUserMethodType_InputArguments:
		return "AddUserMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPriorityMappingEntryMethodType_InputArguments:
		return "AddPriorityMappingEntryMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_InputArguments:
		return "AddSecurityGroupFolderMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddSecurityGroupFolderMethodType_OutputArguments:
		return "AddSecurityGroupFolderMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_InputArguments:
		return "AddPushTargetMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPushTargetMethodType_OutputArguments:
		return "AddPushTargetMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_InputArguments:
		return "AddPushTargetFolderMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddPushTargetFolderMethodType_OutputArguments:
		return "AddPushTargetFolderMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableAdd_AddCommentMethodType_InputArguments:
		return "AddCommentMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableAdd) String() string {
	return e.PLC4XEnumName()
}
