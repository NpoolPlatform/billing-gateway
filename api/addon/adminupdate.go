package addon

import (
	"context"

	addon1 "github.com/NpoolPlatform/billing-gateway/pkg/addon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/gw/v1/addon"
)

func (s *Server) AdminUpdateAddon(ctx context.Context, in *npool.AdminUpdateAddonRequest) (*npool.AdminUpdateAddonResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithID(&in.ID, true),
		addon1.WithEntID(&in.EntID, true),
		addon1.WithAppID(&in.TargetAppID, false),
		addon1.WithPrice(in.Price, false),
		addon1.WithCredit(in.Credit, false),
		addon1.WithSortOrder(in.SortOrder, false),
		addon1.WithEnabled(in.Enabled, false),
		addon1.WithDescription(&in.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateAddon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateAddonResponse{
		Info: info,
	}, nil
}
