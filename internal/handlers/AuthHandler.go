package handlers

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/internal/usecase"
	"encoding/json"
	"net/http"
	"os"
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

var jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))

func (response *POSTAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		token string `json:"access_token"`
	}{
		token: response.AccessToken,
	})
}

func (handler *POSTAuthHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body domain.UserAuthModel
	err := json.NewDecoder(request.Body).Decode(&body)

	token, err := handler.useCase.Login(body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	response := &POSTAuthResponse{
		AccessToken: token,
	}

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
