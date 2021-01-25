package template

import (
	"github.com/MaxPolarfox/http_server/pkg/controllers"
	"github.com/MaxPolarfox/http_server/pkg/helpers/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Service struct {
	Router *httprouter.Router
	Options Options
	Logger logger.Logger
}

func NewService(options Options, env string, logger logger.Logger) *Service {
	logger.Infow("Creating New Service", "env", env)

	// Controlers
	healthcheckController := controllers.HealthcheckController{}

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/liveness", healthcheckController.Liveness)
	router.HandlerFunc(http.MethodGet, "/readiness", healthcheckController.Readiness)

	return &Service{router, options, logger}
}
