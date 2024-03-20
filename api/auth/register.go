package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zSnails/VRAirHockey-Backend/db"
	"github.com/zSnails/VRAirHockey-Backend/store"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequestPayload struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func register(ctx context.Context, payload *RegisterRequestPayload) (store.Player, error) {
	DB := db.Get()
	queries := store.New(DB)
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return store.Player{}, err
	}
	defer func() { _ = tx.Rollback() }()

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return store.Player{}, err
	}

	queries = queries.WithTx(tx)
	created, err := queries.CreatePlayer(ctx, store.CreatePlayerParams{
		Name:  payload.Name,
		Email: payload.Email,
	})
	if err != nil {
		return store.Player{}, err
	}

	err = queries.CreateAuth(ctx, store.CreateAuthParams{
		Hash:     string(hash),
		PlayerID: created.ID.(int64),
	})
	if err != nil {
		return store.Player{}, err
	}

	_ = tx.Commit()
	return created, nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterRequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	created, err := register(r.Context(), &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(created)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
