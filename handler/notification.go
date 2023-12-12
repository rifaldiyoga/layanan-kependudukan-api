package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/user"

	firebase "firebase.google.com/go"
)

type notificationHandler struct {
	app         *firebase.App
	userService user.Service
	authService auth.Service
}

func NewNotificationHandler(app *firebase.App, userService user.Service, authService auth.Service) *notificationHandler {
	return &notificationHandler{app, userService, authService}
}
