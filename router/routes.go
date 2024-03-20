package router

import (
	"github.com/gorilla/mux"
	"github.com/zSnails/VRAirHockey-Backend/api/auth"
	"github.com/zSnails/VRAirHockey-Backend/api/scores"
)

func registerAuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth/register", auth.Register).Methods("POST", "PUT")
	router.HandleFunc("/auth/login", auth.Login).Methods("POST", "PUT")
}

func registerScoreRoutes(router *mux.Router) {
	router.HandleFunc("/scores", scores.RegisterScore).Methods("POST", "PUT")
	router.HandleFunc("/scores", scores.GetScores).Methods("GET")
}
