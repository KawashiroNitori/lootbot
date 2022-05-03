package router

import (
	"github.com/KawashiroNitori/lootbot/internal/http/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	lootUploadHandler := handler.NewLootUploadHandler()

	r.Post("/webhook/loot", lootUploadHandler.Post)

	return r
}
