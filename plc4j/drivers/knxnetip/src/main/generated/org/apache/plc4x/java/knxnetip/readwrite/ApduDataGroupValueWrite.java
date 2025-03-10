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
package org.apache.plc4x.java.knxnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ApduDataGroupValueWrite extends ApduData implements Message {

  // Accessors for discriminator values.
  public Byte getApciType() {
    return (byte) 0x2;
  }

  // Properties.
  protected final byte dataFirstByte;
  protected final byte[] data;

  public ApduDataGroupValueWrite(byte dataFirstByte, byte[] data) {
    super();
    this.dataFirstByte = dataFirstByte;
    this.data = data;
  }

  public byte getDataFirstByte() {
    return dataFirstByte;
  }

  public byte[] getData() {
    return data;
  }

  @Override
  protected void serializeApduDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ApduDataGroupValueWrite");

    // Simple Field (dataFirstByte)
    writeSimpleField("dataFirstByte", dataFirstByte, writeSignedByte(writeBuffer, 6));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("ApduDataGroupValueWrite");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ApduDataGroupValueWrite _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (dataFirstByte)
    lengthInBits += 6;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static ApduDataBuilder staticParseApduDataBuilder(ReadBuffer readBuffer, Short dataLength)
      throws ParseException {
    readBuffer.pullContext("ApduDataGroupValueWrite");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte dataFirstByte = readSimpleField("dataFirstByte", readSignedByte(readBuffer, 6));

    byte[] data =
        readBuffer.readByteArray(
            "data", Math.toIntExact(((((dataLength) < (1))) ? 0 : (dataLength) - (1))));

    readBuffer.closeContext("ApduDataGroupValueWrite");
    // Create the instance
    return new ApduDataGroupValueWriteBuilderImpl(dataFirstByte, data);
  }

  public static class ApduDataGroupValueWriteBuilderImpl implements ApduData.ApduDataBuilder {
    private final byte dataFirstByte;
    private final byte[] data;

    public ApduDataGroupValueWriteBuilderImpl(byte dataFirstByte, byte[] data) {
      this.dataFirstByte = dataFirstByte;
      this.data = data;
    }

    public ApduDataGroupValueWrite build() {
      ApduDataGroupValueWrite apduDataGroupValueWrite =
          new ApduDataGroupValueWrite(dataFirstByte, data);
      return apduDataGroupValueWrite;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApduDataGroupValueWrite)) {
      return false;
    }
    ApduDataGroupValueWrite that = (ApduDataGroupValueWrite) o;
    return (getDataFirstByte() == that.getDataFirstByte())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDataFirstByte(), getData());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
