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
	//Messages    []model.Message
}

// Insert new message to db
func (channel *Channel) CreateChannel() error {
	query := `INSERT INTO channels (name, member_id, member_name, role, created_date) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err := db.PSQL.QueryRow(query, channel.Name, channel.MemberId, channel.MemberName, channel.Role, utils.Timestamp()).Scan(&channel.Id)
	return err
}

func (channel *Channel) AddMemberToChannel() error {
	query := `INSERT INTO channels (name, member_id, member_name, role, created_date) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err := db.PSQL.QueryRow(query, channel.Name, channel.MemberId, channel.MemberName, channel.Role, utils.Timestamp()).Scan(&channel.Id)
	return err
}

func (channel *Channel) DeleteChannel() error {
	query := `DELETE FROM channels where name = $1 and member_name = $2 and role = $3;`
	_, err := db.PSQL.Exec(query, channel.Name, channel.MemberName, channel.Role)
	return err
}

func (channel *Channel) GetAllChannelsOfHost() error {
	query := `SELECT * FROM channels where member_id = $1 and role = $2;`
	err := db.PSQL.QueryRow(query, channel.MemberId, channel.Role).Scan(&channel.Id, &channel.Name, &channel.MemberId, &channel.MemberName, &channel.CreatedDate, &channel.Role)
	if err != nil {
		return err
	}
	return nil
}

func (channel *Channel) ChangeNameChannel() error {
	query := `UPDATE channels SET name = $1 WHERE id= $2 and member_name = $3 and role = $4`
	_, err := db.PSQL.Exec(query, channel.Name, channel.Id, channel.MemberName, channel.Role)
	return err
}
