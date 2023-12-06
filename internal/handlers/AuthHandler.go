package handlers

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/internal/usecase"
	"context"
	"encoding/json"
	"net/http"
)

type POSTAuthHandler struct {
	useCase *usecase.CreateUserUseCase
	ctx     context.Context
}

func NewPOSTAuthHandler(useCase *usecase.CreateUserUseCase, ctx context.Context) *POSTAuthHandler {
	return &POSTAuthHandler{useCase: useCase, ctx: ctx}
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
	token, err := handler.useCase.Login(body)
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
