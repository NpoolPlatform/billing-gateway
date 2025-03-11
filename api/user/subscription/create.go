package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/billing-gateway/pkg/user/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/gw/v1/user/subscription"
)

func (s *Server) CreateSubscription(ctx context.Context, in *npool.CreateSubscriptionRequest) (*npool.CreateSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(in.UserID, true),
		subscription1.WithPackageID(in.PackageID, true),
		subscription1.WithStartAt(in.StartAt, true),
		subscription1.WithEndAt(in.EndAt, true),
		subscription1.WithOrderID(in.OrderID, true),
		subscription1.WithUsageState(in.UsageState, true),
		subscription1.WithSubscriptionCredit(in.SubscriptionCredit, true),
		subscription1.WithAddonCredit(in.AddonCredit, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSubscriptionResponse{
		Info: info,
	}, nil
}
