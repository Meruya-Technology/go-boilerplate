package requests

type CreateClientRequest struct {
	Name   string `json:"name" example:"Mobile client"`
	Secret string `json:"secret" example:"6wTqKFJ1c3QTJ3dkQ8fsKg=="`
}
