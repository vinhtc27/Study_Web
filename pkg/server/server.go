package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Server Struct
type Server struct {
	srv *http.Server
}

// Server Configuration Struct
type serverConfig struct {
	IP   string
	Port string
}

// Server Configuration Variable
var ServerCfg serverConfig

// NewServer Function to Create a New Server Handler
func NewServer(handler http.Handler) *Server {
	// Initialize New Server
	return &Server{
		srv: &http.Server{
			Addr:    net.JoinHostPort(ServerCfg.IP, ServerCfg.Port),
			Handler: handler,
		},
	}
}

// Start Method for Server
func (s *Server) Start() {
	// Initialize Context Handler Without Timeout
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start server
	fmt.Println("{\"label\":\"server-http\",\"level\":\"info\",\"msg\":\"server worker started at pid " + strconv.Itoa(os.Getpid()) + " listening on " + net.JoinHostPort(ServerCfg.IP, ServerCfg.Port) + "\",\"service\":\"" + Config.GetString("SERVER_NAME") + "\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")

	// Server handle all incoming request
	errs := make(chan error, 1)
	go func() {
		errs <- s.srv.ListenAndServe()

		// Show error
		<-errs
	}()

}

// Stop Method for Server
func (s *Server) Stop() {
	// Initialize Timeout
	timeout := 5 * time.Second

	fmt.Println("{\"label\":\"server-http\",\"level\":\"info\",\"msg\":\"server worker stoped at pid " + strconv.Itoa(os.Getpid()) + " listening on " + net.JoinHostPort(ServerCfg.IP, ServerCfg.Port) + "\",\"service\":\"" + Config.GetString("SERVER_NAME") + "\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")

	// Initialize Context Handler With Timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Hanlde Any Error While Stopping Server
	if err := s.srv.Shutdown(ctx); err != nil {
		if err = s.srv.Close(); err != nil {
			log.Fatalln("{\"label\":\"server-http\",\"level\":\"error\",\"msg\":\"" + err.Error() + "\",\"service\":\"" + Config.GetString("SERVER_NAME") + "\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")
			return
		}
	}
}
