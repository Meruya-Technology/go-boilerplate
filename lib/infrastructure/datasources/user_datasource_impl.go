package datasources

import (
	"fmt"

	c "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/database"
)

type UserDatasourceImpl struct{}

func (impl UserDatasourceImpl) User() (string, error) {
	var cfg c.Config
	cfg = cfg.LoadConfig()

	/// Open connection
	db := new(database.PostgresqlHandler)
	session, err := db.ConnectAndGet(cfg)

	/// Error handler
	if err != nil {
		fmt.Print(err.Error())
		return err.Error(), err
	}

	/// Close connection
	db.Close(*session)
	return "Here is your user", nil
}
