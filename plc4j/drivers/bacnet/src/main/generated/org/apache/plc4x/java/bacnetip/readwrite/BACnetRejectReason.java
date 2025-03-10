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

public enum BACnetRejectReason {
  OTHER((short) 0x0),
  BUFFER_OVERFLOW((short) 0x1),
  INCONSISTENT_PARAMETERS((short) 0x2),
  INVALID_PARAMETER_DATA_TYPE((short) 0x3),
  INVALID_TAG((short) 0x4),
  MISSING_REQUIRED_PARAMETER((short) 0x5),
  PARAMETER_OUT_OF_RANGE((short) 0x6),
  TOO_MANY_ARGUMENTS((short) 0x7),
  UNDEFINED_ENUMERATION((short) 0x8),
  UNRECOGNIZED_SERVICE((short) 0x9),
  VENDOR_PROPRIETARY_VALUE((short) 0xFF);
  private static final Map<Short, BACnetRejectReason> map;

  static {
    map = new HashMap<>();
    for (BACnetRejectReason value : BACnetRejectReason.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;

  BACnetRejectReason(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static BACnetRejectReason enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
