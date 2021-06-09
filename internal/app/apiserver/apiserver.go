package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lehius/http-resy-api/internal/config"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.Network.Bind, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.Logging.Level)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}

}
