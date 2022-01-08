package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB,
	nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("ping", Pong)
	if err != nil {
		logger.Error("Unable to register \"ping\": %v", err)
		return err
	}
	return nil
}

func Pong(ctx context.Context, logger runtime.Logger, db *sql.DB,
	nk runtime.NakamaModule, payload string) (string, error) {
	return "pong", nil
}
