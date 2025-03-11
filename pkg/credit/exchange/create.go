package exchange

import (
	"context"

	submwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/credit/exchange"
	submwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"

	"github.com/google/uuid"
)

func (h *Handler) CreateExchange(ctx context.Context) (*submwpb.Exchange, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := submwcli.CreateExchange(ctx, &submwpb.ExchangeReq{
		EntID:             h.EntID,
		AppID:             h.AppID,
		UsageType:         h.UsageType,
		Credit:            h.Credit,
		ExchangeThreshold: h.ExchangeThreshold,
	}); err != nil {
		return nil, err
	}
	return h.GetExchange(ctx)
}
