package routes

import (
	"github.com/sourcehaven/mypass-api/internal/db"
	"github.com/sourcehaven/mypass-api/internal/model"
)

type Handler struct {
	Store db.TransactDatastore
	Cfg   *model.Cfg
}
