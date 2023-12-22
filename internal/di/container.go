package di

import (
	"awesomeProject1/internal/handlers"
	"awesomeProject1/internal/pkg/persistence"
	postgrespkg "awesomeProject1/internal/pkg/persistence/postgres"
	"awesomeProject1/internal/repository/postgres"
	"awesomeProject1/internal/usecase"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"os"
)

type Container struct {
	router      http.Handler
	connection  persistence.Connection
	secretKey   string
	databaseURL string
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

func NewContainer() *Container {
	return &Container{}
}
func (c *Container) Pool(ctx context.Context) persistence.Connection {
	if c.connection == nil {
		postgresPool, err := pgxpool.New(ctx, c.DatabaseURL())
		if err != nil {
			panic(err)
		}

		if err := postgresPool.Ping(ctx); err != nil {
			panic(err)
		}

		c.connection = postgrespkg.NewPoolConnection(postgresPool)
	}

	return c.connection
}

func (c *Container) DatabaseURL() string {
	if c.databaseURL == "" {
		c.databaseURL = os.Getenv("DATABASE_URL")
	}

	return c.databaseURL
}

func (c *Container) InitRepository() {
	ctx := context.Background()
	c.userRepository = postgres.NewUserRepository(c.Pool(ctx))
	c.bankRepository = postgres.NewBankAccountRepository(c.Pool(ctx))
	/*c.cardRepository = postgres.NewCardRepository(db)
	c.currencyRepository = postgres.NewCurrencyRepository(db)
	c.operationBARepository = postgres.NewOperationBARepository(db)
	c.operationCardRepository = postgres.NewOperationCardRepository(db)*/
}

func (c *Container) InitUseCases() {
	c.createUser = usecase.NewCreateUserUseCase(c.userRepository, c.bankRepository)
}

func (c *Container) SecretKey() string {
	if c.secretKey == "" {
		c.secretKey = os.Getenv("SECRET_KEY")
	}

	return c.secretKey
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

	router.Handle("/api/users", c.PostUserHandler()).Methods(http.MethodPost)
	router.Handle("/api/tokens", c.PostAuthHandler()).Methods(http.MethodPost)

	c.router = router
	return c.router
}
