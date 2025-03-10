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

package ads

import (
	"context"
	"encoding/binary"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/apache/plc4x/plc4go/internal/ads/model"
	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
	apiValues "github.com/apache/plc4x/plc4go/pkg/api/values"
	driverModel "github.com/apache/plc4x/plc4go/protocols/ads/readwrite/model"
	spiModel "github.com/apache/plc4x/plc4go/spi/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	spiValues "github.com/apache/plc4x/plc4go/spi/values"

	"github.com/pkg/errors"
)

func (m *Connection) ReadRequestBuilder() apiModel.PlcReadRequestBuilder {
	return spiModel.NewDefaultPlcReadRequestBuilder(m.GetPlcTagHandler(), m)
}

func (m *Connection) Read(ctx context.Context, readRequest apiModel.PlcReadRequest) <-chan apiModel.PlcReadRequestResult {
	m.log.Trace().Msg("Reading")
	result := make(chan apiModel.PlcReadRequestResult, 1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				result <- spiModel.NewDefaultPlcReadRequestResult(readRequest, nil, errors.Errorf("panic-ed %v. Stack: %s", err, debug.Stack()))
			}
		}()
		if len(readRequest.GetTagNames()) <= 1 {
			m.singleRead(ctx, readRequest, result)
		} else {
			m.multiRead(ctx, readRequest, result)
		}
	}()
	return result
}

func (m *Connection) singleRead(ctx context.Context, readRequest apiModel.PlcReadRequest, result chan apiModel.PlcReadRequestResult) {
	if len(readRequest.GetTagNames()) != 1 {
		result <- spiModel.NewDefaultPlcReadRequestResult(readRequest, nil, errors.New("this part of the ads driver only supports single-item requests"))
		m.log.Debug().Int("nTags", len(readRequest.GetTagNames())).Msg("this part of the ads driver only supports single-item requests. Got nTags tags")
		return
	}

	// Here we can be sure that we're only handling a single request.
	tagName := readRequest.GetTagNames()[0]
	tag := readRequest.GetTag(tagName)
	if model.NeedsResolving(tag) {
		adsField, err := model.CastToSymbolicPlcTagFromPlcTag(tag)
		if err != nil {
			result <- spiModel.NewDefaultPlcReadRequestResult(readRequest, nil, errors.Wrap(err, "invalid tag item type"))
			m.log.Debug().Type("tag", tag).Msg("Invalid tag item type")
			return
		}
		// Replace the symbolic tag with a direct one
		tag, err = m.resolveSymbolicTag(ctx, adsField)
		if err != nil {
			result <- spiModel.NewDefaultPlcReadRequestResult(
				readRequest,
				nil,
				errors.Wrap(err, "invalid tag item type"),
			)
			m.log.Debug().Type("tag", tag).Msg("Invalid tag item type")
			return
		}
	}
	directAdsTag, ok := tag.(*model.DirectPlcTag)
	if !ok {
		result <- spiModel.NewDefaultPlcReadRequestResult(readRequest, nil, errors.New("invalid tag item type"))
		m.log.Debug().Type("tag", tag).Msg("Invalid tag item type")
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				result <- spiModel.NewDefaultPlcReadRequestResult(readRequest, nil, errors.Errorf("panic-ed %v. Stack: %s", err, debug.Stack()))
			}
		}()
		response, err := m.ExecuteAdsReadRequest(ctx, directAdsTag.IndexGroup, directAdsTag.IndexOffset, directAdsTag.DataType.GetSize())
		if err != nil {
			result <- spiModel.NewDefaultPlcReadRequestResult(
				readRequest,
				nil,
				errors.Wrap(err, "got error executing the read request"),
			)
			return
		}

		if response.GetErrorCode() != 0x00000000 {
			// TODO: Handle this ...
		}

		rb := utils.NewReadBufferByteBased(response.GetData(), utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian))
		responseCodes := map[string]apiModel.PlcResponseCode{}
		plcValues := map[string]apiValues.PlcValue{}
		for _, tagName := range readRequest.GetTagNames() {
			m.log.Debug().Str("tagName", tagName).Msg("get a tag from request with name")
			// Try to parse the value
			plcValue, err := m.parsePlcValue(directAdsTag.DataType, directAdsTag.DataType.GetArrayInfo(), rb)
			if err != nil {
				m.log.Error().Err(err).Msg("Error parsing plc value")
				responseCodes[tagName] = apiModel.PlcResponseCode_INTERNAL_ERROR
			} else {
				plcValues[tagName] = plcValue
				responseCodes[tagName] = apiModel.PlcResponseCode_OK
			}
		}
		// Return the response to the caller.
		result <- spiModel.NewDefaultPlcReadRequestResult(
			readRequest,
			spiModel.NewDefaultPlcReadResponse(readRequest, responseCodes, plcValues),
			nil,
		)
	}()
}

