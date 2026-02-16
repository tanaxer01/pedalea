package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

func NewServer(addr string, userHandler *UserHandler, bikeHandler *BikeHandler, rentalHandler *RentalHandler, jwtMiddleware *JwtMiddleware) *Server {
	s := &Server{addr: addr}
	mux := http.NewServeMux()

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is up"))
	})

	s.registerUserRoutes(mux, userHandler, jwtMiddleware)
	s.registerBikeRoutes(mux, bikeHandler, jwtMiddleware)

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

func (s *Server) registerAdminRoutes(mux *http.ServeMux, handler *AdminHandler) {
	// Bikes
	// Users
	mux.HandleFunc("GET /admin/users", handler.ListUsers)
	mux.HandleFunc("GET /admin/users/{user_id}", handler.GetUser)
	mux.HandleFunc("PATCH /admin/users/{user_id}", handler.UpdateUser)

}
