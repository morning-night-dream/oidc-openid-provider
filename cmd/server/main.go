package main

import "github.com/morning-night-dream/oidc-openid-provider/internal/handler"

func main() {
	srv := handler.NewServer("1234", handler.NewHandler())

	srv.Run()
}
