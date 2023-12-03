package handlers

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/internal/usecase"
	"encoding/json"
	"net/http"
)

type POSTAuthHandler struct {
	useCase *usecase.CreateUserUseCase
}

func NewPOSTAuthHandler(useCase *usecase.CreateUserUseCase) *POSTAuthHandler {
	return &POSTAuthHandler{useCase: useCase}
}

type POSTAuthResponse struct {
	AccessToken string `json:"access_token"`
}

func (response *POSTAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		AccessToken string `json:"accessToken"`
	}{
		AccessToken: response.AccessToken,
	})
}

func (handler *POSTAuthHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body domain.UserAuthModel
	err := json.NewDecoder(request.Body).Decode(&body)

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
