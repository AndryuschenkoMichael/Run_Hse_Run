package handler

import (
	"Run_Hse_Run/pkg/service"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
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
	//router.Use(middleware.Logger, middleware.Timeout(time.Minute))

	router.Route("/auth", func(router chi.Router) {
		router.Post("/send-email", h.sendEmail)
		router.Post("/check-auth", h.checkAuth)
		router.Post("/create-user", h.createUser)
	})

	router.Route("/api", func(router chi.Router) {
		router.Use(h.authorizationOnly)
	})

	router.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "%s", "pong")
		if err != nil {
			writer.WriteHeader(http.StatusOK)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}

	})

	return router
}
