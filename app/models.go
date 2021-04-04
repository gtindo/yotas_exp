package app

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type AppUser struct {
	gorm.Model
	Name      string
	Email     string
	GithubId  string
	AvatarUrl string
	LastLogin time.Time
}

func (u *AppUser) Exists() bool {
	var user AppUser
	result := db.First(&user, "email = ?", u.Email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
