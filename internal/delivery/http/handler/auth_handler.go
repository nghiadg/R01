package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"r01/internal/delivery/http/http_utils"
	"r01/internal/usecase/auth"
)

type AuthHandler struct {
	loginUsecase auth.LoginUsecase
}

func NewAuthHandler(loginUsecase auth.LoginUsecase) *AuthHandler {
	return &AuthHandler{loginUsecase: loginUsecase}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// bind body to struct
	var params auth.LoginParams
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	result, err := ah.loginUsecase.Execute(params)

	if err != nil {
		if errors.Is(err, auth.ErrInvalidPassword) {
			http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
				Status:  http.StatusBadRequest,
				Message: "Email or password is incorrect",
			})

			return
		}

		http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		})
		return
	}

	http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
		Status:  http.StatusOK,
		Message: "Login success",
		Data:    result,
	})
}
