// package: suryaintro.api
// file: suryaintro.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as suryaintro_pb from "./suryaintro_pb";

interface ISuryaintroService
  extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  ping: ISuryaintroService_IPing;
  pong: ISuryaintroService_IPong;
  joke: ISuryaintroService_IJoke;
}

interface ISuryaintroService_IPing
  extends grpc.MethodDefinition<
    suryaintro_pb.PingRequest,
    suryaintro_pb.PingResponse
  > {
  path: "/suryaintro.api.Suryaintro/Ping";
  requestStream: false;
  responseStream: false;
  requestSerialize: grpc.serialize<suryaintro_pb.PingRequest>;
  requestDeserialize: grpc.deserialize<suryaintro_pb.PingRequest>;
  responseSerialize: grpc.serialize<suryaintro_pb.PingResponse>;
  responseDeserialize: grpc.deserialize<suryaintro_pb.PingResponse>;
}
interface ISuryaintroService_IPong
  extends grpc.MethodDefinition<
    suryaintro_pb.PongRequest,
    suryaintro_pb.PongResponse
  > {
  path: "/suryaintro.api.Suryaintro/Pong";
  requestStream: false;
  responseStream: false;
  requestSerialize: grpc.serialize<suryaintro_pb.PongRequest>;
  requestDeserialize: grpc.deserialize<suryaintro_pb.PongRequest>;
  responseSerialize: grpc.serialize<suryaintro_pb.PongResponse>;
  responseDeserialize: grpc.deserialize<suryaintro_pb.PongResponse>;
}
interface ISuryaintroService_IJoke
  extends grpc.MethodDefinition<
    suryaintro_pb.JokeRequest,
    suryaintro_pb.JokeResponse
  > {
  path: "/suryaintro.api.Suryaintro/Joke";
  requestStream: false;
  responseStream: false;
  requestSerialize: grpc.serialize<suryaintro_pb.JokeRequest>;
  requestDeserialize: grpc.deserialize<suryaintro_pb.JokeRequest>;
  responseSerialize: grpc.serialize<suryaintro_pb.JokeResponse>;
  responseDeserialize: grpc.deserialize<suryaintro_pb.JokeResponse>;
}

export const SuryaintroService: ISuryaintroService;

export interface ISuryaintroServer extends grpc.UntypedServiceImplementation {
  ping: grpc.handleUnaryCall<
    suryaintro_pb.PingRequest,
    suryaintro_pb.PingResponse
  >;
  pong: grpc.handleUnaryCall<
    suryaintro_pb.PongRequest,
    suryaintro_pb.PongResponse
  >;
  joke: grpc.handleUnaryCall<
    suryaintro_pb.JokeRequest,
    suryaintro_pb.JokeResponse
  >;
}

export interface ISuryaintroClient {
  ping(
    request: suryaintro_pb.PingRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PingResponse
    ) => void
  ): grpc.ClientUnaryCall;
  ping(
    request: suryaintro_pb.PingRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PingResponse
    ) => void
  ): grpc.ClientUnaryCall;
  ping(
    request: suryaintro_pb.PingRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PingResponse
    ) => void
  ): grpc.ClientUnaryCall;
  pong(
    request: suryaintro_pb.PongRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PongResponse
    ) => void
  ): grpc.ClientUnaryCall;
  pong(
    request: suryaintro_pb.PongRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PongResponse
    ) => void
  ): grpc.ClientUnaryCall;
  pong(
    request: suryaintro_pb.PongRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PongResponse
    ) => void
  ): grpc.ClientUnaryCall;
  joke(
    request: suryaintro_pb.JokeRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.JokeResponse
    ) => void
  ): grpc.ClientUnaryCall;
  joke(
    request: suryaintro_pb.JokeRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.JokeResponse
    ) => void
  ): grpc.ClientUnaryCall;
  joke(
    request: suryaintro_pb.JokeRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.JokeResponse
    ) => void
  ): grpc.ClientUnaryCall;
}

export class SuryaintroClient extends grpc.Client implements ISuryaintroClient {
  constructor(
    address: string,
    credentials: grpc.ChannelCredentials,
    options?: Partial<grpc.ClientOptions>
  );
  public ping(
    request: suryaintro_pb.PingRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PingResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public ping(
    request: suryaintro_pb.PingRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PingResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public ping(
    request: suryaintro_pb.PingRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PingResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public pong(
    request: suryaintro_pb.PongRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PongResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public pong(
    request: suryaintro_pb.PongRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PongResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public pong(
    request: suryaintro_pb.PongRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.PongResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public joke(
    request: suryaintro_pb.JokeRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.JokeResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public joke(
    request: suryaintro_pb.JokeRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.JokeResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public joke(
    request: suryaintro_pb.JokeRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: suryaintro_pb.JokeResponse
    ) => void
  ): grpc.ClientUnaryCall;
}
