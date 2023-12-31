package server

import (
	"context"
	"errors"
	"log"
	"main/database"
	"main/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTsecret   string
	DatabaseURL string
	SudoToken   string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("el puerto es requerido")
	}
	if config.JWTsecret == "" {
		return nil, errors.New("la clave secreta es requerida")
	}
	if config.DatabaseURL == "" {
		return nil, errors.New("la URL de la base de datos es requerida")
	}
	if config.SudoToken == "" {
		return nil, errors.New("debes proporcionar un token de sudo")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil

}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.Config().DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	log.Println("Servidor corriendo en el puerto: ", b.Config().Port)
	if err := http.ListenAndServe(b.Config().Port, b.router); err != nil {
		log.Fatal(err)
	}
}
