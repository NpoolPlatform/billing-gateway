package exchange

import (
	"context"

	submwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/credit/exchange"
	submwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateExchange(ctx context.Context) (*submwpb.Exchange, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkExchange(ctx); err != nil {
		return nil, err
	}
	if err := submwcli.UpdateExchange(ctx, &submwpb.ExchangeReq{
		ID:                h.ID,
		EntID:             h.EntID,
		Credit:            h.Credit,
		ExchangeThreshold: h.ExchangeThreshold,
	}); err != nil {
		return nil, err
	}
	return h.GetExchange(ctx)
}
