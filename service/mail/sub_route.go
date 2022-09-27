package mail

import (
	"web-service/service/mail/controller"

	"github.com/go-chi/chi"
)

var MailSubRoute = chi.NewRouter()

// Init package with sub-router for mails service
func init() {
	MailSubRoute.Group(func(_ chi.Router) {
		MailSubRoute.Post("/validate-email", controller.ValidateAndSendEmail)
		MailSubRoute.Post("/forgot-password", controller.ForgotPasswordAndSendEmail)
	})
}
