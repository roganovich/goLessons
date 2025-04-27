package main

import (
	"fmt"
	"go-lessons/junior-rest-api/internal/config"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // Путь к файлу .env

	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err.Error()))
	}

	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)

	log.Info(
		fmt.Sprintf("app: %s", cfg.AppName),
		slog.String("env", cfg.Env),
		slog.String("server", cfg.HTTPServer.Address),
	)

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
