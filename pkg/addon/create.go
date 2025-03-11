package addon

import (
	"context"

	submwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/addon"
	submwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/addon"

	"github.com/google/uuid"
)

func (h *Handler) CreateAddon(ctx context.Context) (*submwpb.Addon, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := submwcli.CreateAddon(ctx, &submwpb.AddonReq{
		EntID:       h.EntID,
		AppID:       h.AppID,
		Price:       h.Price,
		Credit:      h.Credit,
		SortOrder:   h.SortOrder,
		Enabled:     h.Enabled,
		Description: h.Description,
	}); err != nil {
		return nil, err
	}
	return h.GetAddon(ctx)
}
