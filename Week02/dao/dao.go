package dao

import (
	"database/sql"
	xerrors "github.com/pkg/errors"
)

// raise err
func UpdateModel() error {
	err := sql.ErrNoRows
	return xerrors.Wrap(err, "raise a sql err")
}
