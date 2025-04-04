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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class APDUIFormat extends APDU implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final int receiveSequenceNo;
  protected final ASDU asdu;

  public APDUIFormat(int command, int receiveSequenceNo, ASDU asdu) {
    super(command);
    this.receiveSequenceNo = receiveSequenceNo;
    this.asdu = asdu;
  }

  public int getReceiveSequenceNo() {
    return receiveSequenceNo;
  }

  public ASDU getAsdu() {
    return asdu;
  }

  @Override
  protected void serializeAPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("APDUIFormat");

    // Simple Field (receiveSequenceNo)
    writeSimpleField(
        "receiveSequenceNo",
        receiveSequenceNo,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (asdu)
    writeSimpleField(
        "asdu",
        asdu,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("APDUIFormat");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    APDUIFormat _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (receiveSequenceNo)
    lengthInBits += 16;

    // Simple field (asdu)
    lengthInBits += asdu.getLengthInBits();

    return lengthInBits;
  }

  public static APDUBuilder staticParseAPDUBuilder(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("APDUIFormat");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int receiveSequenceNo =
        readSimpleField(
            "receiveSequenceNo",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    ASDU asdu =
        readSimpleField(
            "asdu",
            new DataReaderComplexDefault<>(() -> ASDU.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("APDUIFormat");
    // Create the instance
    return new APDUIFormatBuilderImpl(receiveSequenceNo, asdu);
  }

  public static class APDUIFormatBuilderImpl implements APDU.APDUBuilder {
    private final int receiveSequenceNo;
    private final ASDU asdu;

    public APDUIFormatBuilderImpl(int receiveSequenceNo, ASDU asdu) {
      this.receiveSequenceNo = receiveSequenceNo;
      this.asdu = asdu;
    }

    public APDUIFormat build(int command) {
      APDUIFormat aPDUIFormat = new APDUIFormat(command, receiveSequenceNo, asdu);
      return aPDUIFormat;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof APDUIFormat)) {
      return false;
    }
    APDUIFormat that = (APDUIFormat) o;
    return (getReceiveSequenceNo() == that.getReceiveSequenceNo())
        && (getAsdu() == that.getAsdu())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getReceiveSequenceNo(), getAsdu());
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
