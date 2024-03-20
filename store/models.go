// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package store

import ()

type Auth struct {
	ID          int64       `json:"id"`
	Hash        string      `json:"hash"`
	PlayerID    int64       `json:"playerId"`
	PlayerEmail string      `json:"playerEmail"`
	Foreign     interface{} `json:"foreign"`
}

type Player struct {
	ID    interface{} `json:"id"`
	Email string      `json:"email"`
	Name  string      `json:"name"`
}

type Score struct {
	ID       int64 `json:"id"`
	Score    int64 `json:"score"`
	PlayerID int64 `json:"playerId"`
}
