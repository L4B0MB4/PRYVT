package query

import (
	"time"

	"github.com/google/uuid"
)

type UserInfo struct {
	ID           uuid.UUID
	DisplayName  string
	Name         string
	PasswordHash string
	Email        string
	ChangeDate   time.Time
}
