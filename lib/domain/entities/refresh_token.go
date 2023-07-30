package entities

import . "time"

type RefreshToken struct {
	Id            int    `json:"id" example:"1"`
	CreatedAt     Time   `json:"createdAt" example:"2021-11-23 14:09:57.113282+07"`
	AccessTokenId int    `json:"accessTokenId" example:"082115100216"`
	Token         string `json:"token" example:""`
	ExpiredAt     Time   `json:"expiredAt" example:"2021-11-23 14:09:57.113282+07"`
	IsRevoked     bool   `json:"isRevoked" example:"false"`
}
