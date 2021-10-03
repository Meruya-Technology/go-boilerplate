package datasources

type RefreshTokenDatasource interface {
	Create() /// Need to return userId also
	Check()  /// Need to return userId also
	Revoke() /// Need to return userId also
}
