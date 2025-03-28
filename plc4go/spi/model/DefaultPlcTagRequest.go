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
	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
)

//go:generate go run ../../tools/plc4xgenerator/gen.go -type=DefaultPlcTagRequest
type DefaultPlcTagRequest struct {
	tags     map[string]apiModel.PlcTag
	tagNames []string `ignore:"true"`
}

func NewDefaultPlcTagRequest(tags map[string]apiModel.PlcTag, tagNames []string) *DefaultPlcTagRequest {
	return &DefaultPlcTagRequest{tags: tags, tagNames: tagNames}
}

func (d *DefaultPlcTagRequest) IsAPlcMessage() bool {
	return true
}

func (d *DefaultPlcTagRequest) GetTagNames() []string {
	return d.tagNames
}

func (d *DefaultPlcTagRequest) GetTag(name string) apiModel.PlcTag {
	tag, ok := d.tags[name]
	if !ok {
		return nil
	}
	return tag
}
