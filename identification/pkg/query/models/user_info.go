package models

import "time"

type UserInfo struct {
	DisplayName string
	Name        string
	Email       string
	ChangeDate  time.Time
}
