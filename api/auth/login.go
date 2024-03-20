package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/zSnails/VRAirHockey-Backend/db"
	"github.com/zSnails/VRAirHockey-Backend/session"
	"github.com/zSnails/VRAirHockey-Backend/store"
)

type LogingRequestPayload struct {
	Email     string `json:"email"`
	Passsword string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var payload LogingRequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	DB := db.Get()
	queries := store.New(DB)
	registered, err := queries.IsRegistered(r.Context(), payload.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if registered != 1 {
		http.Error(w, "User not registered", http.StatusNotFound)
		return
	}

	sess, err := session.Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(time.Hour * 24),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	player, err := queries.GetPlayerByMail(r.Context(), payload.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sess.Values["authenticated"] = true
	sess.Values["player"] = player

	err = sess.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
