<#--
<#--
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
-->
<#-- Prevent freemarker from escaping stuff -->
<#outputformat "undefined">
<#-- Declare the name and type of variables passed in to the template -->
<#-- @ftlvariable name="languageName" type="java.lang.String" -->
<#-- @ftlvariable name="protocolName" type="java.lang.String" -->
<#-- @ftlvariable name="outputFlavor" type="java.lang.String" -->
<#-- @ftlvariable name="helper" type="org.apache.plc4x.language.java.JavaLanguageTemplateHelper" -->
<#-- @ftlvariable name="tracer" type="org.apache.plc4x.plugins.codegenerator.protocol.freemarker.Tracer" -->
<#-- @ftlvariable name="type" type="org.apache.plc4x.plugins.codegenerator.types.definitions.EnumTypeDefinition" -->
${helper.packageName(protocolName, languageName, outputFlavor)?replace(".", "/")}/${type.name}.py
#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
<#macro importSectionWithContentBelow>
    <#local sectionContent><#nested></#local>
    <#if helper.getRequiredImports()?has_content>
        <#list helper.getRequiredImports() as import>
${import}
        </#list>
    </#if>
    ${sectionContent}
</#macro>
<#macro emitImport import>${helper.emitRequiredImport(import)}</#macro>
<@importSectionWithContentBelow>
<@emitImport import="from enum import IntEnum" />

class ${type.name}(IntEnum):
    <#list type.enumValues as enumValue>
    ${enumValue.name}: <@compress single_line=true>
            <#if type.type.isPresent()>
                ${helper.getLanguageTypeNameForTypeReference(type.type.orElseThrow(), true)}
                <#if type.type.orElseThrow().isNonSimpleTypeReference()>
                    <#if type.type.orElseThrow().isEnumTypeReference()>
                         = <#if type.constantNames?has_content>(</#if>${helper.getLanguageTypeNameForTypeReference(type.type.orElseThrow(), true)}.${enumValue.value}
                    <#else>
                         = <#if type.constantNames?has_content>(</#if>${enumValue.value}
                    </#if>
                <#else>
                    = <#if type.constantNames?has_content>(</#if>${enumValue.value}
                </#if>
            </#if>
            <#if type.constantNames?has_content>
                <#if type.type?has_content>, </#if>
                <#list type.constantNames as constantName>
                    <#if type.getConstantType(constantName).isNonSimpleTypeReference()>
                        <#if helper.escapeValue(type.getConstantType(constantName), enumValue.getConstant(constantName).orElse(null)) == 'null'>
                            = None
                        <#elseif type.getConstantType(constantName).isEnumTypeReference()>
                            = ${helper.getLanguageTypeNameForTypeReference(type.getConstantType(constantName), true)}.${helper.escapeValue(type.getConstantType(constantName), enumValue.getConstant(constantName).orElseThrow())}<#else>${helper.getLanguageTypeNameForTypeReference(type.getConstantType(constantName), true)}(${helper.escapeValue(type.getConstantType(constantName), enumValue.getConstant(constantName).orElseThrow())})</#if><#else>${helper.getLanguageTypeNameForTypeReference(type.getConstantType(constantName), true)}(${helper.escapeValue(type.getConstantType(constantName), enumValue.getConstant(constantName).orElseThrow())})
                        </#if>
                    <#sep>, </#sep>
                </#list>
                )
            </#if>
    </@compress>

    </#list>

    <#if type.constantNames?has_content>
    def __new__(cls, value, <@compress single_line=true>
        <#list type.constantNames as constantName>
            ${helper.camelCaseToSnakeCase(constantName)}
            <#sep>, </#sep>
        </#list>):
        </@compress>

        obj = object.__new__(cls)
        obj._value_ = value
        <#list type.constantNames as constantName>obj.${constantName} = ${constantName}</#list>
        return obj
    </#if>

</@importSectionWithContentBelow>

</#outputformat>
