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

type POSTUserRequest struct {
	Name string `json:"name"`
}

type POSTUserResponse struct {
	id uuid.UUID
}

func (response *POSTUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID uuid.UUID `json:"id"`
	}{
		ID: response.id,
	})
}

func (handler *POSTUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTUserRequest
	err := json.NewDecoder(request.Body).Decode(&body)

	user := domain.User{}

	uuser, err := handler.useCase.Handle(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &POSTUserResponse{
		id: uuser.ID,
	}

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
