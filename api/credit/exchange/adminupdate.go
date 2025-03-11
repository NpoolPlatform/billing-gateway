package exchange

import (
	"context"

	exchange1 "github.com/NpoolPlatform/billing-gateway/pkg/credit/exchange"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/gw/v1/credit/exchange"
)

func (s *Server) AdminUpdateExchange(ctx context.Context, in *npool.AdminUpdateExchangeRequest) (*npool.AdminUpdateExchangeResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithID(&in.ID, true),
		exchange1.WithEntID(&in.EntID, true),
		exchange1.WithAppID(&in.TargetAppID, true),
		exchange1.WithCredit(&in.Credit, false),
		exchange1.WithExchangeThreshold(&in.ExchangeThreshold, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateExchange",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateExchange(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateExchange",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateExchangeResponse{
		Info: info,
	}, nil
}
