package responses

import (
	bas "github.com/Meruya-Technology/go-boilerplate/lib/common/base"
	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type CreateClientResponse struct {
	Base bas.BaseResponse
	Data *ent.Client `json:"data"`
}
