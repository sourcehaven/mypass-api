package db

import (
	"github.com/sourcehaven/mypass-api/internal/validate"
	"gorm.io/gorm"
)

type Vault struct {
	*gorm.Model

	// Foreign key to the User table, joins the User with this Vault record
	UserID uint `gorm:"not null" json:"user_id"`

	// Fields assigned by client
	Username string `json:"username"`
	Password string `json:"password"`
	Title    string `json:"title"`
	Website  string `json:"website"`
	Notes    string `json:"notes"`
	Folder   string `json:"folder"`

	// Active   bool   `json:"active"`

	// Self-Referential Has One
	// For now this is not used. No history is kept.
	// TODO: Optionally create an "analytical" database, where history is kept using slowly changing dimensions.
	// 		 For now let's only have a transactional database, which follows the basic CRUD operations.
	// ParentID *uint
	// Parent   *Vault

	// Many-to-many with vault_tags join table.
	Tags []*Tag `gorm:"many2many:vault_tags;"`
}

func (o *Vault) TableName() string {
	return "Vault"
}

func (o *Vault) Validate() error {
	if err := validate.PositiveInt(o.ID); err != nil {
		return err
	}
	if err := validate.PositiveInt(o.UserID); err != nil {
		return err
	}
	return nil
}
