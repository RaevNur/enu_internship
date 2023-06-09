package web

import (
	"fmt"
	"net/http"

	"enu_internship/internal/service"
	"enu_internship/internal/web/routes"
)

type MainHandler struct {
	routes *routes.Routes
}

func NewMainHandler(service *service.Service) (*MainHandler, error) {
	r, err := routes.NewRoutes(service)
	if err != nil {
		return nil, fmt.Errorf("MainHandler.NewMainHandler: %w", err)
	}

	return &MainHandler{
		routes: r,
	}, nil
}

func (m *MainHandler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	m.routes.InitRoutes(mux)
	return mux
}
