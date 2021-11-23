package datasources

import (
	ctx "context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type UserDatasourceImpl struct {
	Config        cfg.Config
	DBTransaction *sql.Tx
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

func (i UserDatasourceImpl) Login(ctx ctx.Context, Username string, Password string) (*mdl.UserModel, error) {
	h := hmac.New(sha256.New, []byte(i.Config.Secret))

	// Write Data to it
	h.Write([]byte(Password))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("encypted: " + sha)
	result := mdl.UserModel{
		Id:    1,
		Email: "Dwi Kurnianto Mulyadien",
	}
	return &result, nil
}

func (i UserDatasourceImpl) Register(ctx ctx.Context, Request mdl.RegisterRequestModel) (*mdl.UserModel, error) {

	result := mdl.UserModel{
		Id:    1,
		Email: "Dwi Kurnianto Mulyadien",
	}
	return &result, nil
}

func (i UserDatasourceImpl) CheckPhone(ctx ctx.Context, Phone string) (bool, error) {
	return true, nil
}

func (i UserDatasourceImpl) CheckEmail(ctx ctx.Context, Email string) (bool, error) {
	return true, nil
}
