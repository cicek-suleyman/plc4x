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
package org.apache.plc4x.java.bacnetip.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum BACnetNodeType {
  UNKNOWN((short) 0x00),
  SYSTEM((short) 0x01),
  NETWORK((short) 0x02),
  DEVICE((short) 0x03),
  ORGANIZATIONAL((short) 0x04),
  AREA((short) 0x05),
  EQUIPMENT((short) 0x06),
  POINT((short) 0x07),
  COLLECTION((short) 0x08),
  PROPERTY((short) 0x09),
  FUNCTIONAL((short) 0x0A),
  OTHER((short) 0x0B),
  SUBSYSTEM((short) 0x0C),
  BUILDING((short) 0x0D),
  FLOOR((short) 0x0E),
  SECTION((short) 0x0F),
  MODULE((short) 0x10),
  TREE((short) 0x11),
  MEMBER((short) 0x12),
  PROTOCOL((short) 0x13),
  ROOM((short) 0x14),
  ZONE((short) 0x15);
  private static final Map<Short, BACnetNodeType> map;

  static {
    map = new HashMap<>();
    for (BACnetNodeType value : BACnetNodeType.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;

  BACnetNodeType(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static BACnetNodeType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
