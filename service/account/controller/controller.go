package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"web-service/constant"
	"web-service/pkg/auth"
	"web-service/pkg/cache"
	"web-service/pkg/log"
	"web-service/pkg/router"
	"web-service/pkg/utils"
	"web-service/service/account/model"

	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
)

// Status account enum
const (
	_defaultLoginKeyPrefix       = "login"
	_defaultPasswordGenerateCost = 12

	_maxFailSigninCount = 3
	_defaultSigninTTLs  = 1 * time.Hour
)

type AuthenticationForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordForm struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type ForgotPasswordForm struct {
	NewPassword string `json:"newPassword"`
}

type TokenData struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

type FailSignin struct {
	Email         string   `json:"email"`
	Counter       int      `json:"counter"`
	FailTimestamp []string `json:"failTimestamp"`
}

// Create new account and write to the database
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	//Form for register
	var registerForm = &AuthenticationForm{}
	_ = json.NewDecoder(r.Body).Decode(registerForm)

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerForm.Password), _defaultPasswordGenerateCost)
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: GenerateFromPassword", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	// Split name from gmail
	nameFromEmail := strings.SplitN(string(registerForm.Email), "@", 2)[0]

	var user = &model.User{
		Name:  nameFromEmail,
		Email: registerForm.Email,
	}

	err = user.GetUserByEmail()
	if user.Id != 0 {
		log.Println(log.LogLevelDebug, "CreateAccount: Email already registered", err)
		router.ResponseBadRequest(w, "B.ACC.400.C2", " Email already registered")
		return
	}

	err = user.InsertUser()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: InsertUser", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var account = &model.Account{
		UserId:        user.Id,
		Email:         registerForm.Email,
		Password:      string(hashedPassword),
		AccountType:   constant.ACCOUNT_TYPE_NORMAL,
		AccountStatus: constant.ACCOUNT_STATUS_INACTIVATED,
	}

	err = account.InsertAccount()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: InsertAccount", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	err = account.ValidateAndSendEmail()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: SendEmailVerifyAccount", err)
		router.ResponseBadRequest(w, "B.MAI.400.C2", err.Error())
		return
	}

	router.ResponseSuccess(w, "B.ACC.201.C1", "Check your email to confirm your account")
}

func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ConfirmEmail: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "ConfirmEmail: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	uuid := chi.URLParam(r, "uuid")
	if payload.Uuid != uuid {
		log.Println(log.LogLevelInfo, "ConfirmEmail: JWT.uuid mismatch URL.uuid", "")
		router.ResponseBadRequest(w, "B.ACC.401.C8", "This link is invalid.")
		return
	}

	var account = &model.Account{
		Email: payload.Email,
	}

	err = account.GetAccountByEmail()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ConfirmEmail: GetAccountByEmail", err)
		return
	}

	//Check if email already activated or deactivated
	if account.AccountStatus != constant.ACCOUNT_STATUS_INACTIVATED {
		if account.AccountStatus == constant.ACCOUNT_STATUS_ACTIVATED {
			log.Println(log.LogLevelInfo, "ConfirmEmail: Already activated", "")
			router.ResponseBadRequest(w, "B.ACC.400.C9", "Account is already activated")
		} else {
			log.Println(log.LogLevelInfo, "ConfirmEmail: Account has been deactivated", "")
			router.ResponseBadRequest(w, "B.ACC.400.C5", "Account has been deactivated")
		}
		return
	}

	//If account is inactivated -> activate account
	err = account.UpdateAccountStatusById(constant.ACCOUNT_STATUS_ACTIVATED)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ConfirmEmail: UpdateAccountStatusById", err)
		return
	}

	router.ResponseSuccess(w, "B.ACC.200.C7", "Confirm email successfully")
}

// UpdateProfileById Function to Update User's Profile by User ID
func UpdateCurrentProfile(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "UpdateCurrentProfile: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateCurrentProfile: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var user = &model.User{
		Id: payload.UserId,
	}

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	err = user.UpdateUser()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "UpdateCurrentProfile: UPDATE users Database", err)
		return
	}

	router.ResponseUpdated(w, "B.ACC.200.C12")
}

// GetProfileById Function to Get User's Profile by User ID
func GetCurrentProfile(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		log.Println(log.LogLevelDebug, "GetCurrentProfile: GetJWTClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "GetCurrentProfile: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var user = &model.User{
		Id: payload.UserId,
	}

	err = user.GetUserById()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "GetCurrentProfile: GetUserById", err)
		return
	}

	msg := fmt.Sprintf("Get profile user id %d sussess!", payload.UserId)

	router.ResponseSuccessWithData(w, "", msg, user)
}

