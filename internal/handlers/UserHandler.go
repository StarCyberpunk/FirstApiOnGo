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
	id uuid.UUID `json:"id"`
}

func (response *POSTUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		id uuid.UUID `json:"id"`
	}{
		id: response.id,
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

	response := &POSTUserResponse{
		id: id_us,
	}

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
