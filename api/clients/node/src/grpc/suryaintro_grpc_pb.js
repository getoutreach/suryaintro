// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2023 Outreach Corporation. All Rights Reserved.
// Please modify this to match the interface specified in suryaintro.go
'use strict';
var grpc = require('@grpc/grpc-js');
var suryaintro_pb = require('./suryaintro_pb.js');

function serialize_suryaintro_api_JokeRequest(arg) {
  if (!(arg instanceof suryaintro_pb.JokeRequest)) {
    throw new Error('Expected argument of type suryaintro.api.JokeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_suryaintro_api_JokeRequest(buffer_arg) {
  return suryaintro_pb.JokeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_suryaintro_api_JokeResponse(arg) {
  if (!(arg instanceof suryaintro_pb.JokeResponse)) {
    throw new Error('Expected argument of type suryaintro.api.JokeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_suryaintro_api_JokeResponse(buffer_arg) {
  return suryaintro_pb.JokeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_suryaintro_api_PingRequest(arg) {
  if (!(arg instanceof suryaintro_pb.PingRequest)) {
    throw new Error('Expected argument of type suryaintro.api.PingRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_suryaintro_api_PingRequest(buffer_arg) {
  return suryaintro_pb.PingRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_suryaintro_api_PingResponse(arg) {
  if (!(arg instanceof suryaintro_pb.PingResponse)) {
    throw new Error('Expected argument of type suryaintro.api.PingResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_suryaintro_api_PingResponse(buffer_arg) {
  return suryaintro_pb.PingResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_suryaintro_api_PongRequest(arg) {
  if (!(arg instanceof suryaintro_pb.PongRequest)) {
    throw new Error('Expected argument of type suryaintro.api.PongRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_suryaintro_api_PongRequest(buffer_arg) {
  return suryaintro_pb.PongRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_suryaintro_api_PongResponse(arg) {
  if (!(arg instanceof suryaintro_pb.PongResponse)) {
    throw new Error('Expected argument of type suryaintro.api.PongResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_suryaintro_api_PongResponse(buffer_arg) {
  return suryaintro_pb.PongResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Suryaintro is the suryaintro service.
var SuryaintroService = exports.SuryaintroService = {
  ping: {
    path: '/suryaintro.api.Suryaintro/Ping',
    requestStream: false,
    responseStream: false,
    requestType: suryaintro_pb.PingRequest,
    responseType: suryaintro_pb.PingResponse,
    requestSerialize: serialize_suryaintro_api_PingRequest,
    requestDeserialize: deserialize_suryaintro_api_PingRequest,
    responseSerialize: serialize_suryaintro_api_PingResponse,
    responseDeserialize: deserialize_suryaintro_api_PingResponse,
  },
  pong: {
    path: '/suryaintro.api.Suryaintro/Pong',
    requestStream: false,
    responseStream: false,
    requestType: suryaintro_pb.PongRequest,
    responseType: suryaintro_pb.PongResponse,
    requestSerialize: serialize_suryaintro_api_PongRequest,
    requestDeserialize: deserialize_suryaintro_api_PongRequest,
    responseSerialize: serialize_suryaintro_api_PongResponse,
    responseDeserialize: deserialize_suryaintro_api_PongResponse,
  },
  joke: {
    path: '/suryaintro.api.Suryaintro/Joke',
    requestStream: false,
    responseStream: false,
    requestType: suryaintro_pb.JokeRequest,
    responseType: suryaintro_pb.JokeResponse,
    requestSerialize: serialize_suryaintro_api_JokeRequest,
    requestDeserialize: deserialize_suryaintro_api_JokeRequest,
    responseSerialize: serialize_suryaintro_api_JokeResponse,
    responseDeserialize: deserialize_suryaintro_api_JokeResponse,
  },
};

exports.SuryaintroClient = grpc.makeGenericClientConstructor(SuryaintroService);
