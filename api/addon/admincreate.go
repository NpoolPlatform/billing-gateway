package addon

import (
	"context"

	addon1 "github.com/NpoolPlatform/billing-gateway/pkg/addon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/billing/gw/v1/addon"
)

func (s *Server) AdminCreateAddon(ctx context.Context, in *npool.AdminCreateAddonRequest) (*npool.AdminCreateAddonResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithAppID(&in.TargetAppID, true),
		addon1.WithPrice(&in.Price, true),
		addon1.WithCredit(&in.Credit, true),
		addon1.WithSortOrder(&in.SortOrder, false),
		addon1.WithEnabled(&in.Enabled, false),
		addon1.WithDescription(&in.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateAddon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAddonResponse{
		Info: info,
	}, nil
}
