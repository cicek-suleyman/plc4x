/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.iec608705104;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.PlcDriverManager;

import java.util.concurrent.CompletableFuture;

public class ManualTest {

    public static void main(String[] args) throws Exception {
        CompletableFuture<Void> shutdown = new CompletableFuture<>();
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            shutdown.complete(null);
        }));
        try (PlcConnection plcConnection = PlcDriverManager.getDefault().getConnectionManager().getConnection("iec-60870-5-104://192.168.23.10")) {
            // Wait till shutdown.
            shutdown.get();
        }
    }

}
