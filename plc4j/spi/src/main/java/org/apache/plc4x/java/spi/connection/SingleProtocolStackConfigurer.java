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
package org.apache.plc4x.java.spi.connection;

import static org.apache.plc4x.java.spi.configuration.ConfigurationFactory.*;

import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandler;
import io.netty.channel.ChannelPipeline;
import io.netty.handler.codec.MessageToMessageCodec;
import org.apache.plc4x.java.api.authentication.PlcAuthentication;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.listener.EventListener;
import org.apache.plc4x.java.spi.Plc4xNettyWrapper;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.configuration.Configuration;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.generation.ByteOrder;
import org.apache.plc4x.java.spi.generation.Message;
import org.apache.plc4x.java.spi.generation.MessageInput;
import org.apache.plc4x.java.spi.generation.MessageOutput;

import java.lang.reflect.InvocationTargetException;
import java.util.List;
import java.util.function.Consumer;
import java.util.function.ToIntFunction;

/**
 * Builds a Protocol Stack.
 */
public class SingleProtocolStackConfigurer<BASE_PACKET_CLASS extends Message> implements ProtocolStackConfigurer<BASE_PACKET_CLASS> {

    private final Class<BASE_PACKET_CLASS> basePacketClass;
    private final ByteOrder byteOrder;
    private final Class<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocolClass;
    private final Class<? extends DriverContext> driverContextClass;
    private final MessageInput<BASE_PACKET_CLASS> messageInput;
    private final MessageOutput<BASE_PACKET_CLASS> messageOutput;
    private final Class<? extends ToIntFunction<ByteBuf>> packetSizeEstimatorClass;
    private final Class<? extends Consumer<ByteBuf>> corruptPacketRemoverClass;
    private final MessageToMessageCodec<ByteBuf, ByteBuf> encryptionHandler;
    private final Object[] parserArgs;

    public static <BPC extends Message> SingleProtocolStackBuilder<BPC> builder(Class<BPC> basePacketClass, MessageInput<BPC> messageInput) {
        return new SingleProtocolStackBuilder<>(basePacketClass, messageInput, null);
    }

    public static <BPC extends Message> SingleProtocolStackBuilder<BPC> builder(Class<BPC> basePacketClass, MessageInput<BPC> messageInput, MessageOutput<BPC> messageOutput) {
        return new SingleProtocolStackBuilder<>(basePacketClass, messageInput, messageOutput);
    }

    /**
     * Only accessible via Builder
     */
    SingleProtocolStackConfigurer(Class<BASE_PACKET_CLASS> basePacketClass,
                                  ByteOrder byteOrder,
                                  Object[] parserArgs,
                                  Class<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocol,
                                  Class<? extends DriverContext> driverContextClass,
                                  MessageInput<BASE_PACKET_CLASS> messageInput,
                                  MessageOutput<BASE_PACKET_CLASS> messageOutput,
                                  Class<? extends ToIntFunction<ByteBuf>> packetSizeEstimatorClass,
                                  Class<? extends Consumer<ByteBuf>> corruptPacketRemoverClass,
                                  MessageToMessageCodec<ByteBuf, ByteBuf> encryptionHandler) {
        this.basePacketClass = basePacketClass;
        this.byteOrder = byteOrder;
        this.parserArgs = parserArgs;
        this.protocolClass = protocol;
        this.driverContextClass = driverContextClass;
        this.messageInput = messageInput;
        this.messageOutput = messageOutput;
        this.packetSizeEstimatorClass = packetSizeEstimatorClass;
        this.corruptPacketRemoverClass = corruptPacketRemoverClass;
        this.encryptionHandler = encryptionHandler;
    }

    private ChannelHandler getMessageCodec(Configuration configuration) {
        return new GeneratedProtocolMessageCodec<>(basePacketClass, messageInput, messageOutput, byteOrder, parserArgs,
            packetSizeEstimatorClass != null ? configure(configuration, createInstance(packetSizeEstimatorClass)) : null,
            corruptPacketRemoverClass != null ? configure(configuration, createInstance(corruptPacketRemoverClass)) : null);
    }

