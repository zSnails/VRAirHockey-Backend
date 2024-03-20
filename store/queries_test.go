package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/zSnails/VRAirHockey-Backend/db"
)

func init() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
}

func TestCreatePlayer(t *testing.T) {
	DB := db.Get()
	queries := New(DB)
	created, err := queries.CreatePlayer(context.Background(), CreatePlayerParams{
		Name:  "Fuap",
		Email: "erizojuan33@gmail.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("created: %v\n", created)

}
