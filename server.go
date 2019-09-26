package server

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	communication "github.com/ltsnuggie/communication-service"
	log "github.com/sirupsen/logrus"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	// Route{
	// 	"Get Player History",
	// 	"GET",
	// 	"/api/v1/lookup/{platform}/{name}",
	// 	PlayerLookup,
	// },
}

type Server struct {
	router *mux.Router
	server *http.Server
}

// var router *mux.Router
var server *http.Server

func New(port string) Server {

	r := mux.NewRouter()

	r.Use(communication.LoggingMiddleware)
	log.WithFields(log.Fields{
		"date": time.Now(),
	}).Warn("Starting server...")

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})
	// Access-Control-Allow-Headers: Authorization

	// Launch server with CORS validations
	server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":" + port,
		Handler:      handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(r),
	}
	log.WithFields(log.Fields{
		"date": time.Now(),
	}).Warn("Server started...")
	go server.ListenAndServe()

	return Server{
		router: r,
		server: server,
	}
}

func (s *Server) LoadEndpoint(name, path, method string, handlerFunc http.HandlerFunc) {
	log.WithFields(log.Fields{
		"httpMethod": method,
		"endpoint":   path,
		"date":       time.Now(),
	}).Infof("%v end point loaded...", name)

	s.router.
		Methods(method).
		Path(path).
		Name(name).
		Handler(handlerFunc)
}

func (s *Server) Stop() {
	s.server.Close()
}
