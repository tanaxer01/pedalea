package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

func NewServer(addr string, userHandler *UserHandler, bikeHandler *BikeHandler, rentalHandler *RentalHandler, adminHandler *AdminHandler, jwtMiddleware *AuthMiddleware, adminMiddleware *AuthMiddleware) *Server {
	s := &Server{addr: addr}
	mux := http.NewServeMux()

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is up"))
	})

	s.registerUserRoutes(mux, userHandler, jwtMiddleware)
	s.registerBikeRoutes(mux, bikeHandler, jwtMiddleware)
	s.registerRentalRoutes(mux, rentalHandler, jwtMiddleware)
	s.registerAdminRoutes(mux, adminHandler, adminMiddleware)

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

func (s *Server) registerUserRoutes(mux *http.ServeMux, h *UserHandler, m *AuthMiddleware) {
	mux.HandleFunc("POST /user/register", h.Register)
	mux.HandleFunc("POST /user/login", h.Login)
	mux.HandleFunc("GET /user/profile", m.AuthMiddleware(h.GetUserData))
	mux.HandleFunc("PATCH /user/profile", m.AuthMiddleware(h.UpdateUser))
}

func (s *Server) registerBikeRoutes(mux *http.ServeMux, h *BikeHandler, m *AuthMiddleware) {
	mux.HandleFunc("GET /bikes/available", m.AuthMiddleware(h.ListAvailableBikes))
}

func (s *Server) registerRentalRoutes(mux *http.ServeMux, h *RentalHandler, m *AuthMiddleware) {
	mux.HandleFunc("GET /rentals/start", m.AuthMiddleware(h.StartRental))
	mux.HandleFunc("GET /rentals/end", m.AuthMiddleware(h.EndRental))
	mux.HandleFunc("GET /rentals/history", m.AuthMiddleware(h.ListUserRentals))

}

func (s *Server) registerAdminRoutes(mux *http.ServeMux, handler *AdminHandler, m *AuthMiddleware) {
	// Bikes
	mux.HandleFunc("POST /admin/bikes", m.AuthMiddleware(handler.InsertBike))
	mux.HandleFunc("PATCH /admin/bikes", m.AuthMiddleware(handler.UpdateBike))
	mux.HandleFunc("GET /admin/bikes", m.AuthMiddleware(handler.ListBikes))

	// Users
	mux.HandleFunc("GET /admin/users", m.AuthMiddleware(handler.ListUsers))
	mux.HandleFunc("GET /admin/users/{user_id}", m.AuthMiddleware(handler.GetUser))
	mux.HandleFunc("PATCH /admin/users/{user_id}", m.AuthMiddleware(handler.UpdateUser))

	// Rentals
	mux.HandleFunc("PATCH /admin/rentals", m.AuthMiddleware(handler.UpdateRental))
	mux.HandleFunc("GET /admin/rentals/{rental_id}", m.AuthMiddleware(handler.GetRental))
	mux.HandleFunc("GET /admin/rentals", m.AuthMiddleware(handler.ListRentals))
}
