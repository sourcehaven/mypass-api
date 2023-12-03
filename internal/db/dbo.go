package db

import "github.com/sourcehaven/mypass-api/internal/model"

type DbObject interface {
	User | Tag | Vault | model.Credentials
}

type DbObjects interface {
	[]User | []Tag | []Vault | []model.Credentials
}
