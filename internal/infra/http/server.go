package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

func NewServer(addr string, userHandler *UserHandler, bikeHandler *BikeHandler, rentalHandler *RentalHandler, adminHandler *AdminHandler, jwtMiddleware *JwtMiddleware) *Server {
	s := &Server{addr: addr}
	mux := http.NewServeMux()

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is up"))
	})

	s.registerUserRoutes(mux, userHandler, jwtMiddleware)
	s.registerBikeRoutes(mux, bikeHandler, jwtMiddleware)
	s.registerRentalRoutes(mux, rentalHandler, jwtMiddleware)
	s.registerAdminRoutes(mux, adminHandler)

	s.httpServer = &http.Server{Addr: addr, Handler: mux}

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

func (s *Server) registerUserRoutes(mux *http.ServeMux, h *UserHandler, m *JwtMiddleware) {
	mux.HandleFunc("POST /user/register", h.Register)
	mux.HandleFunc("POST /user/login", h.Login)
	mux.HandleFunc("GET /user/profile", m.JwtValidationMiddleware(h.GetUserData))
	mux.HandleFunc("PATCH /user/profile", m.JwtValidationMiddleware(h.UpdateUser))
}

func (s *Server) registerBikeRoutes(mux *http.ServeMux, h *BikeHandler, m *JwtMiddleware) {
	mux.HandleFunc("GET /bikes/available", m.JwtValidationMiddleware(h.ListAvailableBikes))
}

func (s *Server) registerRentalRoutes(mux *http.ServeMux, h *RentalHandler, m *JwtMiddleware) {
	mux.HandleFunc("GET /rentals/start", m.JwtValidationMiddleware(h.StartRental))
	mux.HandleFunc("GET /rentals/end", m.JwtValidationMiddleware(h.EndRental))
	mux.HandleFunc("GET /rentals/history", m.JwtValidationMiddleware(h.ListUserRentals))

}

func (s *Server) registerAdminRoutes(mux *http.ServeMux, handler *AdminHandler) {
	// Bikes
	mux.HandleFunc("POST /admin/bikes", handler.InsertBike)
	mux.HandleFunc("PATCH /admin/bikes", handler.UpdateBike)
	mux.HandleFunc("GET /admin/bikes", handler.ListBikes)

	// Users
	mux.HandleFunc("GET /admin/users", handler.ListUsers)
	mux.HandleFunc("GET /admin/users/{user_id}", handler.GetUser)
	mux.HandleFunc("PATCH /admin/users/{user_id}", handler.UpdateUser)

	// Rentals
	mux.HandleFunc("PATCH /admin/rentals", handler.UpdateRental)
	mux.HandleFunc("GET /admin/rentals/{rental_id}", handler.GetRental)
	mux.HandleFunc("GET /admin/rentals", handler.ListRentals)
}
