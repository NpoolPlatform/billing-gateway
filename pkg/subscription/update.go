package subscription

import (
	"context"

	submwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/subscription"
	submwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/subscription"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateSubscription(ctx context.Context) (*submwpb.Subscription, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSubscription(ctx); err != nil {
		return nil, err
	}
	if err := submwcli.UpdateSubscription(ctx, &submwpb.SubscriptionReq{
		ID:          h.ID,
		EntID:       h.EntID,
		PackageName: h.PackageName,
		Price:       h.Price,
		Description: h.Description,
		SortOrder:   h.SortOrder,
		Credit:      h.Credit,
		ResetType:   h.ResetType,
		QPSLimit:    h.QPSLimit,
	}); err != nil {
		return nil, err
	}
	return h.GetSubscription(ctx)
}
