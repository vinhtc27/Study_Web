package account

import (
	"web-service/pkg/auth"
	"web-service/service/account/controller"

	"github.com/go-chi/chi"
)

var AccountSubRoute = chi.NewRouter()

// Init package with sub-router for account service
func init() {
	AccountSubRoute.Group(func(_ chi.Router) {
		AccountSubRoute.Post("/signup", controller.CreateAccount)
		AccountSubRoute.Post("/signin", controller.Signin)
		AccountSubRoute.With(auth.JWT).Get("/confirm-email/uuid={uuid}", controller.ConfirmEmail)
		AccountSubRoute.With(auth.JWT).Patch("/forgot-password/uuid={uuid}", controller.ForgotPassword)
		AccountSubRoute.With(auth.JWT).Patch("/change-password", controller.ResetPassword)
		AccountSubRoute.With(auth.JWT).Get("/profile/current-profile", controller.GetCurrentProfile)
		AccountSubRoute.With(auth.JWT).Patch("/profile/current-profile", controller.UpdateCurrentProfile)
	})
}
