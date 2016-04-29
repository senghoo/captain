package models

import "time"

type User struct {
	ID        int64
	LowerName string `xorm:"UNIQUE NOT NULL"`
	Name      string `xorm:"UNIQUE NOT NULL"`
	FullName  string
	Email     string    `xorm:"NOT NULL"`
	Passwd    string    `xorm:"NOT NULL"`
	Created   time.Time `xorm:"CREATED"`
	Updated   time.Time `xorm:"UPDATED"`

	IsActive bool
	IsAdmin  bool
}
