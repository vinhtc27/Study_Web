package model

import (
	"errors"
	"net"
	"strings"
	"web-service/constant"
	"web-service/pkg/auth"
	"web-service/pkg/db"
	"web-service/pkg/utils"
	"web-service/service/mail/model"

	"github.com/google/uuid"
)

// Account Struct
type Account struct {
	Id            int    `json:"id"`
	UserId        int    `json:"userId"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	AccountType   string `json:"accountType"`
	AccountStatus string `json:"accountStatus"`
	CreatedDate   string `json:"createdDate"`
	UpdatedDate   string `json:"updatedDate"`
}

func (account *Account) InsertAccount() error {
	query := `INSERT INTO public.accounts (userId, email, password,  accountType, accountStatus, createdDate, updatedDate) VALUES ( $1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	err := db.PSQL.QueryRow(query, account.UserId, account.Email, account.Password, account.AccountType, account.AccountStatus, utils.Timestamp(), utils.Timestamp()).Scan(&account.Id)
	return err
}

func (account *Account) GetAccountByEmail() error {
	query := `SELECT * FROM accounts WHERE email = $1`

	err := db.PSQL.QueryRow(query, account.Email).Scan(&account.Id, &account.UserId, &account.Email, &account.Password, &account.AccountType, &account.AccountStatus, &account.CreatedDate, &account.UpdatedDate)

	return err
}

func (account *Account) GetAccountById() error {
	query := `SELECT * FROM accounts WHERE id = $1`
	err := db.PSQL.QueryRow(query, account.Id).Scan(&account.Id, &account.UserId, &account.Email, &account.Password, &account.AccountType, &account.AccountStatus, &account.CreatedDate, &account.UpdatedDate)

	return err
}

func (account *Account) UpdateAccountPasswordById(password string) error {
	query := `UPDATE accounts SET password = $1, updatedDate = $2 WHERE id= $3`

	_, err := db.PSQL.Exec(query, password, utils.Timestamp(), account.Id)

	return err
}

func (account *Account) UpdateAccountStatusById(accountStatus string) error {
	query := `UPDATE accounts SET accountStatus = $1, updatedDate = $2 WHERE id= $3`

	_, err := db.PSQL.Exec(query, accountStatus, utils.Timestamp(), account.Id)

	return err
}

func (account *Account) DeactivatedAccountByEmail() error {
	query := `UPDATE accounts SET accountStatus = $1, updatedDate = $2 WHERE id= $3`

	_, err := db.PSQL.Exec(query, constant.ACCOUNT_STATUS_DEACTIVATED, utils.Timestamp(), account.Id)

	return err
}

func (account *Account) ValidateAndSendEmail() error {
	//Then make sure that domain name is valid.
	emailPart := strings.SplitN(account.Email, "@", 2)
	emailHost := emailPart[1]
	mx, err := net.LookupMX(emailHost)
	if err != nil {
		return err
	}
	if len(mx) == 0 {
		return errors.New("email is invalid")
	}

	//Sending a message to email make sure the mailbox really exists
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var payload = &auth.Payload{
		Email: account.Email,
		Uuid:  uuid,
	}

	tokenData, err := payload.GetTokenDataJWT()
	if err != nil {
		return err
	}

	templateData := &struct {
		URL string
	}{
		URL: constant.CLIENT_BASE_URL + constant.CONFIRM_EMAIL_PATH + "?uuid=" + payload.Uuid + "&token=" + tokenData.Type + " " + tokenData.Token,
	}

	var email = model.CreateEmail(
		[]string{account.Email},
		"[CONFIRM YOUR EMAIL]",
	)

	err = email.ParseTemplate("./template/confirm_email.html", templateData)
	if err != nil {
		return err
	}

	err = email.SendEmail()
	if err != nil {
		return err
	}
	return nil
}
