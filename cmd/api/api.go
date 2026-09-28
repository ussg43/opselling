package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/ussg43/opselling/internal/store"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	httpPort string
	grpcPort string
	sheets   string
}

type Application struct {
	config  config
	httpSrv *http.Server
	grpcSrv *grpc.Server
	db      *sqlx.DB
	storage store.Storage
}

func (a *Application) mount() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	a.httpSrv = &http.Server{
		Addr:         a.config.httpPort,
		Handler:      r,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	a.grpcSrv = grpc.NewServer()
	reflection.Register(a.grpcSrv)
}

func (a *Application) run(ctx context.Context) error {
	httpL, err := net.Listen("tcp", a.config.httpPort)
	if err != nil {
		log.Fatal()
		return err
	}

	grpcL, err := net.Listen("tcp", a.config.grpcPort)
	if err != nil {
		httpL.Close()
		log.Fatal()
		return err
	}

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return a.grpcSrv.Serve(grpcL)
	})

	g.Go(func() error {
		return a.httpSrv.ListenAndServe()
	})

	g.Go(func() error {
		<-gctx.Done()

		shutdown, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		done := make(chan struct{})
		go func() {
			a.grpcSrv.GracefulStop()
			close(done)
		}()

		go func() {
			_ = a.httpSrv.Shutdown(shutdown)
		}()

		select {
		case <-done:
		case <-shutdown.Done():
			a.grpcSrv.Stop()
		}

		return nil
	})

	return g.Wait()
}
