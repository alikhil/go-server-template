package hello

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Server - struct for managing http server
type Server struct {
	*http.Server
	ctx       *Context
	appRouter *mux.Router
}

// NewServer - create new instance of server
func NewServer(ctx *Context) *Server {
	var s = &Server{
		ctx:       ctx,
		appRouter: mux.NewRouter(),
	}
	s.initRoutes()
	return s
}

func addCORSMiddlewares(cfg *Config, next http.Handler) http.Handler {

	// TODO: update methods, origins and headers according to your needs
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                      // All origins
		AllowedMethods: []string{"GET", "POST", "OPTIONS"}, // Allowing only get, just an example
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Authorization,Content-Type,Access-Content-Allow-Origin")

		crs.Handler(next).ServeHTTP(w, r)
	})
}

func (s *Server) initRoutes() {

	s.appRouter.HandleFunc("/health/live", s.ctx.LivenessHealthcheckHandler).Methods("GET")
	s.appRouter.HandleFunc("/health/ready", s.ctx.ReadinessHealthcheckHandler).Methods("GET")
	s.appRouter.HandleFunc("/", s.ctx.HelloHandler).Methods("GET")

	s.Server = &http.Server{
		Addr:    s.ctx.Config.ServeAddress,
		Handler: addCORSMiddlewares(s.ctx.Config, s.appRouter),
	}
}
