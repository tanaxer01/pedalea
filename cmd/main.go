package main

import (
	"github.com/tanaxer01/pedalea/internal/core/admin"
	"github.com/tanaxer01/pedalea/internal/core/bike"
	"github.com/tanaxer01/pedalea/internal/core/rental"
	"github.com/tanaxer01/pedalea/internal/core/user"
	"github.com/tanaxer01/pedalea/internal/infra/auth"
	"github.com/tanaxer01/pedalea/internal/infra/crypto"
	"github.com/tanaxer01/pedalea/internal/infra/http"
	"github.com/tanaxer01/pedalea/internal/infra/sqlite"
)

func main() {
	db, err := sqlite.NewDB("pedalea.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := sqlite.NewUserRepository(db)
	bikeRepo := sqlite.NewBikeRepository(db)
	rentalRepo := sqlite.NewRentalRepository(db)

	jwtAuth := auth.NewJwtAuth("secret")
	jwtMiddleware := http.NewAuthMiddleware(jwtAuth)

	basicAuth := auth.NewBasicAuth("YWRtaW46cGFzc3dvcmQ=")
	basicMiddleware := http.NewAuthMiddleware(basicAuth)

	userService := user.NewService(userRepo, &crypto.Crypto{}, jwtAuth)
	userHandler := http.NewUserHandler(userService)

	bikeService := bike.NewService(bikeRepo)
	bikeHandler := http.NewBikeHandler(bikeService)

	rentalService := rental.NewService(bikeRepo, rentalRepo)
	rentalHandler := http.NewRentalHandler(rentalService)

	adminService := admin.NewService(userRepo, bikeRepo, rentalRepo)
	adminHandler := http.NewAdminHandler(adminService)

	server := http.NewServer(":8080", userHandler, bikeHandler, rentalHandler, adminHandler, jwtMiddleware, basicMiddleware)
	defer server.Close()

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
