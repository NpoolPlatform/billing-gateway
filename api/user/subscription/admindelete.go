package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-gateway/pkg/user/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/gw/v1/user/subscription"
)

func (s *Server) AdminDeleteSubscription(ctx context.Context, in *npool.AdminDeleteSubscriptionRequest) (*npool.AdminDeleteSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteSubscriptionResponse{
		Info: info,
	}, nil
}
