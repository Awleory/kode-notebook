package app

import (
	"fmt"
	"os"

	api "github.com/awleory/kode/notebook/internal/controller/http/v1"
	"github.com/awleory/kode/notebook/internal/repository/psql"
	"github.com/awleory/kode/notebook/internal/service"
	"github.com/awleory/kode/notebook/pkg/config"
	"github.com/awleory/kode/notebook/pkg/database"
	"github.com/awleory/kode/notebook/pkg/hash"
	"github.com/awleory/kode/notebook/pkg/httpserver"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	CONFIG_DIR  = "config"
	CONFIG_FILE = "main"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connection(database.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	noteRepo := psql.NewNotes(db)
	userRepo := psql.NewUsers(db)
	hasher := hash.NewSHA1Hasher(os.Getenv("SALT"))

	noteService := service.NewNote(noteRepo)
	userService := service.NewUsers(
		userRepo,
		hasher,
	)

	handler := api.NewHandler(
		userService,
		noteService,
	)

	srv := httpserver.New(
		handler.InitRouter(),
		fmt.Sprintf("%d", cfg.Server.Port),
		cfg.Server.Timeout,
	)

	log.Info("SERVER STARTED")
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
