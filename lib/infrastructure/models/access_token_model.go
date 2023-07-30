package models

import . "time"

type AccessTokenModel struct {
	Id        int
	CreatedAt Time
	UserId    int
	ClientId  int
	Token     string
	ExpiredAt Time
	DeviceId  string
	IsRevoked bool
}
