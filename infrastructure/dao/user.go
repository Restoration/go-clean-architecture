package dao

import (
	"encoding/json"
	"go-clean-app/domain"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;unique;autoIncrement;not null;"`
	Name      string    `json:"name" gorm:"type:varchar(60)"`
	CreatedAt time.Time `json:"createdAt" gorm:"<-:false"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"<-:false"`
}

func (m *User) ToEntity() (*domain.User, error) {
	var entityUser domain.User
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &entityUser); err != nil {
		return nil, err
	}
	return &entityUser, nil
}

type Users []*User

func (m *Users) ToEntity() (domain.Users, error) {
	var entityUsers domain.Users
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &entityUsers); err != nil {
		return nil, err
	}
	return entityUsers, nil
}
