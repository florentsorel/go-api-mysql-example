package data

import (
	"database/sql"
	"errors"
)

var ErrorRecordNotFound = errors.New("record not found")

type Models struct {
	Actor ActorModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Actor: ActorModel{DB: db},
	}
}
