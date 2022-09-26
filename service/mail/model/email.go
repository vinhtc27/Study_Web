package model

import (
	"bytes"
	"net/smtp"
	"text/template"
)

const (
	_host                = "smtp.gmail.com"
	_address             = "smtp.gmail.com:587"
	_senderEmail         = "20021478@vnu.edu.vn" // FROM
	_applicationPassword = "khqpqplyrqquevxq"    //PASSWORD
)

var _defaultAuth = smtp.PlainAuth("", _senderEmail, _applicationPassword, _host)

type Email struct {
	to      []string
	subject string
	body    string
}

func CreateEmail(to []string, subject string) *Email {
	return &Email{
		to:      to,
		subject: subject,
	}
}

func (email *Email) SendEmail() error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := "Subject:" + email.subject + "\n" + mime + "\n" + email.body
	err := smtp.SendMail(_address, _defaultAuth, _senderEmail, email.to, []byte(msg))
	return err
}

func (email *Email) ParseTemplate(templateFileName string, data interface{}) error {
	template, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = template.Execute(buf, data)
	if err != nil {
		return err
	}
	email.body = buf.String()
	return nil
}
