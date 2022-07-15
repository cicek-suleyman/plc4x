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

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type CbusParserHelper struct {
}

func (m CbusParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (interface{}, error) {
	switch typeName {
	case "CALData":
		var requestContext model.RequestContext
		return model.CALDataParse(io, requestContext)
	case "ApplicationAddress2":
		return model.ApplicationAddress2Parse(io)
	case "LabelOptions":
		return model.LabelOptionsParse(io)
	case "ApplicationAddress1":
		return model.ApplicationAddress1Parse(io)
	case "Checksum":
		return model.ChecksumParse(io)
	case "RequestContext":
		return model.RequestContextParse(io)
	case "CALReply":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.CALReplyParse(io, cBusOptions, requestContext)
	case "CustomManufacturer":
		return model.CustomManufacturerParse(io)
	case "NetworkRoute":
		return model.NetworkRouteParse(io)
	case "NetworkNumber":
		return model.NetworkNumberParse(io)
	case "RequestTermination":
		return model.RequestTerminationParse(io)
	case "ReplyOrConfirmation":
		var cBusOptions model.CBusOptions
		messageLength, err := utils.StrToUint16(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		var requestContext model.RequestContext
		return model.ReplyOrConfirmationParse(io, cBusOptions, messageLength, requestContext)
	case "StandardFormatStatusReply":
		return model.StandardFormatStatusReplyParse(io)
	case "CBusMessage":
		isResponse, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		var requestContext model.RequestContext
		var cBusOptions model.CBusOptions
		messageLength, err := utils.StrToUint16(arguments[3])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.CBusMessageParse(io, isResponse, requestContext, cBusOptions, messageLength)
	case "ResponseTermination":
		return model.ResponseTerminationParse(io)
	case "CBusOptions":
		return model.CBusOptionsParse(io)
	case "SALData":
		return model.SALDataParse(io)
	case "CBusCommand":
		var cBusOptions model.CBusOptions
		return model.CBusCommandParse(io, cBusOptions)
	case "IdentifyReplyCommand":
		attribute, _ := model.AttributeByName(arguments[0])
		numBytes, err := utils.StrToUint8(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.IdentifyReplyCommandParse(io, attribute, numBytes)
	case "CALDataOrSetParameter":
		return model.CALDataOrSetParameterParse(io)
	case "IdentifyReplyCommandUnitSummary":
		return model.IdentifyReplyCommandUnitSummaryParse(io)
	case "CBusConstants":
		return model.CBusConstantsParse(io)
	case "BridgeCount":
		return model.BridgeCountParse(io)
	case "PowerUp":
		return model.PowerUpParse(io)
	case "Reply":
		var cBusOptions model.CBusOptions
		replyLength, err := utils.StrToUint16(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		var requestContext model.RequestContext
		return model.ReplyParse(io, cBusOptions, replyLength, requestContext)
	case "SerialInterfaceAddress":
		return model.SerialInterfaceAddressParse(io)
	case "InterfaceOptions1PowerUpSettings":
		return model.InterfaceOptions1PowerUpSettingsParse(io)
	case "BridgeAddress":
		return model.BridgeAddressParse(io)
	case "MonitoredSAL":
		var cBusOptions model.CBusOptions
		return model.MonitoredSALParse(io, cBusOptions)
	case "ParameterChange":
		return model.ParameterChangeParse(io)
	case "StatusByte":
		return model.StatusByteParse(io)
	case "ReplyNetwork":
		return model.ReplyNetworkParse(io)
	case "SerialNumber":
		return model.SerialNumberParse(io)
	case "ExtendedStatusHeader":
		return model.ExtendedStatusHeaderParse(io)
	case "CustomTypes":
		return model.CustomTypesParse(io)
	case "CommandHeader":
		return model.CommandHeaderParse(io)
	case "Confirmation":
		return model.ConfirmationParse(io)
	case "CBusPointToMultiPointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToMultiPointCommandParse(io, cBusOptions)
	case "StatusHeader":
		return model.StatusHeaderParse(io)
	case "StatusRequest":
		return model.StatusRequestParse(io)
	case "EncodedReply":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.EncodedReplyParse(io, cBusOptions, requestContext)
	case "InterfaceOptions3":
		return model.InterfaceOptions3Parse(io)
	case "InterfaceOptions1":
		return model.InterfaceOptions1Parse(io)
	case "InterfaceOptions2":
		return model.InterfaceOptions2Parse(io)
	case "UnitAddress":
		return model.UnitAddressParse(io)
	case "NetworkProtocolControlInformation":
		return model.NetworkProtocolControlInformationParse(io)
	case "ExtendedFormatStatusReply":
		return model.ExtendedFormatStatusReplyParse(io)
	case "CBusHeader":
		return model.CBusHeaderParse(io)
	case "Request":
		var cBusOptions model.CBusOptions
		messageLength, err := utils.StrToUint16(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.RequestParse(io, cBusOptions, messageLength)
	case "CBusPointToPointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToPointCommandParse(io, cBusOptions)
	case "Alpha":
		return model.AlphaParse(io)
	case "CBusPointToPointToMultipointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToPointToMultipointCommandParse(io, cBusOptions)
	case "LogicAssignment":
		return model.LogicAssignmentParse(io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
