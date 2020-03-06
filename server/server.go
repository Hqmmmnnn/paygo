package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hqqm/paygo/internal/_lib"
	_authHttpAdapter "github.com/Hqqm/paygo/internal/adapters/http"
	"github.com/Hqqm/paygo/internal/entities"
	"github.com/Hqqm/paygo/internal/repository"
	"github.com/Hqqm/paygo/internal/usescases"
	"github.com/Hqqm/paygo/server/maindb"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type App struct {
	httpServer  *http.Server
	authService *_authHttpAdapter.AuthService
	userService *_authHttpAdapter.UserService
}

func NewApp(dsn string) *App {
	pg, err := maindb.NewPgStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}

	signingKey := []byte(viper.GetString("auth.signing_key"))
	tokenTTL := viper.GetDuration("auth.token_ttl")

	accountRepository := repository.NewAccountRepository(pg.GetDB())
	authUC := usescases.NewAuthUsecases(accountRepository, signingKey, tokenTTL)
	authMiddleware := _authHttpAdapter.NewAuthMiddleware(authUC)
	authService := _authHttpAdapter.NewAuthService(authUC, *authMiddleware)

	userRepository := repository.NewUserRepository(pg.GetDB())
	userUC := usescases.NewUserUsecase(userRepository, accountRepository)
	userService := _authHttpAdapter.NewAccountService(userUC)

	return &App{
		authService: authService,
		userService: userService,
	}
}

func (app *App) Run(port string) error {
	r := mux.NewRouter()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signUp", app.authService.SignUp).Methods("POST")
	auth.HandleFunc("/signIn", app.authService.SignIn).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/hi", hi)
	api.HandleFunc("/addUserInfo", app.userService.AddUserInfoToAccount).Methods("POST")
	api.HandleFunc("/getUserInfo", app.userService.GetUserById).Methods("GET")
	api.Use(LoggerMiddleware)
	api.Use(app.authService.Middleware.VerifyToken)

	app.httpServer = &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("Starting Server on port %s", port)
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return app.httpServer.Shutdown(ctx)
}

func hi(w http.ResponseWriter, r *http.Request) {
	if account := r.Context().Value("account").(*entities.Account); account != nil {
		_lib.MarshalJsonAndWrite(account, w)
	}
}
