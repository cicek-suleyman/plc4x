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

// CallMethodResult is the corresponding interface of CallMethodResult
type CallMethodResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfInputArgumentResults returns NoOfInputArgumentResults (property field)
	GetNoOfInputArgumentResults() int32
	// GetInputArgumentResults returns InputArgumentResults (property field)
	GetInputArgumentResults() []StatusCode
	// GetNoOfInputArgumentDiagnosticInfos returns NoOfInputArgumentDiagnosticInfos (property field)
	GetNoOfInputArgumentDiagnosticInfos() int32
	// GetInputArgumentDiagnosticInfos returns InputArgumentDiagnosticInfos (property field)
	GetInputArgumentDiagnosticInfos() []DiagnosticInfo
	// GetNoOfOutputArguments returns NoOfOutputArguments (property field)
	GetNoOfOutputArguments() int32
	// GetOutputArguments returns OutputArguments (property field)
	GetOutputArguments() []Variant
}

// CallMethodResultExactly can be used when we want exactly this type and not a type which fulfills CallMethodResult.
// This is useful for switch cases.
type CallMethodResultExactly interface {
	CallMethodResult
	isCallMethodResult() bool
}

// _CallMethodResult is the data-structure of this message
type _CallMethodResult struct {
	*_ExtensionObjectDefinition
	StatusCode                       StatusCode
	NoOfInputArgumentResults         int32
	InputArgumentResults             []StatusCode
	NoOfInputArgumentDiagnosticInfos int32
	InputArgumentDiagnosticInfos     []DiagnosticInfo
	NoOfOutputArguments              int32
	OutputArguments                  []Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CallMethodResult) GetIdentifier() string {
	return "709"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CallMethodResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CallMethodResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CallMethodResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_CallMethodResult) GetNoOfInputArgumentResults() int32 {
	return m.NoOfInputArgumentResults
}

func (m *_CallMethodResult) GetInputArgumentResults() []StatusCode {
	return m.InputArgumentResults
}

func (m *_CallMethodResult) GetNoOfInputArgumentDiagnosticInfos() int32 {
	return m.NoOfInputArgumentDiagnosticInfos
}

func (m *_CallMethodResult) GetInputArgumentDiagnosticInfos() []DiagnosticInfo {
	return m.InputArgumentDiagnosticInfos
}

func (m *_CallMethodResult) GetNoOfOutputArguments() int32 {
	return m.NoOfOutputArguments
}

