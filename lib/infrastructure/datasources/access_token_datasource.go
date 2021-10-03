package datasources

type AccessTokenDatasource interface {
	Create() (bool, error) /// Need to return userId also
	Check() (bool, error)  /// Need to return userId also
	Revoke() (bool, error) /// Need to return userId also
}
