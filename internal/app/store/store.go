package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	Config *Config
	Db *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		Config: config,
	}
}

func (s *Store) Open() error{
	fmt.Printf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Config.HostPort, s.Config.Hostname, s.Config.Username, s.Config.Password, s.Config.DatabaseName)
	connection := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Config.HostPort, s.Config.Hostname, s.Config.Username, s.Config.Password, s.Config.DatabaseName)
	db, err := sql.Open("postgres", connection)
	if err != nil{
		return err
	}

	if err := db.Ping(); err != nil{
		return err
	}

	s.Db = db
	return nil
}

func (s *Store) Close(){
	s.Db.Close()
}

func (s *Store) getDb() {
	return
}