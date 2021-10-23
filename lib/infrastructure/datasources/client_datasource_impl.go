package datasources

import (
	ctx "context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type ClientDatasourceImpl struct {
	Config   cfg.Config
	Database sql.DB
}

/// Refactor guidelines fo more reusable datasources
/// The code structure should have 3 blocks
/// 1. Begin Transactional
///    - Start Transaction
/// 2. Functional block
///    - Prepare context
///    - Run / execute queries
///	   - Scan & Get result if any
/// 3. Finishing Transactional
///    - Commit
/// 1 & 3 its on the outter layer
/// Meanwhile 2 is inside the datasource layer

func (i ClientDatasourceImpl) Create(ctx ctx.Context, Name string, Secret string) (*mdl.ClientModel, error) {

	/// Open connection
	session := i.Database
	jwtHandler := new(sec.JwtHandler)

	/// Functional code
	SecretKey := jwtHandler.Generate(Secret)
	var Id int
	const createClient = `INSERT INTO authentication.client (created_at, name, secret) VALUES ($1, $2, $3) RETURNING id`

	dbTx, err := session.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := dbTx.PrepareContext(ctx, createClient)
	if err != nil {
		return nil, err
	}

	fmt.Println(Name, SecretKey)
	stmtContext, err := dbTx.StmtContext(ctx, stmt).QueryContext(ctx, time.Now(), Name, SecretKey)
	if err != nil {
		return nil, err
	}

	inserStatus := stmtContext.Next()
	if !inserStatus {
		return nil, errors.New("Failed to insert client")
	}

	err = stmtContext.Scan(&Id)
	if err != nil {
		return nil, err
	}

	err = dbTx.Commit()
	if err != nil {
		return nil, err
	}

	result := mdl.ClientModel{
		Id:     int(Id),
		Name:   Name,
		Secret: SecretKey,
	}
	return &result, nil
}
