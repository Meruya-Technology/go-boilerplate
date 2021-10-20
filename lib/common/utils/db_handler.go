package utils

import (
	ctx "context"
	"database/sql"
)

type DbHandler struct {
	db *sql.DB
}

func (s *DbHandler) BeginTx(ctx ctx.Context) (*sql.Tx, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s *DbHandler) RollbackTx(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (s *DbHandler) CommitTx(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
