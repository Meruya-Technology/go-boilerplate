package usecases

import (
	"encoding/json"
	"net/http"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	imp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
)

type RetrieveProfile struct{}

// RetrieveUser example
// @Description An API for retrieve profile data
// @ID profile-get
// @tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} base.SuccessResponse{data=entities.User} "Success response"
// @Success 500 {object} base.InternalServerError "Internal Server Error"
// @Success 400 {object} base.BadRequestError "Bad Request"
// @Success 401 {object} base.UnauthorizedError "Unauthorized"
// @Success 403 {object} base.ForbidenError "Forbiden"
// @Success 404 {object} base.NotFoundError "Not Found"
// @Router /profile [get]
func (r RetrieveProfile) Execute(res http.ResponseWriter, req *http.Request) {
	/// Init repository
	var repository rep.UserRepository
	repository = new(imp.UserRepositoryImpl)

	/// Build
	result := build(repository)

	/// Return
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(result)
}

func build(r rep.UserRepository) ent.User {
	return r.GetProfile()
}
