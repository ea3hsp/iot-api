package api

import (
	"context"

	"github.com/ea3hsp/iot-api/internal/models"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints endpoints definions
type Endpoints struct {
	CreateTelemetry endpoint.Endpoint
}

// MakeEndpoints creates enpoints logic
func MakeEndpoints(svc DomoService, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		CreateTelemetry: wrapEndpoint(makeCreateTelemetry(svc), middlewares),
	}
}

func makeCreateTelemetry(svc DomoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.PostTelemetryReq)
		return svc.PostTelemetry(ctx, req)
	}
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	for _, m := range middlewares {
		e = m(e)
	}
	return e
}
