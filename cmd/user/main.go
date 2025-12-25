package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"usermic/internal/config"
	"usermic/internal/infrastructure/postgres"
	grpcserv "usermic/internal/transport/grpc"
	"usermic/internal/usecase/registration"
	"usermic/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Storage.User,
		cfg.Storage.Password,
		cfg.Storage.Host,
		strconv.Itoa(cfg.Storage.Port),
		cfg.Storage.Dbname,
		cfg.Storage.SslMode,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	repo := postgres.NewPostgresStorage(log, db)

	regUc := registration.NewRegistrationUC(log, repo)

	serv := grpcserv.NewServer(log, regUc, cfg.Timeout)

	go serv.MustStart(cfg.GRPC.Port)

	sysChan := make(chan os.Signal, 1)
	signal.Notify(sysChan, syscall.SIGTERM, syscall.SIGINT)

	<-sysChan

	serv.Stop()
}
