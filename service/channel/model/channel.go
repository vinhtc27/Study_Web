package model

import (
	"web-service/pkg/db"
	"web-service/pkg/utils"
)

type Channel struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	MemberId    int    `json:"memberId"`
	MemberName  string `json:"memberName"`
	CreatedDate string `json:"createdDate"`
	Role        int    `json:"role"`
}

// Insert new message to db
func (channel *Channel) CreateChannel() error {
	query := `INSERT INTO channels (name, member_id, member_name, role, created_date) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err := db.PSQL.QueryRow(query, channel.Name, channel.MemberId, channel.MemberName, channel.Role, utils.Timestamp()).Scan(&channel.Id)
	return err
}

func (channel *Channel) DeleteChannel() error {
	query := `DELETE FROM channels where id = $1;`
	_, err := db.PSQL.Exec(query, channel.Id)
	return err
}
