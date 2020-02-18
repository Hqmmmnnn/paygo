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

	"github.com/Hqqm/paygo/internal/maindb"
	_userHttpAdapter "github.com/Hqqm/paygo/internal/auth/adapters/http"
	"github.com/Hqqm/paygo/internal/auth/repository"
	"github.com/Hqqm/paygo/internal/auth/usescases"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// App ...
type app struct {
	httpServer  *http.Server
	authService *_userHttpAdapter.AuthService
}

// NewApp ...
func NewApp(dsn string) *app {
	pg, err := maindb.NewPgStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}

	signingKey := []byte(viper.GetString("auth.signing_key"))
	tokenTtl := viper.GetDuration("auth.token_ttl")

	userRep := repository.NewPgUserRepository(pg.GetDB())
	authUC := usescases.NewAuthUsecases(userRep, signingKey, tokenTtl)
	authService := _userHttpAdapter.NewAuthService(authUC)

	return &app{
		authService: authService,
	}
}

// Run ...
func (app *app) Run(port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", hi)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signUp", app.authService.SignUp).Methods("POST")
	auth.HandleFunc("/signIn", app.authService.SignIn).Methods("POST")

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
	if _, err := w.Write([]byte(fmt.Sprintf("Hiii"))); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}
