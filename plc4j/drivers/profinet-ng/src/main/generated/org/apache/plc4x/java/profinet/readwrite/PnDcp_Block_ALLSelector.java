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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnDcp_Block_ALLSelector extends PnDcp_Block implements Message {

  // Accessors for discriminator values.
  public PnDcp_BlockOptions getOption() {
    return PnDcp_BlockOptions.ALL_SELECTOR_OPTION;
  }

  public Short getSuboption() {
    return (short) 0xFF;
  }

  public PnDcp_Block_ALLSelector() {
    super();
  }

  @Override
  protected void serializePnDcp_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnDcp_Block_ALLSelector");

    writeBuffer.popContext("PnDcp_Block_ALLSelector");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnDcp_Block_ALLSelector _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    return lengthInBits;
  }

  public static PnDcp_BlockBuilder staticParsePnDcp_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnDcp_Block_ALLSelector");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    readBuffer.closeContext("PnDcp_Block_ALLSelector");
    // Create the instance
    return new PnDcp_Block_ALLSelectorBuilderImpl();
  }

  public static class PnDcp_Block_ALLSelectorBuilderImpl implements PnDcp_Block.PnDcp_BlockBuilder {

    public PnDcp_Block_ALLSelectorBuilderImpl() {}

    public PnDcp_Block_ALLSelector build() {
      PnDcp_Block_ALLSelector pnDcp_Block_ALLSelector = new PnDcp_Block_ALLSelector();
      return pnDcp_Block_ALLSelector;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_Block_ALLSelector)) {
      return false;
    }
    PnDcp_Block_ALLSelector that = (PnDcp_Block_ALLSelector) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
