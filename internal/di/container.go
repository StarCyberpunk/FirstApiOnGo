package di

import (
	"awesomeProject1/internal/handlers"
	"awesomeProject1/internal/handlers/middleware"
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
	userRepository          *postgres.UserRepository
	bankRepository          *postgres.BankAccountRepository
	cardRepository          *postgres.CardRepository
	currencyRepository      *postgres.CurrencyRepository
	operationBARepository   *postgres.OperationBARepository
	operationCardRepository *postgres.OperationCardRepository
	//Handler
	postUsersHandler *handlers.POSTUserHandler
	postAuthHandler  *handlers.POSTAuthHandler
}

func NewContainer(dbb *sql.DB) *Container {
	return &Container{
		db: dbb,
	}
}

func (c *Container) InitRepository() {
	// db перенести
	db := c.db
	c.userRepository = postgres.NewUserRepository(db)
	c.bankRepository = postgres.NewBankAccountRepository(db)
	c.cardRepository = postgres.NewCardRepository(db)
	c.currencyRepository = postgres.NewCurrencyRepository(db)
	c.operationBARepository = postgres.NewOperationBARepository(db)
	c.operationCardRepository = postgres.NewOperationCardRepository(db)
}

func (c *Container) InitUseCases() {
	c.createUser = usecase.NewCreateUserUseCase(c.userRepository, c.bankRepository)
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
	router.Use(middleware.AuthMidleware)
	router.Handle("/api/users", c.PostUserHandler()).Methods(http.MethodPost)
	router.Handle("/api/tokens", c.PostAuthHandler()).Methods(http.MethodPost)

	c.router = router
	return c.router
}

func (c *Container) CloseConnect() {
	c.db.Close()
}
