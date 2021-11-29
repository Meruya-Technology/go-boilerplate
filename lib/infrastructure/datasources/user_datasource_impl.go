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

type UserDatasourceImpl struct {
	Config        cfg.Config
	Database      sql.DB
	DBTransaction *sql.Tx
	HashHandler   sec.HashHandler
}

func (i UserDatasourceImpl) User() (string, error) {
	// logger, _ := zap.NewProduction()

	// /// Open connection
	// session := i.Database

	// /// Functional code
	// var result int
	// session.QueryRow(`SELECT count(*) FROM user;`).Scan(&result)
	// fmt.Println("The result is : ", result)
	// logger.Info(strconv.Itoa(result))

	// /// Close connection
	// session.Close()
	return "Here is your user", nil
}

func (i UserDatasourceImpl) Login(ctx ctx.Context, Email string, Password string) (*mdl.UserModel, error) {

	/// Initialize transaction
	dbTx := i.DBTransaction

	/// Functional code
	var localId int
	var hashedPassword string
	const checkClient = `SELECT id, password FROM authentication.user WHERE email = $1`

	stmt, err := dbTx.PrepareContext(ctx, checkClient)
	if err != nil {
		return nil, err
	}

	sqlRow := stmt.QueryRowContext(ctx, Email)
	if sqlRow == nil {
		return nil, err
	}

	err = sqlRow.Scan(&localId, &hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("E-mail not registered")
	}

	passwordMatch := i.HashHandler.CheckPasswordHash(Password, hashedPassword)
	if !passwordMatch {
		return nil, fmt.Errorf("Invalid password")
	}

	result := mdl.UserModel{
		Id: int(localId),
	}
	return &result, nil
}

func (i UserDatasourceImpl) Register(ctx ctx.Context, Request mdl.RegisterRequestModel) (*int, error) {

	/// Declare connection
	dbTx := i.DBTransaction
	bytes, err := i.HashHandler.HashPassword(Request.Password)
	if err != nil {
		return nil, err
	}

	var Id int
	const createClient = `INSERT INTO authentication.user (created_at, first_name, last_name, phone, email, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	stmt, err := dbTx.PrepareContext(ctx, createClient)
	if err != nil {
		return nil, err
	}

	stmtContext, err := dbTx.StmtContext(ctx, stmt).QueryContext(ctx, time.Now(), Request.Firstname, Request.Lastname, Request.Phone, Request.Email, bytes)
	if err != nil {
		return nil, err
	}

	insertStatus := stmtContext.Next()
	if !insertStatus {
		return nil, errors.New("Failed to register")
	}

	err = stmtContext.Scan(&Id)
	if err != nil {
		return nil, err
	}

	err = stmtContext.Close()
	if err != nil {
		return nil, err
	}

	return &Id, nil
}

func (i UserDatasourceImpl) CheckPhone(ctx ctx.Context, Phone string) error {

	/// Initialize transaction
	db := i.DBTransaction
	var localId int
	const checkClient = `SELECT count(*) FROM authentication.user WHERE phone = $1`

	stmt, err := db.PrepareContext(ctx, checkClient)
	if err != nil {
		return err
	}

	sqlRow := stmt.QueryRowContext(ctx, Phone)
	if sqlRow == nil {
		return err
	}

	err = sqlRow.Scan(&localId)
	if err != nil {
		return fmt.Errorf("Internal Error")
	}

	if localId != 0 {
		return fmt.Errorf("Phone number already used ")
	}
	return nil
}

func (i UserDatasourceImpl) CheckEmail(ctx ctx.Context, Email string) error {
	/// Initialize transaction
	db := i.DBTransaction
	var localId int
	const checkClient = `SELECT count(*) FROM authentication.user WHERE email = $1`

	stmt, err := db.PrepareContext(ctx, checkClient)
	if err != nil {
		return err
	}

	sqlRow := stmt.QueryRowContext(ctx, Email)
	if sqlRow == nil {
		return err
	}

	err = sqlRow.Scan(&localId)
	if err != nil {
		return fmt.Errorf("Internal Error")
	}

	if localId != 0 {
		return fmt.Errorf("E-mail already used")
	}
	return nil
}
