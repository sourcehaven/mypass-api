package db

import "errors"

// NotFound is the error used when an object is not found in the database.
var NotFound = errors.New("user not found")
