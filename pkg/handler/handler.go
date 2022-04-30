package handler

import (
	"Run_Hse_Run/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() chi.Router {
	router := chi.NewRouter()
	// fix if I will use websocket
	router.Use(middleware.Logger, middleware.Timeout(time.Minute))

	router.Route("/auth", func(router chi.Router) {
		router.Post("/send-email", h.sendEmail)
		router.Post("/check-auth", h.checkAuth)
	})

	router.Route("/api", func(router chi.Router) {
		router.Use(h.authorizationOnly)
	})

	return router
}
