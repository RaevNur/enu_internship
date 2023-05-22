package routes

import (
	"net/http"
	"time"

	"enu_internship/internal/helper"
	"enu_internship/internal/models"
)

func GetUserBySession(route *Routes, w http.ResponseWriter, r *http.Request) *models.User {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}

	session, err := route.service.Session.GetByUuid(cookie.Value)
	if err != nil || session == nil {
		DeleteSessionCookie(w, r)
		return nil
	} else if session.CreatedAt.Add(helper.CookieExpireTime).Before(time.Now()) {
		DeleteSessionCookie(w, r)
		route.service.Session.DeleteSession(session.Id)
		return nil
	}

	user, err := route.service.User.GetByID(session.UserId)
	if err != nil || user == nil {
		DeleteSessionCookie(w, r)
		route.service.Session.DeleteSession(session.Id)
		return nil
	}

	return user
}

// sets the session cookie
func SetSessionCookie(uuid string, w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    uuid,
		MaxAge:   int(helper.CookieExpireTime.Seconds()),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
}

// deletes the session cookie
func DeleteSessionCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return
	}

	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}