func (m *_CallMethodResult) GetOutputArguments() []Variant {
	return m.OutputArguments
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCallMethodResult factory function for _CallMethodResult
func NewCallMethodResult(statusCode StatusCode, noOfInputArgumentResults int32, inputArgumentResults []StatusCode, noOfInputArgumentDiagnosticInfos int32, inputArgumentDiagnosticInfos []DiagnosticInfo, noOfOutputArguments int32, outputArguments []Variant) *_CallMethodResult {
	_result := &_CallMethodResult{
		StatusCode:                       statusCode,
		NoOfInputArgumentResults:         noOfInputArgumentResults,
		InputArgumentResults:             inputArgumentResults,
		NoOfInputArgumentDiagnosticInfos: noOfInputArgumentDiagnosticInfos,
		InputArgumentDiagnosticInfos:     inputArgumentDiagnosticInfos,
		NoOfOutputArguments:              noOfOutputArguments,
		OutputArguments:                  outputArguments,
		_ExtensionObjectDefinition:       NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCallMethodResult(structType any) CallMethodResult {
	if casted, ok := structType.(CallMethodResult); ok {
		return casted
	}
	if casted, ok := structType.(*CallMethodResult); ok {
		return *casted
	}
	return nil
}

func (m *_CallMethodResult) GetTypeName() string {
	return "CallMethodResult"
}

func (m *_CallMethodResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfInputArgumentResults)
	lengthInBits += 32

	// Array field
	if len(m.InputArgumentResults) > 0 {
		for _curItem, element := range m.InputArgumentResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.InputArgumentResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfInputArgumentDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.InputArgumentDiagnosticInfos) > 0 {
		for _curItem, element := range m.InputArgumentDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.InputArgumentDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfOutputArguments)
	lengthInBits += 32

	// Array field
	if len(m.OutputArguments) > 0 {
		for _curItem, element := range m.OutputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.OutputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CallMethodResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CallMethodResultParse(ctx context.Context, theBytes []byte, identifier string) (CallMethodResult, error) {
	return CallMethodResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CallMethodResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CallMethodResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CallMethodResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CallMethodResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (statusCode)
	if pullErr := readBuffer.PullContext("statusCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusCode")
	}
	_statusCode, _statusCodeErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _statusCodeErr != nil {
		return nil, errors.Wrap(_statusCodeErr, "Error parsing 'statusCode' field of CallMethodResult")
	}
	statusCode := _statusCode.(StatusCode)
	if closeErr := readBuffer.CloseContext("statusCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusCode")
	}

	// Simple Field (noOfInputArgumentResults)
	_noOfInputArgumentResults, _noOfInputArgumentResultsErr := readBuffer.ReadInt32("noOfInputArgumentResults", 32)
	if _noOfInputArgumentResultsErr != nil {
		return nil, errors.Wrap(_noOfInputArgumentResultsErr, "Error parsing 'noOfInputArgumentResults' field of CallMethodResult")
	}
	noOfInputArgumentResults := _noOfInputArgumentResults

	// Array field (inputArgumentResults)
	if pullErr := readBuffer.PullContext("inputArgumentResults", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inputArgumentResults")
	}
	// Count array
	inputArgumentResults := make([]StatusCode, utils.Max(noOfInputArgumentResults, 0))
	// This happens when the size is set conditional to 0
	if len(inputArgumentResults) == 0 {
		inputArgumentResults = nil
	}
	{
		_numItems := uint16(utils.Max(noOfInputArgumentResults, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := StatusCodeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'inputArgumentResults' field of CallMethodResult")
			}
			inputArgumentResults[_curItem] = _item.(StatusCode)
		}
	}
	if closeErr := readBuffer.CloseContext("inputArgumentResults", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inputArgumentResults")
	}

	// Simple Field (noOfInputArgumentDiagnosticInfos)
	_noOfInputArgumentDiagnosticInfos, _noOfInputArgumentDiagnosticInfosErr := readBuffer.ReadInt32("noOfInputArgumentDiagnosticInfos", 32)
	if _noOfInputArgumentDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfInputArgumentDiagnosticInfosErr, "Error parsing 'noOfInputArgumentDiagnosticInfos' field of CallMethodResult")
	}
	noOfInputArgumentDiagnosticInfos := _noOfInputArgumentDiagnosticInfos

	// Array field (inputArgumentDiagnosticInfos)
	if pullErr := readBuffer.PullContext("inputArgumentDiagnosticInfos", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inputArgumentDiagnosticInfos")
	}
	// Count array
	inputArgumentDiagnosticInfos := make([]DiagnosticInfo, utils.Max(noOfInputArgumentDiagnosticInfos, 0))
	// This happens when the size is set conditional to 0
	if len(inputArgumentDiagnosticInfos) == 0 {
		inputArgumentDiagnosticInfos = nil
	}
	{
		_numItems := uint16(utils.Max(noOfInputArgumentDiagnosticInfos, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := DiagnosticInfoParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'inputArgumentDiagnosticInfos' field of CallMethodResult")
			}
			inputArgumentDiagnosticInfos[_curItem] = _item.(DiagnosticInfo)
		}
	}
	if closeErr := readBuffer.CloseContext("inputArgumentDiagnosticInfos", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inputArgumentDiagnosticInfos")
	}

	// Simple Field (noOfOutputArguments)
	_noOfOutputArguments, _noOfOutputArgumentsErr := readBuffer.ReadInt32("noOfOutputArguments", 32)
	if _noOfOutputArgumentsErr != nil {
		return nil, errors.Wrap(_noOfOutputArgumentsErr, "Error parsing 'noOfOutputArguments' field of CallMethodResult")
	}
	noOfOutputArguments := _noOfOutputArguments

	// Array field (outputArguments)
	if pullErr := readBuffer.PullContext("outputArguments", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for outputArguments")
	}
	// Count array
	outputArguments := make([]Variant, utils.Max(noOfOutputArguments, 0))
	// This happens when the size is set conditional to 0
	if len(outputArguments) == 0 {
		outputArguments = nil
	}
	{
		_numItems := uint16(utils.Max(noOfOutputArguments, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := VariantParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'outputArguments' field of CallMethodResult")
			}
			outputArguments[_curItem] = _item.(Variant)
		}
	}
	if closeErr := readBuffer.CloseContext("outputArguments", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for outputArguments")
	}

	if closeErr := readBuffer.CloseContext("CallMethodResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CallMethodResult")
	}

	// Create a partially initialized instance
	_child := &_CallMethodResult{
		_ExtensionObjectDefinition:       &_ExtensionObjectDefinition{},
		StatusCode:                       statusCode,
		NoOfInputArgumentResults:         noOfInputArgumentResults,
		InputArgumentResults:             inputArgumentResults,
		NoOfInputArgumentDiagnosticInfos: noOfInputArgumentDiagnosticInfos,
		InputArgumentDiagnosticInfos:     inputArgumentDiagnosticInfos,
		NoOfOutputArguments:              noOfOutputArguments,
		OutputArguments:                  outputArguments,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CallMethodResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CallMethodResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CallMethodResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CallMethodResult")
		}

		// Simple Field (statusCode)
		if pushErr := writeBuffer.PushContext("statusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusCode")
		}
		_statusCodeErr := writeBuffer.WriteSerializable(ctx, m.GetStatusCode())
		if popErr := writeBuffer.PopContext("statusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusCode")
		}
		if _statusCodeErr != nil {
			return errors.Wrap(_statusCodeErr, "Error serializing 'statusCode' field")
		}

		// Simple Field (noOfInputArgumentResults)
		noOfInputArgumentResults := int32(m.GetNoOfInputArgumentResults())
		_noOfInputArgumentResultsErr := writeBuffer.WriteInt32("noOfInputArgumentResults", 32, (noOfInputArgumentResults))
		if _noOfInputArgumentResultsErr != nil {
			return errors.Wrap(_noOfInputArgumentResultsErr, "Error serializing 'noOfInputArgumentResults' field")
		}

		// Array Field (inputArgumentResults)
		if pushErr := writeBuffer.PushContext("inputArgumentResults", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inputArgumentResults")
		}
		for _curItem, _element := range m.GetInputArgumentResults() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetInputArgumentResults()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'inputArgumentResults' field")
			}
		}
		if popErr := writeBuffer.PopContext("inputArgumentResults", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inputArgumentResults")
		}

		// Simple Field (noOfInputArgumentDiagnosticInfos)
		noOfInputArgumentDiagnosticInfos := int32(m.GetNoOfInputArgumentDiagnosticInfos())
		_noOfInputArgumentDiagnosticInfosErr := writeBuffer.WriteInt32("noOfInputArgumentDiagnosticInfos", 32, (noOfInputArgumentDiagnosticInfos))
		if _noOfInputArgumentDiagnosticInfosErr != nil {
			return errors.Wrap(_noOfInputArgumentDiagnosticInfosErr, "Error serializing 'noOfInputArgumentDiagnosticInfos' field")
		}

		// Array Field (inputArgumentDiagnosticInfos)
		if pushErr := writeBuffer.PushContext("inputArgumentDiagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inputArgumentDiagnosticInfos")
		}
		for _curItem, _element := range m.GetInputArgumentDiagnosticInfos() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetInputArgumentDiagnosticInfos()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'inputArgumentDiagnosticInfos' field")
			}
		}
		if popErr := writeBuffer.PopContext("inputArgumentDiagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inputArgumentDiagnosticInfos")
		}

		// Simple Field (noOfOutputArguments)
		noOfOutputArguments := int32(m.GetNoOfOutputArguments())
		_noOfOutputArgumentsErr := writeBuffer.WriteInt32("noOfOutputArguments", 32, (noOfOutputArguments))
		if _noOfOutputArgumentsErr != nil {
			return errors.Wrap(_noOfOutputArgumentsErr, "Error serializing 'noOfOutputArguments' field")
		}

		// Array Field (outputArguments)
		if pushErr := writeBuffer.PushContext("outputArguments", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for outputArguments")
		}
		for _curItem, _element := range m.GetOutputArguments() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetOutputArguments()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'outputArguments' field")
			}
		}
		if popErr := writeBuffer.PopContext("outputArguments", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for outputArguments")
		}

		if popErr := writeBuffer.PopContext("CallMethodResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CallMethodResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CallMethodResult) isCallMethodResult() bool {
	return true
}

func (m *_CallMethodResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
