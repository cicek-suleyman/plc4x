<!--
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
 -->
# PLC4X Apache NiFi Integration

## Common properties
The following properties applies to all Plc4x Processors:
* Connection String: A constant connection string such as `s7://10.105.143.7:102?remote-rack=0&remote-slot=1&controller-type=S7_1200`.
* Read/Write/Subscribe timeout (miliseconds): Specifies the time in milliseconds for the connection to return a timeout. In case of subscription the timeout is used to renew connections.
* Address Access Strategy: defines how the processor obtains the PLC addresses. It can take 2 values:
  * **Properties as Addreses:** 
      For each variable, add a new property to the processor where the property name matches the variable name, and the variable value corresponds to the address tag. 

    An *example* of these properties for reading values from a S7-1200:
    - *var1:* *%DB1:DBX0.0:BOOL*
    - *var2:* *%DB1:DBX0.1:BOOL*
    - *var3:* *%DB1:DBB01:BYTE*
    - *var4:* *%DB1:DBW02:WORD*
    - *var5:* *%DB1:DBW04:INT*

  * **Address Text:**
    Property *Address Text* must be supplied in JSON format that contains variable name and address tag. Expression Language is supported.

    Using the same example as before:
    - *Address Text*:  
    ```json
    {
      "var1" : "%DB1:DBX0.0:BOOL",
      "var2" : "%DB1:DBX0.1:BOOL",
      "var3" : "%DB1:DBB01:BYTE",
      "var4" : "%DB1:DBW02:WORD",
      "var5" : "%DB1:DBW04:INT" 
    }
    ```
    If this JSON is in an attribute `plc4x.addresses` it can be accessed with *Address Text*=`${plc4x.addresses}`. 


When reading from a PLC the response is used to create a mapping between Plc types into Avro. The mapping is done as follows:

Table of data mapping between plc data and Avro types (as specified in [Avro specification](https://avro.apache.org/docs/1.11.1/specification/#primitive-types)).


| PLC type | Avro Type |
|----------:|-----------|
| PlcBOOL | boolean |
| PlcBYTE | bytes |
| PlcSINT | int |
| PlcINT | int |
| PlcLINT | long |
| PlcREAL | float |
| PlcLREAL | double |
| PlcCHAR | string |
| PlcDATE_AND_TIME | string |
| PlcDATE | string |
| PlcDINT | string |
| PlcDWORD | string |
| PlcLTIME | string |
| PlcLWORD | string |
| PlcNull | string |
| PlcSTRING | string |
| PlcTIME_OF_DAY | string |
| PlcTIME | string |
| PlcUDINT | string |
| PlcUINT | string |
| PlcULINT | string |
| PlcUSINT | string |
| PlcWCHAR | string |
| PlcWORD | string |
| ELSE | string |


Also, it is important to keep in mind the Processor Scheduling Configuration. Using the parameter **Run Schedule** (for example to *1 sec*), the reading frequency can be set. Note that by default, this value is defined to 0 sec (as fast as possible).


## Plc4xSinkProcessor

## Plc4xSourceProcessor

## Plc4xSinkRecordProcessor

This processor is <ins>record oriented</ins>, reads from a formated input flowfile content using a Record Reader (for further information see [NiFi Documentation](https://nifi.apache.org/docs/nifi-docs/html/record-path-guide.html#overview)). 

The Plc4xSinkRecord Processor can be configured using the common properties defined above and the following property:
- *Record Reader:* Specifies the Controller Service to use for reading input variables from a flowfile. The Record Reader may use Inherit Schema to emulate the inferred schema behavior.


For the **Record Reader** property, any reader included in NiFi could be used, such as JSON, CSV, etc (also custom readers can be created).


## Plc4xSourceRecordProcessor

This processor is <ins>record oriented</ins>, formatting output flowfile content using a Record Writer (for further information see [NiFi Documentation](https://nifi.apache.org/docs/nifi-docs/html/record-path-guide.html#overview)). 

The Plc4xSourceRecord Processor can be configured using the common properties defined above and the following **properties**:

- *Record Writer:* Specifies the Controller Service to use for writing results to a FlowFile. The Record Writer may use Inherit Schema to emulate the inferred schema behavior.

## Plc4xListenRecordProcessor
This processor is <ins>record oriented</ins>, formatting output flowfile content using a Record Writer (for further information see [NiFi Documentation](https://nifi.apache.org/docs/nifi-docs/html/record-path-guide.html#overview)). 

The Plc4xListenRecordProcessor can be configured using the common properties defined above and the following properties:
- *Subscription Type*: sets the subscription type. It can be "Change", "Event" or "Cyclic". The subscritpion types available for each driver are stated in the documentation.
- *Cyclic polling interval*: In case of "Cyclic" subscription type a time interval must be provided. Must be smaller than the subscription timeout.

# Example

An *example* for reading values from a S7-1200:

- *PLC connection String:* *s7://10.105.143.7:102?remote-rack=0&remote-slot=1&controller-type=S7_1200*
- *Record Writer:* *PLC4x Embedded - AvroRecordSetWriter*
- *Read timeout (miliseconds):* *10000*
- *var1:* *%DB1:DBX0.0:BOOL*
- *var2:* *%DB1:DBX0.1:BOOL*
- *var3:* *%DB1:DBB01:BYTE*
- *var4:* *%DB1:DBW02:WORD*
- *var5:* *%DB1:DBW04:INT*

Reading values using OPCUA:
- *PLC connection String:* *opcua:tcp://10.105.143.6:4840?discovery=false*
- *Record Writer:* *PLC4x Embedded - AvroRecordSetWriter*
- *Read timeout (miliseconds):* *10000*
- *AcyclicReceiveBit00:* *ns=2;i=11*
- *MaxCurrentI_max:*  *ns=2;i=33*

For the **Record Writer** property, any writer included in NiFi could be used, such as JSON, CSV, etc (also custom writers can be created). In this example, an Avro Writer is supplied, configured as follows:

- *Schema Write Strategy:* Embed Avro Schema
- *Schema Cache:* No value set
- *Schema Protocol Version:* 1
- *Schema Access Strategy:* Inherit Record Schema
- *Schema Registry:* No value set
- *Schema Name:* ${schema.name}
- *Schema Version:* No value set
- *Schema Branch:* No value set
- *Schema Text:* ${avro.schema}
- *Compression Format:* NONE
- *Cache Size:* 1000
- *Encoder Pool Size:* 32


The output flowfile will contain the PLC read values. This information is included in the flowfile content, following the Record Oriented presentation using a **schema** and the configuration specified in the Record Writer (format, schema inclusion, etc). In the schema, one tag will be included for each of the variables defined taking into account the specified datatype. Also, a *ts* (timestamp) field is additionally included containing the read date. An example of the content of a flowfile for the previously defined properties:

```
[ {
  "var1" : true,
  "var2" : false,
  "var3" : "\u0005",
  "var5" : 1992,
  "var4" : "4",
  "ts" : 1628783058433
} ]
```