package apiserver

import (
	"covid-monitoring/internal/app/model"
	"covid-monitoring/internal/app/store"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// New ...
func New(config *Config) *APIServer {
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

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/user", s.handleUser())
}

func (s *APIServer) configureStore() error{
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil{
		return err
	}


	s.store = st

	return nil
}

func (s *APIServer) handleUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.Data{
			ID: 0,
			Country: "duda",
			Cases: 2121,
			Deaths: 333,
			Recovered: 222,
		}

		if err := s.store.Db.QueryRow(
			"INSERT INTO data (country, cases, deaths, recovered) VALUES ($1, $2, $3, $4) RETURNING id",
			data.Country,
			data.Cases,
			data.Deaths,
			data.Recovered,
		).Scan(&data.ID); err != nil {
			panic(err)
		}

		fmt.Println("inserted new value with id ", data.ID)
	}
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}