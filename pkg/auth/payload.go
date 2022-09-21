package auth

import (
	"encoding/json"
)

type Payload struct {
	UserId    int `json:"userId"`
	AccountId int `json:"accountId"`
}

func (payload *Payload) GetDataFromClaims(claims string) error {
	err := json.Unmarshal([]byte(claims), &payload)
	if err != nil {
		return err
	}
	return nil
}

func (payload *Payload) GetTokenDataJWT() (any, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	token, err := GetJWTToken(string(data))
	if err != nil {
		return nil, err
	}

	tokenData := struct {
		Token string `json:"token"`
		Type  string `json:"type"`
	}{
		Token: token,
		Type:  "Bearer",
	}

	return tokenData, nil
}
