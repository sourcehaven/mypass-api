package db

import (
	"github.com/sourcehaven/mypass-api/internal/validate"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Username     string `gorm:"unique;not null" json:"username"`
	Email        string `gorm:"unique;not null;size:320" json:"email"`
	Password     string `gorm:"-" json:"password"`
	PasswordHash []byte `gorm:"not null"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`

	// Explicitly referring to the foreign key
	VaultEntries []*Vault `gorm:"foreignKey:UserID"`
}

func (o *User) Validate(minLength, minNumeric, minCapital, minSpecial uint64) error {
	if err := validate.PositiveInt(o.ID); err != nil {
		return err
	}
	if err := validate.NonEmptyStr(o.Username); err != nil {
		return err
	}
	if err := validate.IsEmailAddress(o.Email); err != nil {
		return err
	}
	if err := validate.IsStrongPassword(o.Password, minLength, minNumeric, minCapital, minSpecial); err != nil {
		return err
	}
	return nil
}
