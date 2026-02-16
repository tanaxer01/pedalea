package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

func NewServer(addr string, userHandler *UserHandler) *Server {
	s := &Server{addr: addr}
	mux := http.NewServeMux()

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is up"))
	})

	s.registerUserRoutes(mux, userHandler)
	// TODO: Add handlers

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

func (s *Server) registerUserRoutes(mux *http.ServeMux, handler *UserHandler) {
	mux.HandleFunc("POST /user/register", handler.Register)
	mux.HandleFunc("POST /user/login", handler.Login)
	mux.HandleFunc("GET /user/profile", handler.GetUserData)
	mux.HandleFunc("PATCH /user/profile", handler.UpdateUser)
}
