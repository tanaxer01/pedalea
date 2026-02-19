// @title Pedalea API
// @version 1.0
// @description Bike rental API for Pedalea
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer {token}
// @securityDefinitions.basic BasicAuth
package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/tanaxer01/pedalea/internal/core/admin"
	"github.com/tanaxer01/pedalea/internal/core/bike"
	"github.com/tanaxer01/pedalea/internal/core/rental"
	"github.com/tanaxer01/pedalea/internal/core/user"
	"github.com/tanaxer01/pedalea/internal/infra/auth"
	"github.com/tanaxer01/pedalea/internal/infra/crypto"
	"github.com/tanaxer01/pedalea/internal/infra/http"
	"github.com/tanaxer01/pedalea/internal/infra/logs"
	"github.com/tanaxer01/pedalea/internal/infra/sqlite"
)

type Specification struct {
	Port             string `envconfig:"PORT" default:"8080"`
	DbFile           string `envconfig:"DB_FILE" default:"pedalea.db"`
	JwtSecret        string `envconfig:"JWT_SECRET" required:"true"`
	AdminCredentials string `envconfig:"ADMIN_CREDENTIALS" required:"true"`
}

func main() {
	var s Specification

	if err := godotenv.Load(".env"); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	err := envconfig.Process("pedalea", &s)
	if err != nil {
		panic(err)
	}

	db, err := sqlite.NewDB(s.DbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	logger := logs.NewLogger()

	userRepo := sqlite.NewUserRepository(db)
	bikeRepo := sqlite.NewBikeRepository(db)
	rentalRepo := sqlite.NewRentalRepository(db)
	txRunner := sqlite.NewTxRunner(db)

	jwtAuth := auth.NewJwtAuth(s.JwtSecret)
	jwtMiddleware := http.NewAuthMiddleware(jwtAuth)

	basicAuth := auth.NewBasicAuth(s.AdminCredentials)
	basicMiddleware := http.NewAuthMiddleware(basicAuth)

	userService := user.NewService(userRepo, &crypto.Crypto{}, jwtAuth)
	userHandler := http.NewUserHandler(userService)

	bikeService := bike.NewService(bikeRepo)
	bikeHandler := http.NewBikeHandler(bikeService)

	rentalService := rental.NewServiceWithTxRunner(bikeRepo, rentalRepo, txRunner)
	rentalHandler := http.NewRentalHandler(rentalService)

	adminService := admin.NewService(userRepo, bikeRepo, rentalRepo)
	adminHandler := http.NewAdminHandler(adminService)

	server := http.NewServer(":"+s.Port, userHandler, bikeHandler, rentalHandler, adminHandler, jwtMiddleware, basicMiddleware, logger)
	defer server.Close()

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
