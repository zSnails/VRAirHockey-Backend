package session

import (
	"os"

	"github.com/gorilla/sessions"
)

var (
	Store sessions.Store
)

func Setup() {
	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}
