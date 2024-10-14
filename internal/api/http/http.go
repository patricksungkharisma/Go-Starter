package http

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/patricksungkharisma/go-starter/internal/config"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
)

type Resource struct {
	Database *sqlx.DB
}

func ServerHTTP() error {
	config, err := config.InitConfig()
	if err != nil {
		log.Println(errs.ErrInitConfig)
		return err
	}

	resource, err := initResource(config)
	if err != nil {
		log.Println(errs.ErrInitResource)
		return err
	}

	repo := initRepo(config, resource)
	usecase := initUsecase(config, repo)
	handler := initHandler(config, usecase)
	router := initRouting(handler)

	server := &http.Server{
		Addr:    strings.Join([]string{config.Server.Host, config.Server.Port}, ":"),
		Handler: router,
	}

	go func() {
		log.Println("server running on port ", config.Server.Port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	time.Sleep(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln(errs.ErrShutdownServer)
	}

	return nil
}
