package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/morning-night-dream/oidc-openid-provider/pkg/openapi"
)

var _ openapi.ServerInterface = (*Handler)(nil)

type Handler struct{}

func NewHandler() http.Handler {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {})

	hdl := openapi.HandlerWithOptions(&Handler{}, openapi.ChiServerOptions{
		BaseRouter: router,
	})

	return hdl
}

func (hdl *Handler) OpenIDConfiguration(
	w http.ResponseWriter,
	r *http.Request,
) {
	res := openapi.OpenIDConfigurationSchema{
		Issuer:                "http://localhost:1234",
		AuthorizationEndpoint: "http://localhost:1234/op/auth",
		TokenEndpoint:         "http://localhost:1234/op/token",
		UserinfoEndpoint:      "http://localhost:1234/op/userinfo",
		RevocationEndpoint:    "http://localhost:1234/op/revoke",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("failed to encode response: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (hdl *Handler) Auth(
	w http.ResponseWriter,
	r *http.Request,
) {
}
