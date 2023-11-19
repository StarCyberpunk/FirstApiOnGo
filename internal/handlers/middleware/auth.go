package middleware

import (
	"context"
	"github.com/gofrs/uuid"
	"net/http"
)

type authKey struct {
}

func AuthMidleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		auth := request.Header.Get("Authorization")
		userId, err := uuid.FromString(auth)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		req := contextWithUserID(request, userId)
		handler.ServeHTTP(writer, req)
	})
}

func contextWithUserID(request *http.Request, userId uuid.UUID) *http.Request {
	ctx := context.WithValue(request.Context(), authKey{}, userId)
	request = request.WithContext(ctx)
	return request
}
