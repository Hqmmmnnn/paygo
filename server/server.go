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
	"github.com/gorilla/mux"
)

type Server struct {
	httpServer             *http.Server
	authService            *_authHttpAdapter.AuthService
	accSettingsService     *_authHttpAdapter.AccountSettingsService
	moneyOperationsService *_authHttpAdapter.MoneyOperationsService
}

func NewServer(
	authService *_authHttpAdapter.AuthService,
	accSettingsService *_authHttpAdapter.AccountSettingsService,
	moneyOpService *_authHttpAdapter.MoneyOperationsService) *Server {
	return &Server{
		authService:            authService,
		accSettingsService:     accSettingsService,
		moneyOperationsService: moneyOpService,
	}
}

func (server *Server) Run(port string) error {
	server.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      server.handler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("Starting Server on port %s", port)
		if err := server.httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return server.httpServer.Shutdown(ctx)
}

func (server *Server) handler() http.Handler {
	r := mux.NewRouter()
	r.Use(LoggerMiddleware)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signUp", server.authService.SignUp).Methods("POST")
	auth.HandleFunc("/signIn", server.authService.SignIn).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/hi", hi)
	api.HandleFunc("/addUserInfo", server.accSettingsService.AddUserInfoToAccount).Methods("POST")
	api.HandleFunc("/getUserInfo", server.accSettingsService.GetUserById).Methods("GET")
	api.HandleFunc("/getAccount", server.accSettingsService.GetAccountById).Methods("GET")

	api.HandleFunc("/replenishmentBalance", server.moneyOperationsService.ReplenishmentBalance).Methods("POST")

	api.Use(server.authService.Middleware.VerifyToken)

	return r
}

func hi(w http.ResponseWriter, r *http.Request) {
	if account := r.Context().Value("account").(*entities.Account); account != nil {
		_lib.MarshalJsonAndWrite(account, w)
	}
}
