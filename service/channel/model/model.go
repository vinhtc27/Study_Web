package model

import (
	"encoding/json"
	"web-service/pkg/db"
	"web-service/pkg/utils"
	"web-service/service/account/model"
)

type Message struct {
	ChannelId int    `json:"channelId"`
	SenderId  int    `json:"senderId"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
type Member struct {
	UserId int    `json:"userId"`
	Role   string `json:"role"`
}

type Channel struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Avatar      string       `json:"avatar"`
	Members     []Member     `json:"members"`
	Messages    []Message    `json:"messages"`
	TaskColumns []TaskColumn `json:"taskColumns"`
	CreatedDate string       `json:"createdDate"`
	UpdatedDate string       `json:"updatedDate"`
}

type TaskColumn struct {
	Title            string `json:"title"`
	TaskColumnDetail any    `json:"taskColumnDetail"`
}

// Insert new message to db
func (channel *Channel) CreateChannel() error {
	query := `INSERT INTO channels (name, avatar, members, createdDate, updatedDate) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	membersJSON, err := json.Marshal(channel.Members)
	if err != nil {
		return err
	}
	channel.CreatedDate = utils.Timestamp()
	channel.UpdatedDate = utils.Timestamp()

	err = db.PSQL.QueryRow(query, channel.Name, channel.Avatar, membersJSON, channel.CreatedDate, channel.UpdatedDate).Scan(&channel.Id)
	if err != nil {
		return err
	}

	channelIdJSON, err := json.Marshal(&model.ChannelId{Id: channel.Id})
	if err != nil {
		return err
	}

	query = `UPDATE users SET channels = channels || $1::jsonb, updatedDate = $2 WHERE id = $3;`
	_, err = db.PSQL.Exec(query, channelIdJSON, utils.Timestamp(), channel.Members[0].UserId)

	return err
}

func (channel *Channel) UpdateChannelName() error {
	query := `UPDATE channels SET name = $1, updatedDate = $2 WHERE id= $3;`
	_, err := db.PSQL.Exec(query, channel.Name, utils.Timestamp(), channel.Id)
	return err
}

func (channel *Channel) UpdateChannelAvatar() error {
	query := `UPDATE channels SET avatar = $1, updatedDate = $2 WHERE id= $3;`
	_, err := db.PSQL.Exec(query, channel.Avatar, utils.Timestamp(), channel.Id)
	return err
}

func (channel *Channel) GetChannelById() error {
	membersJSON := []byte{}
	messagesJSON := []byte{}
	taskColumnsJSON := []byte{}

	query := `SELECT * from channels WHERE id = $1`
	err := db.PSQL.QueryRow(query, channel.Id).Scan(&channel.Id, &channel.Name, &channel.Avatar, &membersJSON, &messagesJSON, &taskColumnsJSON, &channel.CreatedDate, &channel.UpdatedDate)
	if err != nil {
		return err
	}

	err = json.Unmarshal(membersJSON, &channel.Members)
	if err != nil {
		return err
	}
	err = json.Unmarshal(messagesJSON, &channel.Messages)
	if err != nil {
		return err
	}
	err = json.Unmarshal(taskColumnsJSON, &channel.TaskColumns)
	if err != nil {
		return err
	}

	return nil
}

func (channel *Channel) GetChannelHostId() (int, error) {
	query := `SELECT members->0->>'userId' FROM channels WHERE id = $1;`
	var hostId int
	err := db.PSQL.QueryRow(query, channel.Id).Scan(&hostId)
	if err != nil {
		return 0, err
	}
	return hostId, nil
}

func (channel *Channel) UpdateChannelById() error {
	query := `Update channels SET name = $1 WHERE id = $2;`
	_, err := db.PSQL.Exec(query, channel.Name, channel.Id)
	return err
}

func (channel *Channel) DeleteChannelById() error {
	query := `DELETE FROM channels WHERE id = $1;`
	_, err := db.PSQL.Exec(query, channel.Id)
	return err
}

func (channel *Channel) AddNewMember(newMember *Member) error {
	newMemberJSON, err := json.Marshal(newMember)
	if err != nil {
		return err
	}

	query := `UPDATE channels SET members = members || $1::jsonb, updatedDate = $2 WHERE id = $3;`
	_, err = db.PSQL.Exec(query, newMemberJSON, utils.Timestamp(), channel.Id)
	if err != nil {
		return err
	}

	channelIdJSON, err := json.Marshal(&model.ChannelId{Id: channel.Id})
	if err != nil {
		return err
	}

	query = `UPDATE users SET channels = channels || $1::jsonb, updatedDate = $2 WHERE id = $3;`
	_, err = db.PSQL.Exec(query, channelIdJSON, utils.Timestamp(), newMember.UserId)
	return err
}

func (channel *Channel) DeleteMember(member *Member) error {
	query := `UPDATE users SET channels = (
			SELECT jsonb_agg(elems) FROM users,
    		jsonb_array_elements(channels::jsonb) AS elems 
			WHERE elems ->> 'id' <> $1 AND id = $2)
			WHERE id = $2;`
	_, err := db.PSQL.Exec(query, channel.Id, member.UserId)
	if err != nil {
		return err
	}

	query = `UPDATE channels SET members = (
			SELECT jsonb_agg(elems) FROM channels,
    		jsonb_array_elements(members::jsonb) AS elems 
			WHERE elems ->> 'userId' <> $1 AND id = $2)
			WHERE id = $2;`

	_, err = db.PSQL.Exec(query, member.UserId, channel.Id)
	return err
}

func (channel *Channel) AddMessage(newMessage *Message) error {
	newMessageJSON, err := json.Marshal(newMessage)
	if err != nil {
		return err
	}

	query := `UPDATE channels SET messages = messages || $1::jsonb, updatedDate = $2 WHERE id = $3;`
	_, err = db.PSQL.Exec(query, newMessageJSON, utils.Timestamp(), channel.Id)

	return err
}

func (channel *Channel) AddTaskColumn(taskColumn *TaskColumn) error {
	newTaskColumnJSON, err := json.Marshal(taskColumn)
	if err != nil {
		return err
	}

	query := `UPDATE channels SET taskcolumns = taskcolumns || $1::jsonb, updatedDate = $2 WHERE id = $3;`
	_, err = db.PSQL.Exec(query, newTaskColumnJSON, utils.Timestamp(), channel.Id)

	return err
}

func (channel *Channel) DeleteTaskColumnByTitle(taskColumnTitle string) error {
	query := `UPDATE channels SET taskcolumns = (
			SELECT jsonb_agg(elems) FROM channels,
    		jsonb_array_elements(taskcolumns::jsonb) AS elems
			WHERE elems ->> 'title' <> $1 AND id = $2)
			WHERE id = $2;`

	_, err := db.PSQL.Exec(query, taskColumnTitle, channel.Id)
	return err
}
