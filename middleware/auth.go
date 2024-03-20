package middleware

import (
	"context"
	"net/http"

	"github.com/zSnails/VRAirHockey-Backend/session"
	"github.com/zSnails/VRAirHockey-Backend/store"
)

type PlayerKey string

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, err := session.Store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if sess.IsNew {
			http.Error(w, "New Session", http.StatusUnauthorized)
			return
		}

		value, ok := sess.Values["authenticated"]
		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		auth, ok := value.(bool)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if !auth {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		player, ok := sess.Values["player"]
		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		re := r.WithContext(context.WithValue(r.Context(), PlayerKey("player"), player.(store.Player)))
		h.ServeHTTP(w, re)
	})
}
