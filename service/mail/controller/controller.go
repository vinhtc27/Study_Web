package controller

import (
	"encoding/json"
	"net"
	"net/http"
	"net/mail"
	"strings"
	"web-service/constant"
	"web-service/pkg/auth"
	"web-service/pkg/log"
	"web-service/pkg/router"
	"web-service/service/mail/model"

	"github.com/google/uuid"
)

// Request Body API Update Connection
type ConfirmEmailForm struct {
	Email string `json:"email"`
}

type ForgotPasswordForm struct {
	Email string `json:"email"`
}

func ValidateAndSendEmail(w http.ResponseWriter, r *http.Request) {
	var emailForm = &ConfirmEmailForm{}
	err := json.NewDecoder(r.Body).Decode(emailForm)
	if err != nil {
		log.Println(log.LogLevelError, "ValidateEmail: Decode(&emailForm)", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	// First it checks for email address format.
	_, err = mail.ParseAddress(emailForm.Email)
	if err != nil {
		log.Println(log.LogLevelError, "ValidateEmail: ParseAddress", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	//Then make sure that domain name is valid.
	emailPart := strings.SplitN(emailForm.Email, "@", 2)
	emailHost := emailPart[1]
	mx, err := net.LookupMX(emailHost)
	if err != nil {
		log.Println(log.LogLevelError, "ValidateEmail: LookupMX", err)
		router.ResponseBadRequest(w, "B.MAI.400.C2", err.Error())
		return
	}
	if len(mx) == 0 {
		router.ResponseBadRequest(w, "B.MAI.400.C2", "Email is invalid")
		return
	}

	//Sending a message to email make sure the mailbox really exists
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var payload = &auth.Payload{
		Email: emailForm.Email,
		Uuid:  uuid,
	}

	tokenData, err := payload.GetTokenDataJWT()
	if err != nil {
		log.Println(log.LogLevelDebug, "ValidateEmail: GetJWTToken", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	templateData := &struct {
		URL string
	}{
		URL: constant.CLIENT_BASE_URL + constant.CONFIRM_EMAIL_PATH + "?uuid=" + payload.Uuid + "&token=" + tokenData.Type + " " + tokenData.Token,
	}

	var email = model.CreateEmail(
		[]string{emailForm.Email},
		"[CONFIRM YOUR EMAIL]",
	)

	err = email.ParseTemplate("./template/confirm_email.html", templateData)
	if err != nil {
		log.Println(log.LogLevelDebug, "ValidateEmail: ParseTemplate", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	err = email.SendEmail()
	if err != nil {
		log.Println(log.LogLevelDebug, "ValidateEmail: SendEmail", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseSuccess(w, "B.MAI.200.C1", "Your email is valid.")
}

func ForgotPasswordAndSendEmail(w http.ResponseWriter, r *http.Request) {
	var forgotPasswordForm = &ForgotPasswordForm{}
	err := json.NewDecoder(r.Body).Decode(forgotPasswordForm)
	if err != nil {
		log.Println(log.LogLevelError, "ForgotPasswordAndSendEmail: Decode(&forgotPasswordForm)", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	// First it checks for email address format.
	_, err = mail.ParseAddress(forgotPasswordForm.Email)
	if err != nil {
		log.Println(log.LogLevelError, "ForgotPasswordAndSendEmail: ParseAddress", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	//Then make sure that domain name is valid.
	emailPart := strings.SplitN(forgotPasswordForm.Email, "@", 2)
	emailHost := emailPart[1]
	mx, err := net.LookupMX(emailHost)
	if err != nil {
		log.Println(log.LogLevelError, "ForgotPasswordAndSendEmail: LookupMX", err)
		router.ResponseBadRequest(w, "B.MAI.200.C2", err.Error())
		return
	}
	if len(mx) == 0 {
		router.ResponseBadRequest(w, "B.MAI.200.C2", "Email is invalid")
		return
	}

	//Sending a message to email make sure the mailbox really exists
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var payload = &auth.Payload{
		Email: forgotPasswordForm.Email,
		Uuid:  uuid,
	}

	tokenData, err := payload.GetTokenDataJWT()
	if err != nil {
		log.Println(log.LogLevelDebug, "ForgotPasswordAndSendEmail: GetJWTToken", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	templateData := &struct {
		URL string
	}{
		URL: constant.CLIENT_BASE_URL + constant.FORGOT_PASSWORD_PATH + "?uuid=" + payload.Uuid + "&token=" + tokenData.Type + " " + tokenData.Token,
	}

	var email = model.CreateEmail(
		[]string{forgotPasswordForm.Email},
		"[RESET YOUR PASSWORD]",
	)

	err = email.ParseTemplate("./template/forgot_pass.html", templateData)
	if err != nil {
		log.Println(log.LogLevelDebug, "ForgotPasswordAndSendEmail: ParseTemplate", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	err = email.SendEmail()
	if err != nil {
		log.Println(log.LogLevelDebug, "ForgotPasswordAndSendEmail: SendEmail", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseSuccess(w, "B.MAI.200.C1", "Your email is valid.")
}
