package main

import (
	"fmt"
	"net"
	"net/http"
	"shokHorizon/go_recipes/internal/config"
	"shokHorizon/go_recipes/internal/recipe"
	"shokHorizon/go_recipes/pkg/logger"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	l := logger.GetLogger()
	cfg := config.GetConfig()
	r := httprouter.New()

	rh := recipe.NewHandler(&l)
	rh.Register(r)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.BindIp, cfg.BindIp))
	if err != nil {
		l.Fatal(err)
	}
	l.Info("Server start on address", cfg.BindIp, cfg.Port)

	server := &http.Server{
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
