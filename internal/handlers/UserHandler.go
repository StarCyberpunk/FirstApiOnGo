package handlers

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/internal/usecase"
	"encoding/json"
	"github.com/gofrs/uuid"
	"net/http"
)

type POSTUserHandler struct {
	useCase *usecase.CreateUserUseCase
}

func NewPOSTUserHandler(useCase *usecase.CreateUserUseCase) *POSTUserHandler {
	return &POSTUserHandler{useCase: useCase}
}

type POSTUserResponse struct {
	AccessToken string `json:"access_token"`
}

func (response *POSTUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		AccessToken string `json:"accessToken"`
	}{
		AccessToken: response.AccessToken,
	})
}
func (handler *POSTUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body domain.UserRegisterModel
	err := json.NewDecoder(request.Body).Decode(&body)

	id_us, err := handler.useCase.Handle(body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if id_us == uuid.Nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}
	body2 := domain.UserAuthModel{Login: body.Login, Password: body.Password}
	token, err := handler.useCase.Login(body2)
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
