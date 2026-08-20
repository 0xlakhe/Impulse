package app

import (
	"context"

	"github.com/0xlakhe/Impluse/internal/config"
	"github.com/0xlakhe/Impluse/internal/database"
)


type App struct{
	Config *config.Config
	DB *database.Database
}

func New(cfg *config.Config, db *database.Database) *App{
	return &App{
		Config: cfg,
		DB: db,
	}
}

func (a *App) Shutdown(ctx context.Context) error{
	a.DB.Pool.Close()
	return nil
}