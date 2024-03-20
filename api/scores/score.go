package scores

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zSnails/VRAirHockey-Backend/db"
	"github.com/zSnails/VRAirHockey-Backend/store"
)

func GetScores(w http.ResponseWriter, r *http.Request) {
	player := getPlayer(r.Context())
	if player == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	DB := db.Get()
	queries := store.New(DB)
	scores, err := queries.GetPlayerScores(r.Context(), player.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(scores)
}

func getPlayer(ctx context.Context) *store.Player {
	player := ctx.Value("player")
	return player.(*store.Player)
}

type RegisterScorePayload struct {
	Score int64 `json:"score"`
}

func RegisterScore(w http.ResponseWriter, r *http.Request) {
	player := getPlayer(r.Context())

	DB := db.Get()
	queries := store.New(DB)

	var payload RegisterScorePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	score, err := queries.RegisterPlayerScore(r.Context(), store.RegisterPlayerScoreParams{
		Score:    payload.Score,
		PlayerID: player.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(score)
}
