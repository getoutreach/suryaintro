// Copyright 2023 Outreach Corporation. All Rights Reserved.

// Description: This file contains the gRPC client implementation for the
// suryaintro service.

package suryaintro //nolint:revive // Why: We allow [-_].

import (
	"context"
	"fmt"

	"github.com/getoutreach/services/pkg/grpcx"
	"github.com/getoutreach/suryaintro/api"

	// imports added by modules
	"github.com/getoutreach/mint/pkg/authn"
	// end imports added by modules
)

// New returns a new grpc client for the suryaintro Service
//
// The client is concurrency safe and handles reconnecting.
// All calls automatically handle logging, tracing, metrics,
// service discovery, and authn.
func New(ctx context.Context) (api.Service, error) {
	// Inserted by modules

	authnClient, err := authn.NewClient(ctx, &authn.Config{
		Audience:             "suryaintro@outreach.cloud",
		ForwardAccountsToken: true,
		MintDisabled:         false,
	})
	if err != nil {
		return nil, err
	}
	// End Inserted by modules

	clientOpts := []grpcx.ClientConnOption{
		grpcx.WithServiceDiscovery(),
		// Inserted by modules
		grpcx.WithAuthn(authnClient),
		// End Inserted by modules
	}

	conn, err := grpcx.NewClientConn(ctx, "suryaintro", clientOpts...)
	if err != nil {
		return nil, err
	}
	return &client{
		closers: []func(ctx context.Context) error{
			// Closers Inserted by modules
			authnClient.Close,
			// End closers inserted by modules
			func(_ context.Context) error { return conn.Close() },
		},
		SuryaintroClient: api.NewSuryaintroClient(conn),
	}, nil
}

// client is the type that actually implements the correct interface to serve as
// a gRPC client for the rms service as per the protobuf files.
type client struct {
	closers []func(ctx context.Context) error
	api.SuryaintroClient
	// Place your client struct data here
}

// Close is necessary to avoid potential resource leaks
func (c client) Close(ctx context.Context) error {
	errors := make([]error, 0)
	for _, fn := range c.closers {
		if err := fn(ctx); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("failed to close client: %v", errors)
	}

	return nil
}

// Place any client handler functions for your service here
func (c client) Ping(ctx context.Context, message string) (string, error) {
	in := &api.PingRequest{Message: message}
	resp, err := c.SuryaintroClient.Ping(ctx, in)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

func (c client) Pong(ctx context.Context, message string) (string, error) {
	in := &api.PongRequest{Message: message}
	resp, err := c.SuryaintroClient.Pong(ctx, in)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

func (c client) Joke(ctx context.Context) (string, error) {
	in := &api.JokeRequest{}
	resp, err := c.SuryaintroClient.Joke(ctx, in)
	if err != nil {
		return "", err
	}
	return resp.Joke, nil
}
