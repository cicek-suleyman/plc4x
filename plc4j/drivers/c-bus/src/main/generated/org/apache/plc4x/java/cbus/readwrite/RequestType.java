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
package org.apache.plc4x.java.cbus.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum RequestType {
  UNKNOWN((short) 0x00, (short) 0x00),
  SMART_CONNECT_SHORTCUT((short) 0x7C, (short) 0x7C),
  RESET((short) 0x7E, (short) 0x7E),
  DIRECT_COMMAND((short) 0x40, (short) 0x40),
  REQUEST_COMMAND((short) 0x5C, (short) 0x5C),
  NULL((short) 0x6E, (short) 0x00),
  EMPTY((short) 0x0D, (short) 0x00);
  private static final Map<Short, RequestType> map;

  static {
    map = new HashMap<>();
    for (RequestType value : RequestType.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final short controlChar;

  RequestType(short value, short controlChar) {
    this.value = value;
    this.controlChar = controlChar;
  }

  public short getValue() {
    return value;
  }

  public short getControlChar() {
    return controlChar;
  }

  public static RequestType firstEnumForFieldControlChar(short fieldValue) {
    for (RequestType _val : RequestType.values()) {
      if (_val.getControlChar() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<RequestType> enumsForFieldControlChar(short fieldValue) {
    List<RequestType> _values = new ArrayList<>();
    for (RequestType _val : RequestType.values()) {
      if (_val.getControlChar() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static RequestType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
