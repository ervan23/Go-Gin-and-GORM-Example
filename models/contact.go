package models

import "time"

type Contact struct {
	Id        string     `json:"id" gorm:"auto_increment;primary_key"`
	Name      string     `json:"name" gorm:"varchar(100)"`
	Phone     string     `json:"phone" gorm:"varchar(15)"`
	Email     string     `json:"email" gorm:"varchar(100)"`
	Photo     string     `json:"photo" gorm:"varchar(255)"`
	CreatedAt *time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:timestamp"`
}

func (c Contact) GetPrimary() string {
	return c.Id
}
