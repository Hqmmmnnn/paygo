package main

import (
	"log"

	"github.com/Hqqm/paygo/config"
	_authHttpAdapter "github.com/Hqqm/paygo/internal/adapters/http"
	"github.com/Hqqm/paygo/internal/repository"
	"github.com/Hqqm/paygo/internal/usescases"
	"github.com/Hqqm/paygo/server"
	"github.com/Hqqm/paygo/server/maindb"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	cfg, err := config.InitConfig("/paygo/config/config.yml")
	if err != nil {
		panic(err)
	}

	dsn := cfg.GetDsn()
	db, err := maindb.NewPgStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}
	dbConnection := db.GetDBConnection()
	defer func() {
		err := dbConnection.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signingKey := []byte(viper.GetString("auth.signing_key"))
	tokenTTL := viper.GetDuration("auth.token_ttl")

	accountRepository := repository.NewAccountRepository(dbConnection)
	authUC := usescases.NewAuthUsecases(accountRepository, signingKey, tokenTTL)
	authMiddleware := _authHttpAdapter.NewAuthMiddleware(authUC)
	authService := _authHttpAdapter.NewAuthService(authUC, *authMiddleware)

	userRepository := repository.NewUserRepository(dbConnection)
	accSettingsUC := usescases.NewAccountSettingsUsecases(userRepository, accountRepository)
	accSettingsService := _authHttpAdapter.NewAccountSettingsService(accSettingsUC)

	transferRepository := repository.NewTransferRepository(dbConnection)
	moneyOperationsUC := usescases.NewMoneyOperationsUsecases(accountRepository, transferRepository)
	moneyOperationsService := _authHttpAdapter.NewMoneyOperationsService(moneyOperationsUC)

	serv := server.NewServer(authService, accSettingsService, moneyOperationsService)
	port := cfg.GetPort()
	if err := serv.Run(port); err != nil {
		log.Fatal(err.Error())
	}
}
