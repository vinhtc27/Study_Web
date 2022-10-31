package model

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"web-service/pkg/db"
	"web-service/pkg/utils"
)

// User Struct
type User struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Dob         string      `json:"dob"`
	Sex         string      `json:"sex"`
	Avatar      string      `json:"avatar"`
	Email       string      `json:"email"`
	Address     string      `json:"address"`
	Phone       string      `json:"phone"`
	IdCard      string      `json:"idCard"`
	National    string      `json:"national"`
	Channels    []ChannelId `json:"channels"`
	CreatedDate string      `json:"createdDate"`
	UpdatedDate string      `json:"updatedDate"`
}

type ChannelId struct {
	Id int `json:"id"`
}

// User Struct
type NullUser struct {
	Id          sql.NullInt32  `json:"id"`
	Name        sql.NullString `json:"name"`
	Dob         sql.NullString `json:"dob"`
	Sex         sql.NullString `json:"sex"`
	Avatar      sql.NullString `json:"avatar"`
	Email       sql.NullString `json:"email"`
	Address     sql.NullString `json:"address"`
	Phone       sql.NullString `json:"phone"`
	IdCard      sql.NullString `json:"idCard"`
	National    sql.NullString `json:"national"`
	CreatedDate sql.NullString `json:"createdDate"`
	UpdatedDate sql.NullString `json:"updatedDate"`
}

// Maping null value to the empty string
func (user *User) ConvertToUser(nullUser *NullUser) {
	if reflect.TypeOf(nullUser.Id) != nil {
		user.Id = int(nullUser.Id.Int32)
	}
	if reflect.TypeOf(nullUser.Name) != nil {
		user.Name = nullUser.Name.String
	}
	if reflect.TypeOf(nullUser.Dob) != nil {
		user.Dob = nullUser.Dob.String
	}
	if reflect.TypeOf(nullUser.Sex) != nil {
		user.Sex = nullUser.Sex.String
	}
	if reflect.TypeOf(nullUser.Avatar) != nil {
		user.Avatar = nullUser.Avatar.String
	}
	if reflect.TypeOf(nullUser.Email) != nil {
		user.Email = nullUser.Email.String
	}
	if reflect.TypeOf(nullUser.Address) != nil {
		user.Address = nullUser.Address.String
	}
	if reflect.TypeOf(nullUser.Phone) != nil {
		user.Phone = nullUser.Phone.String
	}
	if reflect.TypeOf(nullUser.IdCard) != nil {
		user.IdCard = nullUser.IdCard.String
	}
	if reflect.TypeOf(nullUser.National) != nil {
		user.National = nullUser.National.String
	}
	if reflect.TypeOf(nullUser.CreatedDate) != nil {
		user.CreatedDate = nullUser.CreatedDate.String
	}
	if reflect.TypeOf(nullUser.UpdatedDate) != nil {
		user.UpdatedDate = nullUser.UpdatedDate.String
	}
}

// Insert new user into the database
func (user *User) InsertUser() error {
	channelsJSON, err := json.Marshal(user.Channels)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (name, email, avatar, channels, createddate, updateddate) VALUES ( $1, $2, $3, $4, $5, $6) RETURNING id;`

	return db.PSQL.QueryRow(query, user.Name, user.Email, user.Avatar, channelsJSON, utils.Timestamp(), utils.Timestamp()).Scan(&user.Id)
}

func (user *User) UserIsExist() (bool, error) {
	var exist bool
	query := `SELECT exists(SELECT 1 FROM users WHERE email= $1);`
	err := db.PSQL.QueryRow(query, user.Email).Scan(&exist)
	return exist, err
}

// Select user by id from database
func (user *User) GetUserById() error {
	nullUser := &NullUser{}
	channels := []byte{}
	query := `SELECT * FROM users WHERE id= $1;`

	err := db.PSQL.QueryRow(query, user.Id).Scan(&nullUser.Id, &nullUser.Name, &nullUser.Dob, &nullUser.Sex, &nullUser.Avatar, &nullUser.Email, &nullUser.Address, &nullUser.Phone, &nullUser.IdCard, &nullUser.National, &channels, &nullUser.CreatedDate, &nullUser.UpdatedDate)
	if err != nil {
		return err
	}

	user.ConvertToUser(nullUser)
	err = json.Unmarshal(channels, &user.Channels)
	if err != nil {
		return err
	}

	return nil
}

// Update user from database
func (user *User) UpdateUser() error {
	query := `UPDATE users SET name= $1, dob= $2, sex= $3, avatar= $4, address= $5, phone= $6, idCard= $7, national= $8, updatedDate = $9 WHERE id= $10`

	_, err := db.PSQL.Exec(query, user.Name, user.Dob, user.Sex, user.Avatar, user.Address, user.Phone, user.IdCard, user.National, utils.Timestamp(), user.Id)
	return err
}
