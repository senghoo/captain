package models

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/go-macaron/session"
	"github.com/senghoo/captain/modules/utils"
)

type User struct {
	ID        int64
	LowerName string `xorm:"UNIQUE NOT NULL"`
	Name      string `xorm:"UNIQUE NOT NULL"`
	FullName  string
	Email     string    `xorm:"NOT NULL"`
	Passwd    string    `xorm:"NOT NULL"`
	Salt      string    `xorm:"NOT NULL"`
	Created   time.Time `xorm:"CREATED"`
	Updated   time.Time `xorm:"UPDATED"`

	IsActive bool
	IsAdmin  bool
}

func NewUser(name string) *User {
	u := &User{
		LowerName: strings.ToLower(name),
		Name:      name,
	}
	u.UpdateSalt()
	return u
}

func GetUserByID(id int64) *User {
	u := &User{
		ID: id,
	}
	if existed, err := x.Get(u); err != nil && existed {
		return u
	}
	return nil
}

func GetUserFromSession(sess session.Store) *User {
	uid := sess.Get("uid")
	if uid == nil {
		return nil
	}

	if id, ok := uid.(int64); ok {
		return GetUserByID(id)
	}
	return nil
}

func (u *User) ValidatePassword(passwd string) bool {
	return u.Passwd == encodePassword(passwd, u.Salt)
}

func (u *User) SetPassword(passwd string) {
	if u.Salt == "" {
		u.UpdateSalt()
	}
	u.Passwd = encodePassword(passwd, u.Salt)
}

func (u *User) UpdateSalt() {
	u.Salt = utils.RandomString(10)
}

func (u *User) Save() {
	if u.ID == 0 {
		x.Insert(u)
	} else {
		x.Id(u.ID).Update(u)
	}
}

func encodePassword(pass, salt string) string {
	return fmt.Sprintf("%x", utils.PBKDF2([]byte(pass), []byte(salt), 10000, 50, sha256.New))
}

func UserSignIn(username, password string) (*User, error) {
	var u *User
	if strings.Contains(username, "@") {
		u = &User{Email: strings.ToLower(username)}
	} else {
		u = &User{LowerName: strings.ToLower(username)}
	}

	userExists, err := x.Get(u)
	if err != nil {
		return nil, err
	}

	if userExists {
		if u.ValidatePassword(password) {
			return u, nil
		}
	}
	return nil, ErrUserNotExist{u.ID, u.Name}
}

type ErrUserNotExist struct {
	UID  int64
	Name string
}

func IsErrUserNotExist(err error) bool {
	_, ok := err.(ErrUserNotExist)
	return ok
}

func (err ErrUserNotExist) Error() string {
	return fmt.Sprintf("user does not exist [uid: %d, name: %s]", err.UID, err.Name)
}
