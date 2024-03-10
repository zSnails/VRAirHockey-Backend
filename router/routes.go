package router

import (
	"github.com/gorilla/mux"
	"github.com/zSnails/VRAirHockey-Backend/api/auth"
	"github.com/zSnails/VRAirHockey-Backend/api/scores"
)

func registerAuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth/register", auth.Register)
	router.HandleFunc("/auth/login", auth.Login)
}

func registerScoreRoutes(router *mux.Router) {
	router.HandleFunc("/api/scores/upload", scores.RegisterScore)
	router.HandleFunc("/api/scores/", scores.GetScores)
}
