package middleware

/*
import (
	"context"
	"github.com/gofrs/uuid"
	"net/http"
)

type authKey struct {
}

func AuthMidleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		auth := request.Header.Get("Au")
		userId, err := uuid.FromString(auth)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := contextWithUserID(request, userId)
		handler.Register(writer, ctx)
	})
}

func contextWithUserID(request *http.Request, userId uuid.UUID) context.Context {
	ctx := context.WithValue(request.Context(), authKey{}, userId)
	request = request.WithContext(ctx)
	return ctx
}

func UserIDFromConext(ctx context.Context) uuid.UUID {

}
*/
