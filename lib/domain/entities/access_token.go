package entities

import . "time"

type AccessToken struct {
	Id        int
	CreatedAt Time
	UserId    int
	ClientId  int
	Token     string
	ExpiredAt Time
	DeviceId  string
	IsRevoked bool
}
