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

public class InformationObject_SECTION_READY extends InformationObject implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.SECTION_READY;
  }

  // Properties.
  protected final NameOfFile nof;
  protected final NameOfSection nos;
  protected final LengthOfFile lof;
  protected final SectionReadyQualifier srq;

  public InformationObject_SECTION_READY(
      int address, NameOfFile nof, NameOfSection nos, LengthOfFile lof, SectionReadyQualifier srq) {
    super(address);
    this.nof = nof;
    this.nos = nos;
    this.lof = lof;
    this.srq = srq;
  }

  public NameOfFile getNof() {
    return nof;
  }

  public NameOfSection getNos() {
    return nos;
  }

  public LengthOfFile getLof() {
    return lof;
  }

  public SectionReadyQualifier getSrq() {
    return srq;
  }

  @Override
  protected void serializeInformationObjectChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObject_SECTION_READY");

    // Simple Field (nof)
    writeSimpleField(
        "nof",
        nof,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (nos)
    writeSimpleField(
        "nos",
        nos,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (lof)
    writeSimpleField(
        "lof",
        lof,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (srq)
    writeSimpleField(
        "srq",
        srq,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObject_SECTION_READY");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObject_SECTION_READY _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nof)
    lengthInBits += nof.getLengthInBits();

    // Simple field (nos)
    lengthInBits += nos.getLengthInBits();

    // Simple field (lof)
    lengthInBits += lof.getLengthInBits();

    // Simple field (srq)
    lengthInBits += srq.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectBuilder staticParseInformationObjectBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification) throws ParseException {
    readBuffer.pullContext("InformationObject_SECTION_READY");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NameOfFile nof =
        readSimpleField(
            "nof",
            new DataReaderComplexDefault<>(() -> NameOfFile.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    NameOfSection nos =
        readSimpleField(
            "nos",
            new DataReaderComplexDefault<>(() -> NameOfSection.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    LengthOfFile lof =
        readSimpleField(
            "lof",
            new DataReaderComplexDefault<>(() -> LengthOfFile.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    SectionReadyQualifier srq =
        readSimpleField(
            "srq",
            new DataReaderComplexDefault<>(
                () -> SectionReadyQualifier.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObject_SECTION_READY");
    // Create the instance
    return new InformationObject_SECTION_READYBuilderImpl(nof, nos, lof, srq);
  }

  public static class InformationObject_SECTION_READYBuilderImpl
      implements InformationObject.InformationObjectBuilder {
    private final NameOfFile nof;
    private final NameOfSection nos;
    private final LengthOfFile lof;
    private final SectionReadyQualifier srq;

    public InformationObject_SECTION_READYBuilderImpl(
        NameOfFile nof, NameOfSection nos, LengthOfFile lof, SectionReadyQualifier srq) {
      this.nof = nof;
      this.nos = nos;
      this.lof = lof;
      this.srq = srq;
    }

    public InformationObject_SECTION_READY build(int address) {
      InformationObject_SECTION_READY informationObject_SECTION_READY =
          new InformationObject_SECTION_READY(address, nof, nos, lof, srq);
      return informationObject_SECTION_READY;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObject_SECTION_READY)) {
      return false;
    }
    InformationObject_SECTION_READY that = (InformationObject_SECTION_READY) o;
    return (getNof() == that.getNof())
        && (getNos() == that.getNos())
        && (getLof() == that.getLof())
        && (getSrq() == that.getSrq())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNof(), getNos(), getLof(), getSrq());
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
