package api

import _ "github.com/Meruya-Technology/go-boilerplate/docs"

// RetrieveUser
// @Summary Retrieve user profile
// @description Retrieve user profile
// @ID retrieve-user
// @Accept  json
// @Produce  json
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /auth/profile [get]
