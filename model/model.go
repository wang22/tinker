package model

import (
	"time"
)

type ModelInterface interface {
	Insert()
}

type Model struct {
	ID       int `gorm:"primary_key"`
	Created  time.Time
	Modified time.Time
	Deleted  time.Time `gorm:"DEFAULT:null"`
}

func (model *Model) InsertHook() {
	model.Created = time.Now()
	model.Modified = time.Now()
}

func (model *Model) UpdateHook() {
	model.Modified = time.Now()
}

func (model *Model) DeleteHook() {
	model.Deleted = time.Now()
}

type Admin struct {
	Model
	Username string
	Password string
	Avatar   string
}
