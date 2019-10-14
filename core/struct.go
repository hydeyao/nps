package core

import (
	"context"
	"errors"
)

type Stage uint8

// These constants are meant to describe the stage in which the plugin is running.
const (
	STAGE_START_RUN_END Stage = iota
	STAGE_START_RUN
	STAGE_START_END
	STAGE_RUN_END
	STAGE_START
	STAGE_END
	STAGE_RUN
	PROXY_CONNECTION_TYPE = "proxy_target_type"
	PROXY_CONNECTION_ADDR = "proxy_target_addr"
	PROXY_CONNECTION_PORT = "proxy_target_port"
	CLIENT_CONNECTION     = "clientConn"
	BRIDGE                = "bridge"
	CLIENT_ID             = "client_id"
)

var (
	CLIENT_CONNECTION_NOT_EXIST = errors.New("the client connection is not exist")
	BRIDGE_NOT_EXIST            = errors.New("the client connection is not exist")
	REQUEST_EOF                 = errors.New("the request has finished")
	CLIENT_ID_NOT_EXIST        = errors.New("the request has finished")
)

// Plugin interface, all plugins must implement those functions.
type Plugin interface {
	GetConfigName() *NpsConfigs
	GetBeforePlugin() Plugin
	GetStage() Stage
	Start(ctx context.Context, config map[string]string) error
	Run(ctx context.Context, config map[string]string) error
	End(ctx context.Context, config map[string]string) error
}
