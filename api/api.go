package api

// GetStringByInt example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Param   some_id      body config.Pet true  "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} config.APIError "We need ID!!"
// @Failure 404 {object} config.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
