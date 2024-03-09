package server

import (
	"context"
	"fmt"
	"gateway/internal/handler"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

func NewRouter(ctx context.Context) (chi.Router, error) {
	swagger, err := handler.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("error in GetSwagger: %w", err)
	}
	swagger.Servers = nil

	if err := swagger.Validate(ctx, openapi3.DisableExamplesValidation()); err != nil {
		return nil, fmt.Errorf("error in swagger.Validate: %w", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.OapiRequestValidator(swagger))

	return router, nil
}
