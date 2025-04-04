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
package org.apache.plc4x.java.firmata.readwrite;

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

public abstract class FirmataMessage implements Message {

  // Abstract accessors for discriminator values.
  public abstract Byte getMessageType();

  public FirmataMessage() {
    super();
  }

  protected abstract void serializeFirmataMessageChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("FirmataMessage");

    // Discriminator Field (messageType) (Used as input to a switch field)
    writeDiscriminatorField(
        "messageType",
        getMessageType(),
        writeUnsignedByte(writeBuffer, 4),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch field (Serialize the sub-type)
    serializeFirmataMessageChild(writeBuffer);

    writeBuffer.popContext("FirmataMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    FirmataMessage _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (messageType)
    lengthInBits += 4;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static FirmataMessage staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Boolean response;
    if (args[0] instanceof Boolean) {
      response = (Boolean) args[0];
    } else if (args[0] instanceof String) {
      response = Boolean.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Boolean or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, response);
  }

  public static FirmataMessage staticParse(ReadBuffer readBuffer, Boolean response)
      throws ParseException {
    readBuffer.pullContext("FirmataMessage");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte messageType =
        readDiscriminatorField(
            "messageType",
            readUnsignedByte(readBuffer, 4),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    FirmataMessageBuilder builder = null;
    if (EvaluationHelper.equals(messageType, (byte) 0xE)) {
      builder = FirmataMessageAnalogIO.staticParseFirmataMessageBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(messageType, (byte) 0x9)) {
      builder = FirmataMessageDigitalIO.staticParseFirmataMessageBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(messageType, (byte) 0xC)) {
      builder =
          FirmataMessageSubscribeAnalogPinValue.staticParseFirmataMessageBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(messageType, (byte) 0xD)) {
      builder =
          FirmataMessageSubscribeDigitalPinValue.staticParseFirmataMessageBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(messageType, (byte) 0xF)) {
      builder = FirmataMessageCommand.staticParseFirmataMessageBuilder(readBuffer, response);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "messageType="
              + messageType
              + "]");
    }

    readBuffer.closeContext("FirmataMessage");
    // Create the instance
    FirmataMessage _firmataMessage = builder.build();
    return _firmataMessage;
  }

  public interface FirmataMessageBuilder {
    FirmataMessage build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FirmataMessage)) {
      return false;
    }
    FirmataMessage that = (FirmataMessage) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
