package db

import (
	"github.com/sourcehaven/mypass-api/internal/validate"
	"gorm.io/gorm"
)

type Tag struct {
	*gorm.Model
	UserID      uint   `gorm:"uniqueIndex:idx_userid_name" json:"user_id"`
	Name        string `gorm:"uniqueIndex:idx_userid_name" json:"name"`
	Description string `json:"description"`
}

func (o *Tag) TableName() string {
	return "Tag"
}

func (o *Tag) Validate() error {
	if err := validate.PositiveInt(o.ID); err != nil {
		return err
	}
	if err := validate.PositiveInt(o.UserID); err != nil {
		return err
	}
	if err := validate.NonEmptyStr(o.Name); err != nil {
		return err
	}
	return nil
}
