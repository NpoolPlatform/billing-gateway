package api

import (
	"context"

	addon "github.com/NpoolPlatform/billing-gateway/api/addon"
	exchange "github.com/NpoolPlatform/billing-gateway/api/credit/exchange"
	subscription "github.com/NpoolPlatform/billing-gateway/api/subscription"
	record "github.com/NpoolPlatform/billing-gateway/api/user/credit/record"
	usersubscription "github.com/NpoolPlatform/billing-gateway/api/user/subscription"
	servicetmpl "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	servicetmpl.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	servicetmpl.RegisterMiddlewareServer(server, &Server{})
	addon.Register(server)
	subscription.Register(server)
	exchange.Register(server)
	usersubscription.Register(server)
	record.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := servicetmpl.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := addon.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := subscription.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := exchange.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := usersubscription.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := record.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
