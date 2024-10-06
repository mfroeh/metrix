package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("data: no record found")
)

type Models struct {
	Summoners SummonerModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Summoners: SummonerModel{DB: db},
	}
}
