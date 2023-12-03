package model

import (
	"github.com/rs/zerolog"
)

type Env string

type Cfg struct {
	Host               string
	Port               string
	SecretKey          string
	Environment        Env
	LogLevel           zerolog.Level
	DbConnectionUri    string
	PasswordLength     uint64
	PasswordMinNumber  uint64
	PasswordMinCapital uint64
	PasswordMinSpecial uint64
}
