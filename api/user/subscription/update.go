package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-gateway/pkg/user/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/gw/v1/user/subscription"
)

func (s *Server) UpdateSubscription(ctx context.Context, in *npool.UpdateSubscriptionRequest) (*npool.UpdateSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.AppID, false),
		subscription1.WithUserID(&in.UserID, false),
		subscription1.WithStartAt(in.StartAt, false),
		subscription1.WithEndAt(in.EndAt, false),
		subscription1.WithUsageState(in.UsageState, false),
		subscription1.WithSubscriptionCredit(in.SubscriptionCredit, false),
		subscription1.WithAddonCredit(in.AddonCredit, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateSubscriptionResponse{
		Info: info,
	}, nil
}
