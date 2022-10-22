package model

import (
	"web-service/pkg/db"
	"web-service/pkg/utils"
)

type Message struct {
	Id           int    `json:"id"`
	SenderId     int    `json:"senderId"`
	SenderName   string `json:"senderName"`
	Content      string `json:"content"`
	ChannelId    int    `json:"channelId"`
	CreatedDate  string `json:"createdDate"`
	LastModified string `json:"lastModified"`
	//Channel      model.Channel
	//Type         int    `json:"type"`
}

// Insert new message to db
func (message *Message) CreateMessage() error {
	query := `INSERT INTO messages (sender_id, sender_name, content, channel_id, created_date, last_modified) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	err := db.PSQL.QueryRow(query, message.SenderId, message.SenderName, message.Content, message.ChannelId, utils.Timestamp(), utils.Timestamp()).Scan(&message.Id)
	return err
}

func (message *Message) ModifyMessage() error {
	query := `UPDATE messages SET content= $1, last_modified = $2 WHERE id= $3`
	_, err := db.PSQL.Exec(query, message.Content, utils.Timestamp(), message.Id)
	return err
}

func (message *Message) DeleteMessage() error {
	query := `DELETE FROM messages WHERE id= $1 and $channelId= $2`
	_, err := db.PSQL.Exec(query, message.Id, message.ChannelId)
	return err
}
