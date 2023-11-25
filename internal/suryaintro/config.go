// Copyright 2023 Outreach Corporation. All Rights Reserved.

// Description: This file is the focal point of configuration that needs passed
// to various parts of the service.
// Managed: true

package suryaintro //nolint:revive // Why: We allow [-_].

import (
	"context"

	"github.com/getoutreach/gobox/pkg/cfg"
	"github.com/getoutreach/gobox/pkg/log"
	// <<Stencil::Block(configImports)>>
	// <</Stencil::Block>>
)

// Config tracks config needed for suryaintro
type Config struct {
	ListenHost string `yaml:"ListenHost"`
	HTTPPort   int    `yaml:"HTTPPort"`
	GRPCPort   int    `yaml:"GRPCPort"`

	// <<Stencil::Block(config)>>

	// <</Stencil::Block>>
}

// MarshalLog can be used to write config to log
func (c *Config) MarshalLog(addfield func(key string, value interface{})) {
	// <<Stencil::Block(marshalconfig)>>

	// <</Stencil::Block>>
}

// LoadConfig returns a new Config type that has been loaded in accordance to the environment
// that the service was deployed in, with all necessary tweaks made before returning.
// nolint: funlen // Why: This function is long for extensibility reasons since it is generated by stencil.
func LoadConfig(ctx context.Context) (*Config, error) {
	// NOTE: Defaults should generally be set in the config
	// override jsonnet file: deployments/suryaintro/suryaintro.config.jsonnet
	c := Config{
		// Defaults to [::]/0.0.0.0 which will broadcast to all reachable
		// IPs on a server on the given port for the respective service.
		ListenHost: "",
		HTTPPort:   8000,
		GRPCPort:   5000,
		/// !!! DEPRECATED: This block is deprecated and will be removed in an upcoming release.
		/// All configuration should be defined in deployments/suryaintro/suryaintro.config.jsonnet.
		///
		// <<Stencil::Block(defconfig)>>
		// <</Stencil::Block>>
	}

	// Attempt to load a local config file on top of the defaults
	if err := cfg.Load("suryaintro.yaml", &c); err != nil {
		return nil, err
	}

	// Do any necessary tweaks/augmentations to your configuration here
	// <<Stencil::Block(configtweak)>>

	// <</Stencil::Block>>

	log.Info(ctx, "Configuration data of the application:\n", &c)

	return &c, nil
}
