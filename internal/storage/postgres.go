package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type Storage struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
}
