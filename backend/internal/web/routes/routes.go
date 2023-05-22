package routes

import (
	"encoding/json"
	"net/http"

	"enu_internship/internal/service"
)

type Routes struct {
	service *service.Service
}

func NewRoutes(service *service.Service) (*Routes, error) {
	r := &Routes{
		service: service,
	}

	return r, nil
}

func (r *Routes) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", r.SignUp)
	mux.HandleFunc("/login", r.SignIn)
	mux.HandleFunc("/logout", r.Logout)
}

func (r *Routes) respond(rw http.ResponseWriter, code int, data interface{}) {
	rw.WriteHeader(code)
	if data != nil {
		json.NewEncoder(rw).Encode(data)
	}
}
