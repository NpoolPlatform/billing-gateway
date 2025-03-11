package subscription

import (
	"context"

	submwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/user/subscription"
	submwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
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
		ID:                 h.ID,
		EntID:              h.EntID,
		StartAt:            h.StartAt,
		EndAt:              h.EndAt,
		UsageState:         h.UsageState,
		SubscriptionCredit: h.SubscriptionCredit,
		AddonCredit:        h.AddonCredit,
	}); err != nil {
		return nil, err
	}
	return h.GetSubscription(ctx)
}