// Signin function by checking the password and the hashedPassword
func Signin(w http.ResponseWriter, r *http.Request) {
	//Create login form
	var signinForm = &AuthenticationForm{}
	_ = json.NewDecoder(r.Body).Decode(signinForm)

	var account = &model.Account{
		Email: signinForm.Email,
	}

	err := account.GetAccountByEmail()
	if account.Id == 0 {
		log.Println(log.LogLevelDebug, "Signin: Email is not registered", err)
		router.ResponseBadRequest(w, "B.ACC.400.C3", "Wrong email or password")
		return
	}

	//Not allow inactivated or deactivated accounts
	if account.AccountStatus != constant.ACCOUNT_STATUS_ACTIVATED {
		if account.AccountStatus == constant.ACCOUNT_STATUS_INACTIVATED {
			log.Println(log.LogLevelInfo, "Signin: Account inactivated", "")
			router.ResponseBadRequest(w, "B.ACC.400.C4", "Account inactivated")
		} else {
			log.Println(log.LogLevelInfo, "Signin: Account has been deactivated", "")
			router.ResponseBadRequest(w, "B.ACC.400.C5", "Account has been deactivated")
		}
		return
	}

	//Compare the password with the hashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(signinForm.Password))
	if err != nil {
		cacheItemKey := _defaultLoginKeyPrefix + strconv.Itoa(account.Id)
		cacheItem, found := cache.LocalCache.Get(cacheItemKey)
		fmt.Println(cache.LocalCache.MetricsString())
		if !found {
			var failSignin = &FailSignin{
				Email:         account.Email,
				Counter:       1,
				FailTimestamp: []string{"", "", ""},
			}
			cache.LocalCache.SetByKey(cacheItemKey, failSignin, _defaultSigninTTLs)
			cacheItem = failSignin
		}

		if cacheItem.(*FailSignin).Counter < _maxFailSigninCount {
			cacheItem.(*FailSignin).FailTimestamp[cacheItem.(*FailSignin).Counter] = utils.Timestamp()
			cacheItem.(*FailSignin).Counter++
		} else {
			account.DeactivatedAccountByEmail()
			log.Println(log.LogLevelDebug, "Signin: DeactivatedAccountByEmail", err)
			router.ResponseBadRequest(w, "B.ACC.400.C5", "Account is deactivated")
			return
		}
		log.Println(log.LogLevelDebug, "Signin: CompareHashAndPassword", err)
		router.ResponseBadRequest(w, "B.ACC.400.C3", "Wrong email or password")
		return
	}

	var payload = &auth.Payload{
		UserId:    account.UserId,
		AccountId: account.Id,
	}

	tokenData, err := payload.GetTokenDataJWT()
	if err != nil {
		log.Println(log.LogLevelDebug, "Signin: GetJWTToken", err)
		router.ResponseInternalError(w, err.Error())
	}

	router.ResponseSuccessWithData(w, "B.ACC.200.C6", "Signin successfully", tokenData)
}

// Reset password of account
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	//Form for register
	var resetPasswordForm = &ResetPasswordForm{}
	_ = json.NewDecoder(r.Body).Decode(&resetPasswordForm)

	var account = &model.Account{
		Email: resetPasswordForm.Email,
	}

	err := account.GetAccountByEmail()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ResetPassword: GetAccount", err)
		return
	}

	//Compare the password with the hashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(resetPasswordForm.Password))
	if err != nil {
		router.ResponseBadRequest(w, "B.ACC.400.C10", "Old password is incorrect")
		log.Println(log.LogLevelDebug, "ResetPassword: CompareHashAndPassword", err)
		return
	}

	// Generate new hashed password
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(resetPasswordForm.NewPassword), _defaultPasswordGenerateCost)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ResetPassword: GenerateFromPassword", err)
		return
	}

	err = account.UpdateAccountPasswordById(string(newHashedPassword))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ResetPassword: UpdateAccountPassword", err)
		return
	}

	router.ResponseSuccess(w, "B.ACC.200.C11", "Password reset successfully")
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ForgotPassword: GetJWTClaims", err)
		return
	}

	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "ForgotPassword: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	uuid := chi.URLParam(r, "uuid")
	if payload.Uuid != uuid {
		log.Println(log.LogLevelInfo, "ForgotPassword: JWT.uuid mismatch URL.uuid", "")
		router.ResponseBadRequest(w, "B.ACC.401.C8", "This link is invalid.")
		return
	}

	// Form for forget password
	var forgotPasswordForm = &ForgotPasswordForm{}
	_ = json.NewDecoder(r.Body).Decode(&forgotPasswordForm)

	var account = &model.Account{
		Email: payload.Email,
	}

	err = account.GetAccountByEmail()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ForgotPassword: GetAccount", err)
		return
	}

	//Check if email already activated or deactivated
	if account.AccountStatus != constant.ACCOUNT_STATUS_ACTIVATED {
		if account.AccountStatus == constant.ACCOUNT_STATUS_INACTIVATED {
			log.Println(log.LogLevelInfo, "ForgotPassword: Your account is inactivated", "")
			router.ResponseBadRequest(w, "B.ACC.400.C4", "Your account is inactivated")
		} else {
			log.Println(log.LogLevelInfo, "ForgotPassword: Your account has been deactivated", "")
			router.ResponseBadRequest(w, "B.ACC.400.C5", "Your account has been deactivated")
		}
		return
	}

	// Generate new hashed password
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(forgotPasswordForm.NewPassword), _defaultPasswordGenerateCost)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ForgotPassword: GenerateFromPassword", err)
		return
	}

	err = account.UpdateAccountPasswordById(string(newHashedPassword))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ForgotPassword : UpdateAccountPassword", err)
		return
	}

	router.ResponseSuccess(w, "B.ACC.200.C13", "Password updated successfully")
}
