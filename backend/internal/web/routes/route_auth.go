package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"enu_internship/internal/helper"
	"enu_internship/internal/models"
)

func (route *Routes) SignUp(w http.ResponseWriter, r *http.Request) {
	user := GetUserBySession(route, w, r)

	if user != nil {
		route.respond(w, http.StatusSeeOther, nil)
		return
	}

	user = &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		route.respond(w, http.StatusInternalServerError, nil)
		return
	}

	if r.Method != http.MethodPost {
		route.respond(w, http.StatusMethodNotAllowed, nil)
		return
	}

	err = route.service.User.Register(user)
	var validateError *helper.ValidateError
	var existsError *helper.ExistsError
	if err != nil {
		if errors.As(err, &validateError) {
			route.respond(w, http.StatusBadRequest, validateError.Error())
			return
		} else if errors.As(err, &existsError) {
			route.respond(w, http.StatusBadRequest, existsError.Error())
			return
		} else {
			route.respond(w, http.StatusInternalServerError, nil)
			return
		}
	}

	session, err := route.service.Session.GenerateSession(user.Id)
	if err != nil {
		route.respond(w, http.StatusInternalServerError, nil)
		return
	}

	SetSessionCookie(session.Uuid, w)
	route.respond(w, http.StatusCreated, nil)
}

func (route *Routes) SignIn(w http.ResponseWriter, r *http.Request) {
	user := GetUserBySession(route, w, r)

	if user != nil {
		route.respond(w, http.StatusSeeOther, nil)
		return
	}

	user = &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		route.respond(w, http.StatusInternalServerError, nil)
		return
	}

	if r.Method != http.MethodPost {
		route.respond(w, http.StatusMethodNotAllowed, nil)
		return
	}

	err = route.service.User.Login(user)
	var passErr *helper.ValidateError
	var existsErr *helper.ExistsError
	if errors.As(err, &passErr) {
		route.respond(w, http.StatusBadRequest, passErr.Error())
		return
	} else if errors.As(err, &existsErr) {
		route.respond(w, http.StatusBadRequest, existsErr.Error())
		return
	} else if err != nil {
		route.respond(w, http.StatusInternalServerError, nil)
		return
	}

	session, err := route.service.Session.GetByUserId(user.Id)
	if err != nil {
		route.respond(w, http.StatusInternalServerError, nil)
		return
	}

	if session != nil {
		err = route.service.Session.DeleteSession(session.Id)
		if err != nil {
			route.respond(w, http.StatusInternalServerError, nil)
			return
		}
	}

	session, err = route.service.Session.GenerateSession(user.Id)
	if err != nil {
		route.respond(w, http.StatusInternalServerError, nil)
		return
	}

	SetSessionCookie(session.Uuid, w)
	route.respond(w, http.StatusOK, nil)
}

func (route *Routes) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == nil {
		session, err := route.service.Session.GetByUuid(cookie.Value)
		if err != nil || session == nil {
			DeleteSessionCookie(w, r)
		} else {
			DeleteSessionCookie(w, r)
			route.service.Session.DeleteSession(session.Id)
		}
	}

	route.respond(w, http.StatusOK, nil)
}