    /**
     * Applies the given Stack to the Pipeline
     */
    @Override
    public Plc4xProtocolBase<BASE_PACKET_CLASS> configurePipeline(Configuration configuration, ChannelPipeline pipeline,
                                                                  PlcAuthentication authentication, boolean passive,
                                                                  List<EventListener> ignore) {
        if (this.encryptionHandler != null) {
            pipeline.addLast(this.encryptionHandler);
        }
        pipeline.addLast(getMessageCodec(configuration));
        Plc4xProtocolBase<BASE_PACKET_CLASS> protocol = configure(configuration, createInstance(protocolClass));
        if (driverContextClass != null) {
            protocol.setDriverContext(configure(configuration, createInstance(driverContextClass)));
        }
        Plc4xNettyWrapper<BASE_PACKET_CLASS> context = new Plc4xNettyWrapper<>(configuration.getTimeoutManager(), pipeline, passive, protocol,
            authentication, basePacketClass);
        pipeline.addLast(context);
        return protocol;
    }

    private <T> T createInstance(Class<T> clazz, Object... args) {
        try {
            Class<?>[] parameterTypes = new Class<?>[args.length];
            for (int i = 0; i < args.length; i++) {
                parameterTypes[i] = args[i].getClass();
            }
            return clazz.getDeclaredConstructor(parameterTypes).newInstance(args);
        } catch (InstantiationException | IllegalAccessException | InvocationTargetException | NoSuchMethodException e) {
            throw new PlcRuntimeException("Error creating instance of class " + clazz.getName());
        }
    }

    /**
     * Used to Build Instances of {@link SingleProtocolStackConfigurer}.
     *
     * @param <BASE_PACKET_CLASS> Type of Created Message that is Exchanged.
     */
    public static final class SingleProtocolStackBuilder<BASE_PACKET_CLASS extends Message> {

        private final Class<BASE_PACKET_CLASS> basePacketClass;
        private final MessageInput<BASE_PACKET_CLASS> messageInput;
        private final MessageOutput<BASE_PACKET_CLASS> messageOutput;
        private Class<? extends DriverContext> driverContextClass;
        private ByteOrder byteOrder = ByteOrder.BIG_ENDIAN;
        private Object[] parserArgs;
        private Class<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocol;
        private Class<? extends ToIntFunction<ByteBuf>> packetSizeEstimator;
        private Class<? extends Consumer<ByteBuf>> corruptPacketRemover;
        private MessageToMessageCodec<ByteBuf, ByteBuf> encryptionHandler;

        public SingleProtocolStackBuilder(Class<BASE_PACKET_CLASS> basePacketClass, MessageInput<BASE_PACKET_CLASS> messageInput, MessageOutput<BASE_PACKET_CLASS> messageOutput) {
            this.basePacketClass = basePacketClass;
            this.messageInput = messageInput;
            this.messageOutput = messageOutput;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withDriverContext(Class<? extends DriverContext> driverContextClass) {
            this.driverContextClass = driverContextClass;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> byteOrder(ByteOrder byteOrder) {
            this.byteOrder = byteOrder;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> bigEndian() {
            this.byteOrder = ByteOrder.BIG_ENDIAN;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> littleEndian() {
            this.byteOrder = ByteOrder.LITTLE_ENDIAN;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withParserArgs(Object... parserArgs) {
            this.parserArgs = parserArgs;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withProtocol(Class<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocol) {
            this.protocol = protocol;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withPacketSizeEstimator(Class<? extends ToIntFunction<ByteBuf>> packetSizeEstimator) {
            this.packetSizeEstimator = packetSizeEstimator;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withCorruptPacketRemover(Class<? extends Consumer<ByteBuf>> corruptPacketRemover) {
            this.corruptPacketRemover = corruptPacketRemover;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withEncryptionHandler(MessageToMessageCodec<ByteBuf, ByteBuf> encryptionHandler) {
            this.encryptionHandler = encryptionHandler;
            return this;
        }

        public SingleProtocolStackConfigurer<BASE_PACKET_CLASS> build() {
            assert this.protocol != null;
            return new SingleProtocolStackConfigurer<>(basePacketClass, byteOrder, parserArgs, protocol,
                driverContextClass, messageInput, messageOutput, packetSizeEstimator, corruptPacketRemover,
                encryptionHandler);
        }

    }

}
