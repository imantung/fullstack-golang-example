package app

import (
	"context"
	"net/http"
)

func Start(server *http.Server) error {
	return server.ListenAndServe()
}

func Stop(server *http.Server) error {
	ctx := context.Background()
	return server.Shutdown(ctx)
}
