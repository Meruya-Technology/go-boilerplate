package datasources

import (
	"database/sql"
	"fmt"
	"strconv"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"go.uber.org/zap"
)

type UserDatasourceImpl struct {
	Config   cfg.Config
	Database sql.DB
}

func (i UserDatasourceImpl) User() (string, error) {
	logger, _ := zap.NewProduction()

	/// Open connection
	session := i.Database

	/// Functional code
	var result int
	session.QueryRow(`SELECT count(*) FROM user;`).Scan(&result)
	fmt.Println("The result is : ", result)
	logger.Info(strconv.Itoa(result))

	/// Close connection
	session.Close()
	return "Here is your user", nil
}
