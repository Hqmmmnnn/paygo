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

	_authHttpAdapter "github.com/Hqqm/paygo/internal/auth/adapters/http"
	"github.com/Hqqm/paygo/internal/auth/repository"
	"github.com/Hqqm/paygo/internal/auth/usescases"
	"github.com/Hqqm/paygo/internal/maindb"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type App struct {
	httpServer  *http.Server
	authService *_authHttpAdapter.AuthService
}

func NewApp(dsn string) *App {
	pg, err := maindb.NewPgStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}

	signingKey := []byte(viper.GetString("auth.signing_key"))
	tokenTTL := viper.GetDuration("auth.token_ttl")

	accountRepository := repository.NewAccountRepository(pg.GetDB())
	userRepository := repository.NewPgUserRepository(pg.GetDB())
	authUC := usescases.NewAuthUsecases(accountRepository, userRepository, signingKey, tokenTTL)
	authMiddleware := _authHttpAdapter.NewAuthMiddleware(authUC)

	authService := _authHttpAdapter.NewAuthService(authUC, *authMiddleware)

	return &App{authService: authService}
}

func (app *App) Run(port string) error {
	r := mux.NewRouter()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signUp", app.authService.SignUp).Methods("POST")
	auth.HandleFunc("/signIn", app.authService.SignIn).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(AccesLoginMiddleware)
	api.Use(app.authService.Middleware.VerifyToken)
	api.HandleFunc("/hi", hi)

	siteHandler := AccesLoginMiddleware(api)
	siteHandler = app.authService.Middleware.VerifyToken(siteHandler)

	app.httpServer = &http.Server{
		Handler:      siteHandler,
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
	if account := r.Context().Value("account"); account != nil {
		w.Write([]byte(fmt.Sprintf("hiii %+v", account)))
	}
}
