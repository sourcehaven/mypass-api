package model

import "github.com/sourcehaven/mypass-api/internal/validate"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (o *Credentials) Validate() error {
	if err := validate.NonEmptyStr(o.Username); err != nil {
		return err
	}
	if err := validate.NonEmptyStr(o.Password); err != nil {
		return err
	}
	return nil
}
