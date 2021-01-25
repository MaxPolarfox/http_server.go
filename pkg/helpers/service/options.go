package service

import (
	"github.com/MaxPolarfox/http_server/pkg/helpers/logger"
)

type Options struct {
	ServiceName    string         `json:"serviceName"`
	HttpPort       int            `json:"httpPort"`
	Logger         logger.Options `json:"logger"`
	JSONSchemasDir string         `json:"schemaDir"`
}
