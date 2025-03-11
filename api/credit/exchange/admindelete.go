package exchange

import (
	"context"

	exchange1 "github.com/NpoolPlatform/billing-gateway/pkg/credit/exchange"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/gw/v1/credit/exchange"
)

func (s *Server) AdminDeleteExchange(ctx context.Context, in *npool.AdminDeleteExchangeRequest) (*npool.AdminDeleteExchangeResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithID(&in.ID, true),
		exchange1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteExchange",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteExchange(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteExchange",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteExchangeResponse{
		Info: info,
	}, nil
}
