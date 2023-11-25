// Copyright 2023 Outreach Corporation. All Rights Reserved.

// Description: This file is the entrypoint for suryaintro.
// Managed: true

// Package main implements the main entrypoint for the suryaintro service.
package main

import (
	"context"
	"os"

	"github.com/getoutreach/gobox/pkg/app"
	"github.com/getoutreach/gobox/pkg/async"
	"github.com/getoutreach/gobox/pkg/env"
	"github.com/getoutreach/gobox/pkg/events"
	"github.com/getoutreach/gobox/pkg/log"
	"github.com/getoutreach/gobox/pkg/trace"
	"github.com/getoutreach/stencil-golang/pkg/serviceactivities/automemlimit"
	"github.com/getoutreach/stencil-golang/pkg/serviceactivities/gomaxprocs"
	"github.com/getoutreach/stencil-golang/pkg/serviceactivities/shutdown"

	"github.com/getoutreach/suryaintro/internal/suryaintro"

	// Code inserted by modules
	"github.com/getoutreach/tollmon/pkg/tollway"
	// End code inserted by modules
	// Place any extra imports for your startup code here
	// <<Stencil::Block(imports)>>
	// <</Stencil::Block>>
)

// Place any customized code for your service in this block
//
// <<Stencil::Block(customized)>>

// <</Stencil::Block>>

// dependencies is a conglomerate struct used for injecting dependencies into service
// activities.
type dependencies struct {
	privateHTTP suryaintro.PrivateHTTPDependencies
	gRPC        suryaintro.GRPCDependencies
	// <<Stencil::Block(customServiceActivityDependencyInjection)>>

	// <</Stencil::Block>>
}

// main is the entrypoint for the suryaintro service.
func main() { //nolint: funlen // Why: We can't dwindle this down anymore without adding complexity.
	exitCode := 1
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
		os.Exit(exitCode)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env.ApplyOverrides()
	app.SetName("suryaintro")

	cfg, err := suryaintro.LoadConfig(ctx)
	if err != nil {
		log.Error(ctx, "failed to load config", events.NewErrorInfo(err))
		return
	}

	if err := trace.InitTracer(ctx, "suryaintro"); err != nil {
		log.Error(ctx, "tracing failed to start", events.NewErrorInfo(err))
		return
	}
	defer trace.CloseTracer(ctx)

	// Initialize variable for service activity dependency injection.
	var deps dependencies

	log.Info(ctx, "starting", app.Info(), cfg, log.F{"app.pid": os.Getpid()})

	// Code inserted by modules

	// Create a tollway instance to register tollgate instances for rate limiting http/grpc servers onto.
	tw := tollway.NewProxy(ctx)
	deps.gRPC.TollwayProxy = tw
	// End code inserted by modules

	// Place any code for your service to run before registering service activities in this block
	// <<Stencil::Block(initialization)>>

	// <</Stencil::Block>>

	acts := []async.Runner{
		shutdown.New(),
		gomaxprocs.New(),
		automemlimit.New(),
		suryaintro.NewHTTPService(cfg, &deps.privateHTTP),
		suryaintro.NewGRPCService(cfg, &deps.gRPC),

		// Place any additional ServiceActivities that your service has built here to have them handled automatically
		//
		// <<Stencil::Block(services)>>

		// <</Stencil::Block>>
	}

	// Place any code for your service to run during startup in this block
	//
	// <<Stencil::Block(startup)>>

	// <</Stencil::Block>>

	err = async.RunGroup(acts).Run(ctx)
	if shutdown.HandleShutdownConditions(ctx, err) {
		exitCode = 0
	}
}
