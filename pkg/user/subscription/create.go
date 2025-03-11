package subscription

import (
	"context"

	submwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/user/subscription"
	submwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"

	"github.com/google/uuid"
)

func (h *Handler) CreateSubscription(ctx context.Context) (*submwpb.Subscription, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := submwcli.CreateSubscription(ctx, &submwpb.SubscriptionReq{
		EntID:              h.EntID,
		AppID:              h.AppID,
		UserID:             h.UserID,
		PackageID:          h.PackageID,
		StartAt:            h.StartAt,
		EndAt:              h.EndAt,
		OrderID:            h.OrderID,
		UsageState:         h.UsageState,
		SubscriptionCredit: h.SubscriptionCredit,
		AddonCredit:        h.AddonCredit,
	}); err != nil {
		return nil, err
	}
	return h.GetSubscription(ctx)
}
