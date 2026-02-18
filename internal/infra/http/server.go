package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tanaxer01/pedalea/pkg/utils"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

func NewServer(addr string, userHandler *UserHandler, bikeHandler *BikeHandler, rentalHandler *RentalHandler, adminHandler *AdminHandler, jwtMiddleware *AuthMiddleware, adminMiddleware *AuthMiddleware) *Server {
	s := &Server{addr: addr}
	r := chi.NewRouter()

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteResponse(w, "OK")
	})

	r.Mount("/user", s.userRouter(userHandler, jwtMiddleware))
	r.Mount("/bikes", s.bikeRouter(bikeHandler, jwtMiddleware))
	r.Mount("/routes", s.rentalRouter(rentalHandler, jwtMiddleware))
	r.Mount("/admin", s.adminRouter(adminHandler, adminMiddleware))

	s.httpServer = &http.Server{Addr: addr, Handler: r}

	return s
}

func (s *Server) Start() error {
	fmt.Printf("[+] Server running on %v\n", s.addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Close() error {
	// TODO: Change for Shutdown
	return s.httpServer.Close()
}

func (s *Server) userRouter(h *UserHandler, m *AuthMiddleware) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)

	r.With(m.AuthMiddleware).Get("/profile", h.GetUserData)
	r.With(m.AuthMiddleware).Patch("/profile", h.UpdateUser)

	return r
}

func (s *Server) bikeRouter(h *BikeHandler, m *AuthMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(m.AuthMiddleware)

	r.Get("/available", h.ListAvailableBikes)

	return r
}

func (s *Server) rentalRouter(h *RentalHandler, m *AuthMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(m.AuthMiddleware)

	r.Get("/start", h.StartRental)
	r.Get("/end", h.EndRental)
	r.Get("/history", h.ListUserRentals)

	return r
}

func (s *Server) adminRouter(h *AdminHandler, m *AuthMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(m.AuthMiddleware)

	// Bikes
	r.Post("/bikes", h.InsertBike)
	r.Patch("/bikes/{bike_id}", h.UpdateBike)
	r.Get("/bikes", h.ListBikes)

	// Users
	r.Get("/users", h.ListUsers)
	r.Get("/users/{user_id}", h.GetUser)
	r.Patch("/users/{user_id}", h.UpdateUser)

	// Rentals
	r.Patch("/rentals/{rental_id}", h.UpdateRental)
	r.Get("/rentals/{rental_id}", h.GetRental)
	r.Get("/rentals", h.ListRentals)

	return r
}
