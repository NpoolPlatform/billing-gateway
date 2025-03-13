package charge

import (
	"context"
	"fmt"

	exchangemwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/credit/exchange"
	recordmwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/user/credit/record"
	submwcli "github.com/NpoolPlatform/billing-middleware/pkg/client/user/subscription"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/billing/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	exchangemwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
	recordmwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/credit/record"
	submwpb "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription"
)

type calculateHandler struct {
	*Handler
	consumeCredit    uint32
	exchange         *exchangemwpb.Exchange
	userSubscription *submwpb.Subscription
	extra            string
}

func (h *calculateHandler) calculateByToken() {
	// pre token
	const preTokenCount = uint32(2000)
	h.consumeCredit = h.exchange.Credit * preTokenCount
	h.extra = fmt.Sprintf("prededucted %v tokens", preTokenCount)
}

func (h *calculateHandler) calculateByCount() {
	// calculate credit
	h.consumeCredit = h.exchange.Credit / h.exchange.ExchangeThreshold
}

func (h *calculateHandler) calculateByFilePage() {
	// todo get file pages
	pages := uint32(1)
	h.consumeCredit = (h.exchange.Credit / h.exchange.ExchangeThreshold) * pages
}

func (h *calculateHandler) getExchange(ctx context.Context) error {
	conds := &exchangemwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		Path:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.Path},
	}
	exchange, err := exchangemwcli.GetExchangeOnly(ctx, conds)
	if err != nil {
		return err
	}
	h.exchange = exchange
	return nil
}

func (h *calculateHandler) getUserSubscription(ctx context.Context) error {
	sub, err := submwcli.GetSubscriptionOnly(ctx, &submwpb.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	})
	if err != nil {
		return err
	}
	h.userSubscription = sub
	return nil
}

func (h *calculateHandler) calculateCredit() (ok bool, subDeducted, addonDeducted uint32, err error) {
	if h.consumeCredit == 0 {
		return false, 0, 0, nil
	}
	if h.consumeCredit > (h.userSubscription.SubscriptionCredit + h.userSubscription.AddonCredit) {
		return false, 0, 0, wlog.Errorf("insufficient credit balance")
	}

	if h.userSubscription.SubscriptionCredit > h.consumeCredit {
		subDeducted = h.consumeCredit
		return true, subDeducted, 0, nil
	}

	subDeducted = h.userSubscription.SubscriptionCredit
	addonDeducted = h.consumeCredit - subDeducted

	return true, subDeducted, addonDeducted, nil
}

func (h *calculateHandler) deductedCredit(ctx context.Context, subDeducted, addonDeducted uint32) error {
	subBalance := h.userSubscription.SubscriptionCredit - subDeducted
	addonBalance := h.userSubscription.AddonCredit - addonDeducted
	if err := submwcli.UpdateSubscription(ctx, &submwpb.SubscriptionReq{
		ID:                 h.ID,
		EntID:              h.EntID,
		SubscriptionCredit: &subBalance,
		AddonCredit:        &addonBalance,
	}); err != nil {
		return err
	}

	operationType := types.OperationType_Exchange
	creditsChange := 0 - int32(h.consumeCredit)
	if err := recordmwcli.CreateRecord(ctx, &recordmwpb.RecordReq{
		AppID:         h.AppID,
		UserID:        h.UserID,
		OperationType: &operationType,
		CreditsChange: &creditsChange,
		Extra:         &h.extra,
	}); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CalculateCharge(ctx context.Context) (allow bool, msg string, err error) {
	allow = false
	handler := calculateHandler{
		Handler:       h,
		consumeCredit: uint32(0),
	}
	// get exchange usagetype by path
	if err := handler.getExchange(ctx); err != nil {
		return false, "", err
	}

	switch handler.exchange.UsageType {
	case types.UsageType_ChatCount:
		fallthrough //nolint
	case types.UsageType_ImageCount:
		fallthrough //nolint
	case types.UsageType_VideoCount:
		handler.calculateByCount()
	case types.UsageType_TextToken:
		handler.calculateByToken()
	case types.UsageType_FilePageCount:
		handler.calculateByFilePage()
	default:
		return false, "", wlog.Errorf("invalid usagetype")
	}

	if err := handler.getUserSubscription(ctx); err != nil {
		return false, "", err
	}
	ok, subDeducted, addonDeducted, err := handler.calculateCredit()
	if err != nil {
		return false, "", err
	}
	if !ok {
		return true, "", nil
	}

	if err := handler.deductedCredit(ctx, subDeducted, addonDeducted); err != nil {
		return false, "", err
	}
	allow = true

	return allow, msg, err
}
