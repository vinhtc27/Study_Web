package model

import (
	"database/sql"
	"reflect"
	"web-service/pkg/db"
	"web-service/pkg/utils"
)

// User Struct
type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Dob         string `json:"dob"`
	Sex         string `json:"sex"`
	Avartar     string `json:"avartar"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	IdCard      string `json:"idCard"`
	National    string `json:"national"`
	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
}

// User Struct
type NullUser struct {
	Id          sql.NullInt32  `json:"id"`
	Name        sql.NullString `json:"name"`
	Dob         sql.NullString `json:"dob"`
	Sex         sql.NullString `json:"sex"`
	Avartar     sql.NullString `json:"avartar"`
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
	if reflect.TypeOf(nullUser.Avartar) != nil {
		user.Avartar = nullUser.Avartar.String
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
	query := `INSERT INTO users (name, email, createddate, updateddate) VALUES ( $1, $2, $3, $4) RETURNING id;`

	err := db.PSQL.QueryRow(query, user.Name, user.Email, utils.Timestamp(), utils.Timestamp()).Scan(&user.Id)
	return err
}

// Select user by id from database
func (user *User) GetUserByEmail() error {
	var nullUser = &NullUser{}
	query := `SELECT * FROM users WHERE email= $1`

	err := db.PSQL.QueryRow(query, user.Email).Scan(&nullUser.Id, &nullUser.Name, &nullUser.Dob, &nullUser.Sex, &nullUser.Avartar, &nullUser.Email, &nullUser.Address, &nullUser.Phone, &nullUser.IdCard, &nullUser.National, &nullUser.CreatedDate, &nullUser.UpdatedDate)
	if err != nil {
		return err
	}

	user.ConvertToUser(nullUser)

	return nil
}

// Select user by id from database
func (user *User) GetUserById() error {
	var nullUser = &NullUser{}
	query := `SELECT * FROM users WHERE id= $1`

	err := db.PSQL.QueryRow(query, user.Id).Scan(&nullUser.Id, &nullUser.Name, &nullUser.Dob, &nullUser.Sex, &nullUser.Avartar, &nullUser.Email, &nullUser.Address, &nullUser.Phone, &nullUser.IdCard, &nullUser.National, &nullUser.CreatedDate, &nullUser.UpdatedDate)
	if err != nil {
		return err
	}

	user.ConvertToUser(nullUser)

	return nil
}

// Update user from database
func (user *User) UpdateUser() error {
	query := `UPDATE users SET name= $1, dob= $2, sex= $3, address= $4, phone= $5, idCard= $6, national= $7, updatedDate = $8 WHERE id= $9`

	_, err := db.PSQL.Exec(query, user.Name, user.Dob, user.Sex, user.Address, user.Phone, user.IdCard, user.National, utils.Timestamp(), user.Id)
	return err
}
