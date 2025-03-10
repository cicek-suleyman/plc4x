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

public class BACnetConstructedDataFDBBMDAddress extends BACnetConstructedData implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return null;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.FD_BBMD_ADDRESS;
  }

  // Properties.
  protected final BACnetHostNPort fDBBMDAddress;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataFDBBMDAddress(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetHostNPort fDBBMDAddress,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.fDBBMDAddress = fDBBMDAddress;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetHostNPort getFDBBMDAddress() {
    return fDBBMDAddress;
  }

  public BACnetHostNPort getActualValue() {
    return (BACnetHostNPort) (getFDBBMDAddress());
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConstructedDataFDBBMDAddress");

    // Simple Field (fDBBMDAddress)
    writeSimpleField("fDBBMDAddress", fDBBMDAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetHostNPort actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetConstructedDataFDBBMDAddress");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataFDBBMDAddress _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (fDBBMDAddress)
    lengthInBits += fDBBMDAddress.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataFDBBMDAddress");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetHostNPort fDBBMDAddress =
        readSimpleField(
            "fDBBMDAddress",
            new DataReaderComplexDefault<>(
                () -> BACnetHostNPort.staticParse(readBuffer), readBuffer));
    BACnetHostNPort actualValue =
        readVirtualField("actualValue", BACnetHostNPort.class, fDBBMDAddress);

    readBuffer.closeContext("BACnetConstructedDataFDBBMDAddress");
    // Create the instance
    return new BACnetConstructedDataFDBBMDAddressBuilderImpl(
        fDBBMDAddress, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataFDBBMDAddressBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final BACnetHostNPort fDBBMDAddress;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataFDBBMDAddressBuilderImpl(
        BACnetHostNPort fDBBMDAddress,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.fDBBMDAddress = fDBBMDAddress;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataFDBBMDAddress build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataFDBBMDAddress bACnetConstructedDataFDBBMDAddress =
          new BACnetConstructedDataFDBBMDAddress(
              openingTag,
              peekedTagHeader,
              closingTag,
              fDBBMDAddress,
              tagNumber,
              arrayIndexArgument);
      return bACnetConstructedDataFDBBMDAddress;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataFDBBMDAddress)) {
      return false;
    }
    BACnetConstructedDataFDBBMDAddress that = (BACnetConstructedDataFDBBMDAddress) o;
    return (getFDBBMDAddress() == that.getFDBBMDAddress()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFDBBMDAddress());
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
