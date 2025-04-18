package exchange

import (
	"context"

	submwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/credit/exchange"
	submwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

func (h *Handler) GetExchange(ctx context.Context) (*submwpb.Exchange, error) {
	return submwcli.GetExchange(ctx, *h.EntID)
}

func (h *Handler) GetExchanges(ctx context.Context) ([]*submwpb.Exchange, error) {
	return submwcli.GetExchanges(ctx, &submwpb.Conds{}, h.Offset, h.Limit)
}

func (h *Handler) GetExchangesCount(ctx context.Context) (uint32, error) {
	return submwcli.GetExchangesCount(ctx, &submwpb.Conds{})
}
