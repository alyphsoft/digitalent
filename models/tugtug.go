package models

import "gorm.io/gorm"

type Tugtug struct {
	gorm.Model

	ID       uint `gorm:"primaryKey; autoIncrement:true`
	Assignee string
	Task     string
	Deadline string
	IsFinish bool
}
