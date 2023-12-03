package db

import (
	"gorm.io/gorm"
)

type TransactDb struct {
	*gorm.DB
}

func (db *TransactDb) Init() error {
	return db.AutoMigrate(
		&User{},
		&Tag{},
		&Vault{},
	)
}

func (db *TransactDb) GetUserById(id uint) (*User, error) {
	user := &User{}
	tx := db.First(user, id)
	return user, tx.Error
}

func (db *TransactDb) GetUser(username string) (*User, error) {
	user := &User{}
	tx := db.Where("username = ?", username).First(user)
	return user, tx.Error
}

func (db *TransactDb) CreateUser(user *User) (*User, error) {
	tx := db.Create(user)
	return user, tx.Error
}

func (db *TransactDb) UpdateUser(user *User) error {
	tx := db.Save(user)
	return tx.Error
}

func (db *TransactDb) DeleteUser(id uint) error {
	tx := db.Delete(&User{}, id)
	return tx.Error
}
