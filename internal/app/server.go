package app

import (
	"net/http"

	"github.com/imantung/fullstack-golang-example/internal/app/infra/config"
	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
)

var _ = di.Provide(NewEcho)

func NewEcho(cfg *config.Config) *http.Server {
	server := &http.Server{
		Addr: cfg.Address,
	}

	http.Handle("/", http.FileServer(http.Dir("./public")))
	return server
}
