import * as grpc from "@grpc/grpc-js";
import {SuryaintroClient} from "./grpc/suryaintro_grpc_pb";
import {createErrorLoggerInterceptor} from "@getoutreach/grpc-client";
import winston from "winston";
import * as find from "@getoutreach/find";

const level = "error";
const ConsoleTransport = () => {
  return new winston.transports.Console({
    level,
  });
};

/**
 * The ClientOptions interface defines the gRPC service endpoint the SuryaintroClient connects to as well
 * as any gRPC options that should be used by the client.
 */
export interface ClientOptions {
  /** This is the endpoint the client should connect to. If not specified a default is used. */
  endpoint?: string;

  /** This is the gRPC interceptors that should be run on messages passed through the client. */
  interceptors?: grpc.Interceptor[];

  /** This is the ChannelOptions map for overriding default grpc-js ChannelOptions values. */
  /* eslint-disable @typescript-eslint/no-explicit-any */
  channelOptions?: Map<string, any> | null;
}

/**
 * @param options The client options that affect its behavior
 * @returns The newly created SuryaintroClient instance
 */
export function createSuryaintroClient(
  options?: ClientOptions
): SuryaintroClient {
  const logger = winston.createLogger({transports: [ConsoleTransport()]});
  const endpoint =
    options?.endpoint || find.service("suryaintro").dnsName + ":5000";
  const clientName = "suryaintro" + ":gRPCClient";
  logger.info(`${clientName}: Endpoint information: ${endpoint}`);
  const interceptors = [createErrorLoggerInterceptor(logger, endpoint)];
  if (options?.interceptors) {
    interceptors.push(...options.interceptors);
  }
  const channelOptions = options?.channelOptions;
  return new SuryaintroClient(endpoint, grpc.credentials.createInsecure(), {
    interceptors,
    channelOptions,
  });
}
