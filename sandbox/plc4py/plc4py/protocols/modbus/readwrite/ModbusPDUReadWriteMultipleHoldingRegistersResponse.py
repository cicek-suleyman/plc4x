#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

from dataclasses import dataclass

from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDUBuilder
from plc4py.spi.generation.ReadBuffer import ReadBuffer
from plc4py.spi.generation.WriteBuffer import WriteBuffer
from typing import List
import math


@dataclass
class ModbusPDUReadWriteMultipleHoldingRegistersResponse(PlcMessage, ModbusPDU):
    value: List[int]
    # Accessors for discriminator values.
    error_flag: bool = False
    function_flag: int = 0x17
    response: bool = True

    def __post_init__(self):
        super().__init__()

    def serialize_modbus_pdu_child(self, write_buffer: WriteBuffer):
        write_buffer.push_context("ModbusPDUReadWriteMultipleHoldingRegistersResponse")

        # Implicit Field (byte_count) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
        byte_count: int = int(len(self.value))
        write_buffer.write_unsigned_byte(byte_count, logical_name="byteCount")

        # Array Field (value)
        write_buffer.write_byte_array(self.value, logical_name="value")

        write_buffer.pop_context("ModbusPDUReadWriteMultipleHoldingRegistersResponse")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.get_length_in_bits() / 8.0)))

    def get_length_in_bits(self) -> int:
        length_in_bits: int = super().get_length_in_bits()
        _value: ModbusPDUReadWriteMultipleHoldingRegistersResponse = self

        # Implicit Field (byteCount)
        length_in_bits += 8

        # Array field
        if self.value != None:
            length_in_bits += 8 * len(self.value)

        return length_in_bits

    @staticmethod
    def static_parse_builder(read_buffer: ReadBuffer, response: bool):
        read_buffer.push_context("ModbusPDUReadWriteMultipleHoldingRegistersResponse")

        byte_count: int = read_implicit_field("byteCount", read_unsigned_short)

        self.value = read_buffer.read_byte_array("value", int(byte_count))

        read_buffer.pop_context("ModbusPDUReadWriteMultipleHoldingRegistersResponse")
        # Create the instance
        return ModbusPDUReadWriteMultipleHoldingRegistersResponseBuilder(value)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUReadWriteMultipleHoldingRegistersResponse):
            return False

        that: ModbusPDUReadWriteMultipleHoldingRegistersResponse = (
            ModbusPDUReadWriteMultipleHoldingRegistersResponse(o)
        )
        return (self.value == that.value) and super().equals(that) and True

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            write_buffer_box_based.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(write_buffer_box_based.get_box()) + "\n"


@dataclass
class ModbusPDUReadWriteMultipleHoldingRegistersResponseBuilder(ModbusPDUBuilder):
    value: List[int]

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUReadWriteMultipleHoldingRegistersResponse:
        modbus_pdu_read_write_multiple_holding_registers_response: ModbusPDUReadWriteMultipleHoldingRegistersResponse = ModbusPDUReadWriteMultipleHoldingRegistersResponse(
            self.value
        )
        return modbus_pdu_read_write_multiple_holding_registers_response
