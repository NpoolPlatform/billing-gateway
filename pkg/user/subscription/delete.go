package subscription

import (
	"context"

	subscriptionmwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/user/subscription"
	subscriptionmwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteSubscription(ctx context.Context) (*subscriptionmwpb.Subscription, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSubscription(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetSubscription(ctx)
	if err != nil {
		return nil, err
	}
	if err := subscriptionmwcli.DeleteSubscription(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
