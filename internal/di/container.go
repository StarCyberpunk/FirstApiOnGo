package di

import (
	"awesomeProject1/internal/handlers"
	"awesomeProject1/internal/repository/postgres"
	"awesomeProject1/internal/usecase"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type Container struct {
	router http.Handler
	db     *sql.DB
	//USECASE
	createUser *usecase.CreateUserUseCase
	//Repository
	userRepository *postgres.UserRepository
	bankRepository *postgres.BankAccountRepository
	//Handler
	postUsersHandler *handlers.POSTUserHandler
	postAuthHandler  *handlers.POSTAuthHandler
}

func NewContainer() *Container {
	return &Container{
		db: postgres.CreateConnection(),
	}
}

func (c *Container) InitRepository() {
	db := c.db
	c.userRepository = postgres.NewUserRepository(db)
	c.bankRepository = postgres.NewBankAccountRepository(db)
}

func (c *Container) InitUseCases() {
	c.createUser = usecase.NewCreateUserUseCase(c.userRepository)
}

func (c *Container) PostUserHandler() *handlers.POSTUserHandler {
	if c.postUsersHandler == nil {
		c.postUsersHandler = handlers.NewPOSTUserHandler(c.createUser)
	}

	return c.postUsersHandler
}
func (c *Container) PostAuthHandler() *handlers.POSTAuthHandler {
	if c.postAuthHandler == nil {
		c.postAuthHandler = handlers.NewPOSTAuthHandler(c.createUser)
	}

	return c.postAuthHandler
}

func (c *Container) HTTPRouter() http.Handler {
	if c.router != nil {
		return c.router
	}
	router := mux.NewRouter()
	//router.Use(middleware.AuthMidleware)

	router.Handle("/register", c.PostUserHandler()).Methods(http.MethodPost)
	router.Handle("/login", c.PostAuthHandler()).Methods(http.MethodPost)

	c.router = router
	return c.router
}

func (c *Container) CloseConnect() {
	c.db.Close()
}
