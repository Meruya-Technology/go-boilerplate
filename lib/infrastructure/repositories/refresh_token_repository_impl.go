package repositories_impl

import (
	ctx "context"
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	res "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/responses"
)

type RefreshTokenRepositoryImpl struct {
	Config   cfg.Config
	Database *sql.DB
}

///
func (r RefreshTokenRepositoryImpl) Refresh(ctx ctx.Context, Token string) (*res.LoginResponse, error) {
	return nil, nil
}