func (m *Connection) multiRead(ctx context.Context, readRequest apiModel.PlcReadRequest, result chan apiModel.PlcReadRequestResult) {
	// Calculate the size of all tags together.
	// Calculate the expected size of the response data.
	expectedResponseDataSize := uint32(0)
	directAdsTags := map[string]*model.DirectPlcTag{}
	requestItems := make([]driverModel.AdsMultiRequestItem, 0)
	for _, tagName := range readRequest.GetTagNames() {
		tag := readRequest.GetTag(tagName)
		if model.NeedsResolving(tag) {
			adsField, err := model.CastToSymbolicPlcTagFromPlcTag(tag)
			if err != nil {
				result <- spiModel.NewDefaultPlcReadRequestResult(
					readRequest,
					nil,
					errors.Wrap(err, "invalid tag item type"),
				)
				m.log.Debug().Type("tag", tag).Msg("Invalid tag item type")
				return
			}
			// Replace the symbolic tag with a direct one
			tag, err = m.resolveSymbolicTag(ctx, adsField)
			if err != nil {
				result <- spiModel.NewDefaultPlcReadRequestResult(
					readRequest,
					nil,
					errors.Wrap(err, "invalid tag item type"),
				)
				m.log.Debug().Type("tag", tag).Msg("Invalid tag item type")
				return
			}
		}
		directAdsTag, ok := tag.(*model.DirectPlcTag)
		if !ok {
			result <- spiModel.NewDefaultPlcReadRequestResult(
				readRequest,
				nil,
				errors.New("invalid tag item type"),
			)
			m.log.Debug().Type("tag", tag).Msg("Invalid tag item type")
			return
		}

		directAdsTags[tagName] = directAdsTag

		// Size of one element.
		size := directAdsTag.DataType.GetSize()

		// Calculate how many elements in total we'll be reading.
		arraySize := uint32(1)
		if len(tag.GetArrayInfo()) > 0 {
			for _, arrayInfo := range tag.GetArrayInfo() {
				arraySize = arraySize * arrayInfo.GetSize()
			}
		}

		// Status code + payload size
		expectedTagSize := 4 + (size * arraySize)
		expectedResponseDataSize += expectedTagSize

		requestItems = append(requestItems, driverModel.NewAdsMultiRequestItemRead(directAdsTag.IndexGroup, directAdsTag.IndexOffset, size*arraySize))
	}

	response, err := m.ExecuteAdsReadWriteRequest(ctx, uint32(driverModel.ReservedIndexGroups_ADSIGRP_MULTIPLE_READ), uint32(len(directAdsTags)), expectedResponseDataSize, requestItems, nil)
	if err != nil {
		result <- spiModel.NewDefaultPlcReadRequestResult(
			readRequest,
			nil,
			errors.Wrap(err, "error executing multi-item read request"),
		)
		return
	}

	rb := utils.NewReadBufferByteBased(response.GetData(), utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian))

	// Read in the response codes first.
	responseCodes := map[string]apiModel.PlcResponseCode{}
	plcValues := map[string]apiValues.PlcValue{}
	for _, tagName := range readRequest.GetTagNames() {
		returnCodeValue, err := rb.ReadUint32("returnCode", 32)
		if err != nil {
			responseCodes[tagName] = apiModel.PlcResponseCode_INTERNAL_ERROR
		} else if returnCodeValue != 0x00000000 {
			// TODO: Correctly handle this.
			responseCodes[tagName] = apiModel.PlcResponseCode_REMOTE_ERROR
		} else {
			responseCodes[tagName] = apiModel.PlcResponseCode_OK
		}
	}

	// Parse the plc apiValues for those items that were ok.
	for _, tagName := range readRequest.GetTagNames() {
		if responseCodes[tagName] != apiModel.PlcResponseCode_OK {
			continue
		}

		directAdsTag := directAdsTags[tagName]
		m.log.Debug().Str("tagName", tagName).Msg("get a tag from request with name")
		// Try to parse the value
		plcValue, err := m.parsePlcValue(directAdsTag.DataType, directAdsTag.DataType.GetArrayInfo(), rb)
		if err != nil {
			m.log.Error().Err(err).Msg("Error parsing plc value")
			responseCodes[tagName] = apiModel.PlcResponseCode_INTERNAL_ERROR
		} else {
			plcValues[tagName] = plcValue
			responseCodes[tagName] = apiModel.PlcResponseCode_OK
		}
	}

	// Return the response to the caller.
	result <- spiModel.NewDefaultPlcReadRequestResult(
		readRequest,
		spiModel.NewDefaultPlcReadResponse(readRequest, responseCodes, plcValues),
		nil,
	)
}

