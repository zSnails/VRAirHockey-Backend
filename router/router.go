package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/VRAirHockey-Backend/middleware"
)

var logger = logrus.WithField("service", "router")

func NewRouter() http.Handler {
	logger.Info("Starting router")
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	// TODO: auth middleware and all that shit.
	logger.Info("Registering routes")
	registerAuthRoutes(r)
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.Auth)
	registerScoreRoutes(api)
	return r
}
