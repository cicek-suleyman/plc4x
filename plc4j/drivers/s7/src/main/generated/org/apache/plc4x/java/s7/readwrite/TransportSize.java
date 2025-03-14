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
package org.apache.plc4x.java.s7.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum TransportSize {
  BOOL(
      (short) 0x01,
      (boolean) true,
      (boolean) true,
      (short) 0x01,
      (short) 1,
      (boolean) true,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      DataTransportSize.BIT,
      (String) "IEC61131_BOOL",
      null),
  BYTE(
      (short) 0x02,
      (boolean) true,
      (boolean) true,
      (short) 0x02,
      (short) 1,
      (boolean) true,
      (boolean) true,
      (short) 'B',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_BYTE",
      null),
  WORD(
      (short) 0x03,
      (boolean) true,
      (boolean) true,
      (short) 0x04,
      (short) 2,
      (boolean) true,
      (boolean) true,
      (short) 'W',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_WORD",
      null),
  DWORD(
      (short) 0x04,
      (boolean) true,
      (boolean) true,
      (short) 0x06,
      (short) 4,
      (boolean) true,
      (boolean) true,
      (short) 'D',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_DWORD",
      TransportSize.WORD),
  LWORD(
      (short) 0x05,
      (boolean) false,
      (boolean) false,
      (short) 0x00,
      (short) 8,
      (boolean) false,
      (boolean) false,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_LWORD",
      null),
  INT(
      (short) 0x06,
      (boolean) true,
      (boolean) true,
      (short) 0x05,
      (short) 2,
      (boolean) true,
      (boolean) true,
      (short) 'W',
      (boolean) true,
      DataTransportSize.INTEGER,
      (String) "IEC61131_INT",
      null),
  UINT(
      (short) 0x07,
      (boolean) false,
      (boolean) true,
      (short) 0x05,
      (short) 2,
      (boolean) false,
      (boolean) true,
      (short) 'W',
      (boolean) true,
      DataTransportSize.INTEGER,
      (String) "IEC61131_UINT",
      TransportSize.INT),
  SINT(
      (short) 0x08,
      (boolean) false,
      (boolean) true,
      (short) 0x02,
      (short) 1,
      (boolean) false,
      (boolean) true,
      (short) 'B',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_SINT",
      TransportSize.INT),
  USINT(
      (short) 0x09,
      (boolean) false,
      (boolean) true,
      (short) 0x02,
      (short) 1,
      (boolean) false,
      (boolean) true,
      (short) 'B',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_USINT",
      TransportSize.INT),
  DINT(
      (short) 0x0A,
      (boolean) true,
      (boolean) true,
      (short) 0x07,
      (short) 4,
      (boolean) true,
      (boolean) true,
      (short) 'D',
      (boolean) true,
      DataTransportSize.INTEGER,
      (String) "IEC61131_DINT",
      TransportSize.INT),
  UDINT(
      (short) 0x0B,
      (boolean) false,
      (boolean) true,
      (short) 0x07,
      (short) 4,
      (boolean) false,
      (boolean) true,
      (short) 'D',
      (boolean) true,
      DataTransportSize.INTEGER,
      (String) "IEC61131_UDINT",
      TransportSize.INT),
  LINT(
      (short) 0x0C,
      (boolean) false,
      (boolean) false,
      (short) 0x00,
      (short) 8,
      (boolean) false,
      (boolean) false,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_LINT",
      TransportSize.INT),
  ULINT(
      (short) 0x0D,
      (boolean) false,
      (boolean) false,
      (short) 0x00,
      (short) 16,
      (boolean) false,
      (boolean) false,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_ULINT",
      TransportSize.INT),
  REAL(
      (short) 0x0E,
      (boolean) true,
      (boolean) true,
      (short) 0x08,
      (short) 4,
      (boolean) true,
      (boolean) true,
      (short) 'D',
      (boolean) true,
      DataTransportSize.REAL,
      (String) "IEC61131_REAL",
      null),
  LREAL(
      (short) 0x0F,
      (boolean) false,
      (boolean) false,
      (short) 0x30,
      (short) 8,
      (boolean) false,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_LREAL",
      TransportSize.REAL),
  CHAR(
      (short) 0x10,
      (boolean) true,
      (boolean) true,
      (short) 0x03,
      (short) 1,
      (boolean) true,
      (boolean) true,
      (short) 'B',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_CHAR",
      null),
  WCHAR(
      (short) 0x11,
      (boolean) false,
      (boolean) true,
      (short) 0x13,
      (short) 2,
      (boolean) false,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_WCHAR",
      null),
  STRING(
      (short) 0x12,
      (boolean) true,
      (boolean) true,
      (short) 0x03,
      (short) 1,
      (boolean) true,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_STRING",
      null),
  WSTRING(
      (short) 0x13,
      (boolean) false,
      (boolean) true,
      (short) 0x00,
      (short) 2,
      (boolean) false,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_WSTRING",
      null),
  TIME(
      (short) 0x14,
      (boolean) true,
      (boolean) true,
      (short) 0x0B,
      (short) 4,
      (boolean) true,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_TIME",
      null),
  LTIME(
      (short) 0x16,
      (boolean) false,
      (boolean) false,
      (short) 0x00,
      (short) 8,
      (boolean) false,
      (boolean) false,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_LTIME",
      TransportSize.TIME),
  DATE(
      (short) 0x17,
      (boolean) true,
      (boolean) true,
      (short) 0x09,
      (short) 2,
      (boolean) true,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_DATE",
      null),
  TIME_OF_DAY(
      (short) 0x18,
      (boolean) true,
      (boolean) true,
      (short) 0x06,
      (short) 4,
      (boolean) true,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_TIME_OF_DAY",
      null),
  TOD(
      (short) 0x19,
      (boolean) true,
      (boolean) true,
      (short) 0x06,
      (short) 4,
      (boolean) true,
      (boolean) true,
      (short) 'X',
      (boolean) true,
      DataTransportSize.BYTE_WORD_DWORD,
      (String) "IEC61131_TIME_OF_DAY",
      null),
  DATE_AND_TIME(
      (short) 0x1A,
      (boolean) true,
      (boolean) false,
      (short) 0x0F,
      (short) 12,
      (boolean) true,
      (boolean) false,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_DATE_AND_TIME",
      null),
  DT(
      (short) 0x1B,
      (boolean) true,
      (boolean) false,
      (short) 0x0F,
      (short) 12,
      (boolean) true,
      (boolean) false,
      (short) 'X',
      (boolean) true,
      null,
      (String) "IEC61131_DATE_AND_TIME",
      null);
  private static final Map<Short, TransportSize> map;

  static {
    map = new HashMap<>();
    for (TransportSize value : TransportSize.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final boolean supported_S7_300;
  private final boolean supported_LOGO;
  private final short code;
  private final short sizeInBytes;
  private final boolean supported_S7_400;
  private final boolean supported_S7_1200;
  private final short shortName;
  private final boolean supported_S7_1500;
  private final DataTransportSize dataTransportSize;
  private final String dataProtocolId;
  private final TransportSize baseType;

  TransportSize(
      short value,
      boolean supported_S7_300,
      boolean supported_LOGO,
      short code,
      short sizeInBytes,
      boolean supported_S7_400,
      boolean supported_S7_1200,
      short shortName,
      boolean supported_S7_1500,
      DataTransportSize dataTransportSize,
      String dataProtocolId,
      TransportSize baseType) {
    this.value = value;
    this.supported_S7_300 = supported_S7_300;
    this.supported_LOGO = supported_LOGO;
    this.code = code;
    this.sizeInBytes = sizeInBytes;
    this.supported_S7_400 = supported_S7_400;
    this.supported_S7_1200 = supported_S7_1200;
    this.shortName = shortName;
    this.supported_S7_1500 = supported_S7_1500;
    this.dataTransportSize = dataTransportSize;
    this.dataProtocolId = dataProtocolId;
    this.baseType = baseType;
  }

  public short getValue() {
    return value;
  }

  public boolean getSupported_S7_300() {
    return supported_S7_300;
  }

  public static TransportSize firstEnumForFieldSupported_S7_300(boolean fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_S7_300() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldSupported_S7_300(boolean fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_S7_300() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public boolean getSupported_LOGO() {
    return supported_LOGO;
  }

  public static TransportSize firstEnumForFieldSupported_LOGO(boolean fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_LOGO() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldSupported_LOGO(boolean fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_LOGO() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public short getCode() {
    return code;
  }

  public static TransportSize firstEnumForFieldCode(short fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getCode() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldCode(short fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getCode() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public short getSizeInBytes() {
    return sizeInBytes;
  }

  public static TransportSize firstEnumForFieldSizeInBytes(short fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSizeInBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldSizeInBytes(short fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSizeInBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public boolean getSupported_S7_400() {
    return supported_S7_400;
  }

  public static TransportSize firstEnumForFieldSupported_S7_400(boolean fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_S7_400() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldSupported_S7_400(boolean fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_S7_400() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public boolean getSupported_S7_1200() {
    return supported_S7_1200;
  }

  public static TransportSize firstEnumForFieldSupported_S7_1200(boolean fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_S7_1200() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldSupported_S7_1200(boolean fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_S7_1200() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public short getShortName() {
    return shortName;
  }

  public static TransportSize firstEnumForFieldShortName(short fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getShortName() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldShortName(short fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getShortName() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public boolean getSupported_S7_1500() {
    return supported_S7_1500;
  }

  public static TransportSize firstEnumForFieldSupported_S7_1500(boolean fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_S7_1500() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldSupported_S7_1500(boolean fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getSupported_S7_1500() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public DataTransportSize getDataTransportSize() {
    return dataTransportSize;
  }

  public static TransportSize firstEnumForFieldDataTransportSize(DataTransportSize fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getDataTransportSize() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldDataTransportSize(DataTransportSize fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getDataTransportSize() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public String getDataProtocolId() {
    return dataProtocolId;
  }

  public static TransportSize firstEnumForFieldDataProtocolId(String fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getDataProtocolId().equals(fieldValue)) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldDataProtocolId(String fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getDataProtocolId().equals(fieldValue)) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public TransportSize getBaseType() {
    return baseType;
  }

  public static TransportSize firstEnumForFieldBaseType(TransportSize fieldValue) {
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getBaseType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TransportSize> enumsForFieldBaseType(TransportSize fieldValue) {
    List<TransportSize> _values = new ArrayList<>();
    for (TransportSize _val : TransportSize.values()) {
      if (_val.getBaseType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static TransportSize enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