func (m *Connection) parsePlcValue(dataType driverModel.AdsDataTypeTableEntry, arrayInfo []driverModel.AdsDataTypeArrayInfo, rb utils.ReadBufferByteBased) (apiValues.PlcValue, error) {
	// Decode the data according to the information from the request
	// Based on the AdsDataTypeTableEntry in tag.DataType() parse the data
	if len(arrayInfo) > 0 {
		// This is an Array/List type.
		curArrayInfo := arrayInfo[0]
		arrayItemTypeName := dataType.GetDataTypeName()[strings.Index(dataType.GetDataTypeName(), " OF ")+4:]
		arrayItemType, ok := m.driverContext.dataTypeTable[arrayItemTypeName]
		if !ok {
			return nil, fmt.Errorf("couldn't resolve array item type %s", arrayItemTypeName)
		}
		var plcValues []apiValues.PlcValue
		for i := uint32(0); i < curArrayInfo.GetNumElements(); i++ {
			restArrayInfo := arrayInfo[1:]
			plcValue, err := m.parsePlcValue(arrayItemType, restArrayInfo, rb)
			if err != nil {
				return nil, errors.Wrap(err, "error decoding list item")
			}
			plcValues = append(plcValues, plcValue)
		}
		return spiValues.NewPlcList(plcValues), nil
	} else if len(dataType.GetChildren()) > 0 {
		// This is a Struct type.
		plcValues := map[string]apiValues.PlcValue{}
		startPos := uint32(rb.GetPos())
		curPos := uint32(0)
		for _, child := range dataType.GetChildren() {
			childName := child.GetPropertyName()
			childDataType, ok := m.driverContext.dataTypeTable[child.GetDataTypeName()]
			if !ok {
				return nil, fmt.Errorf("couldn't find data type named %s for property %s of type %s", child.GetDataTypeName(), childName, dataType.GetDataTypeName())
			}
			if child.GetOffset() > curPos {
				skipBytes := child.GetOffset() - curPos
				for i := uint32(0); i < skipBytes; i++ {
					_, _ = rb.ReadByte("")
				}
			}
			childValue, err := m.parsePlcValue(childDataType, childDataType.GetArrayInfo(), rb)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("error parsing propery %s of type %s", childName, dataType.GetDataTypeName()))
			}
			plcValues[childName] = childValue
			curPos = uint32(rb.GetPos()) - startPos
		}
		return spiValues.NewPlcStruct(plcValues), nil
	} else {
		// This is a primitive type.
		valueType, stringLength := m.getPlcValueForAdsDataTypeTableEntry(dataType)
		if valueType == apiValues.NULL {
			return nil, errors.New(fmt.Sprintf("error converting %s into plc4x plc-value type", dataType.GetDataTypeName()))
		}
		adsValueType, ok := driverModel.PlcValueTypeByName(valueType.String())
		if !ok {
			return nil, errors.New(fmt.Sprintf("error converting plc4x plc-value type %s into ads plc-value type", valueType.String()))
		}
		return driverModel.DataItemParseWithBuffer(context.Background(), rb, adsValueType, stringLength)
	}
}
