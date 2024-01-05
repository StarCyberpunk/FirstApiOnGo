package handlers

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/internal/usecase/auth"
	"encoding/json"
	"net/http"
)

type POSTAuthHandler struct {
	useCase *auth.CreateUserUseCase
}

func NewPOSTAuthHandler(useCase *auth.CreateUserUseCase) *POSTAuthHandler {
	return &POSTAuthHandler{useCase: useCase}
}

type POSTAuthResponse struct {
	AccessToken string `json:"access_token"`
}

func (handler *POSTAuthHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body domain.UserAuthModel
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := handler.useCase.Login(request.Context(), body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}
	response := &POSTAuthResponse{
		AccessToken: token,
	}

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
