package subscription

import (
	"context"

	subscription "github.com/NpoolPlatform/message/npool/billing/gw/v1/subscription"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	subscription.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	subscription.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := subscription.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
