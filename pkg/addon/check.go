package addon

import (
	"context"
	"fmt"

	addonmwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/addon"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	addonmwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkAddon(ctx context.Context) error {
	exist, err := addonmwcli.ExistAddonConds(ctx, &addonmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid addon")
	}
	return nil
}
