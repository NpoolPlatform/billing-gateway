package subscription

import (
	"context"
	"fmt"

	subscriptionmwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/subscription"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	subscriptionmwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/subscription"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkSubscription(ctx context.Context) error {
	exist, err := subscriptionmwcli.ExistSubscriptionConds(ctx, &subscriptionmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid subscription")
	}
	return nil
}
