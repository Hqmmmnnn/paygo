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
	_userHttpAdapter "github.com/Hqqm/paygo/internal/user/adapters/http"
	"github.com/Hqqm/paygo/internal/user/domain/repository"
	"github.com/Hqqm/paygo/internal/user/domain/usescases"
	"github.com/gorilla/mux"
)

// App ...
type App struct {
	httpServer   *http.Server
	userUsecases usescases.UserUsecases
}

// NewApp ...
func NewApp(dsn string) *App {
	pg, err := maindb.NewPgStorage(dsn)

	if err != nil {
		log.Fatal(err)
	}

	return &App{
		userUsecases: usescases.UserUsecases{UserRepository: repository.NewPgUserRepository(pg.Db)},
	}

}

// Run ...
func (app *App) Run(port string) error {
	userService := _userHttpAdapter.NewUserService(app.userUsecases)

	r := mux.NewRouter()
	r.HandleFunc("/", hi)
	r.HandleFunc("/createUser", userService.CreateUser).Methods("POST")

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
