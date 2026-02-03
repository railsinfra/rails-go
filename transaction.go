// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/sibabale/rails-go/internal/apijson"
	"github.com/sibabale/rails-go/internal/apiquery"
	"github.com/sibabale/rails-go/internal/requestconfig"
	"github.com/sibabale/rails-go/option"
	"github.com/sibabale/rails-go/packages/param"
	"github.com/sibabale/rails-go/packages/respjson"
)

// TransactionService contains methods and other services that help with
// interacting with the rails API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.Options = opts
	return
}

// Retrieve transaction
func (r *TransactionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TransactionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List account transactions
func (r *TransactionService) ListByAccount(ctx context.Context, accountID string, query TransactionListByAccountParams, opts ...option.RequestOption) (res *[]TransactionListByAccountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/accounts/%s/transactions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TransactionGetResponse struct {
	ID           string    `json:"id,required" format:"uuid"`
	AccountID    string    `json:"account_id,required" format:"uuid"`
	Amount       string    `json:"amount,required"`
	BalanceAfter string    `json:"balance_after,required"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	Currency     string    `json:"currency,required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status TransactionGetResponseStatus `json:"status,required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     TransactionGetResponseTransactionType `json:"transaction_type,required"`
	UpdatedAt           time.Time                             `json:"updated_at,required" format:"date-time"`
	Description         string                                `json:"description,nullable"`
	ExternalRecipientID string                                `json:"external_recipient_id,nullable"`
	RecipientAccountID  string                                `json:"recipient_account_id,nullable" format:"uuid"`
	ReferenceID         string                                `json:"reference_id,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountID           respjson.Field
		Amount              respjson.Field
		BalanceAfter        respjson.Field
		CreatedAt           respjson.Field
		Currency            respjson.Field
		Status              respjson.Field
		TransactionType     respjson.Field
		UpdatedAt           respjson.Field
		Description         respjson.Field
		ExternalRecipientID respjson.Field
		RecipientAccountID  respjson.Field
		ReferenceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionGetResponseStatus string

const (
	TransactionGetResponseStatusPending   TransactionGetResponseStatus = "pending"
	TransactionGetResponseStatusCompleted TransactionGetResponseStatus = "completed"
	TransactionGetResponseStatusFailed    TransactionGetResponseStatus = "failed"
	TransactionGetResponseStatusCancelled TransactionGetResponseStatus = "cancelled"
)

type TransactionGetResponseTransactionType string

const (
	TransactionGetResponseTransactionTypeDeposit          TransactionGetResponseTransactionType = "deposit"
	TransactionGetResponseTransactionTypeWithdrawal       TransactionGetResponseTransactionType = "withdrawal"
	TransactionGetResponseTransactionTypeTransfer         TransactionGetResponseTransactionType = "transfer"
	TransactionGetResponseTransactionTypeRecurringPayment TransactionGetResponseTransactionType = "recurring_payment"
	TransactionGetResponseTransactionTypeSavingsWithdraw  TransactionGetResponseTransactionType = "savings_withdraw"
)

type TransactionListByAccountResponse struct {
	ID           string    `json:"id,required" format:"uuid"`
	AccountID    string    `json:"account_id,required" format:"uuid"`
	Amount       string    `json:"amount,required"`
	BalanceAfter string    `json:"balance_after,required"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	Currency     string    `json:"currency,required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status TransactionListByAccountResponseStatus `json:"status,required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     TransactionListByAccountResponseTransactionType `json:"transaction_type,required"`
	UpdatedAt           time.Time                                       `json:"updated_at,required" format:"date-time"`
	Description         string                                          `json:"description,nullable"`
	ExternalRecipientID string                                          `json:"external_recipient_id,nullable"`
	RecipientAccountID  string                                          `json:"recipient_account_id,nullable" format:"uuid"`
	ReferenceID         string                                          `json:"reference_id,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountID           respjson.Field
		Amount              respjson.Field
		BalanceAfter        respjson.Field
		CreatedAt           respjson.Field
		Currency            respjson.Field
		Status              respjson.Field
		TransactionType     respjson.Field
		UpdatedAt           respjson.Field
		Description         respjson.Field
		ExternalRecipientID respjson.Field
		RecipientAccountID  respjson.Field
		ReferenceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListByAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionListByAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListByAccountResponseStatus string

const (
	TransactionListByAccountResponseStatusPending   TransactionListByAccountResponseStatus = "pending"
	TransactionListByAccountResponseStatusCompleted TransactionListByAccountResponseStatus = "completed"
	TransactionListByAccountResponseStatusFailed    TransactionListByAccountResponseStatus = "failed"
	TransactionListByAccountResponseStatusCancelled TransactionListByAccountResponseStatus = "cancelled"
)

type TransactionListByAccountResponseTransactionType string

const (
	TransactionListByAccountResponseTransactionTypeDeposit          TransactionListByAccountResponseTransactionType = "deposit"
	TransactionListByAccountResponseTransactionTypeWithdrawal       TransactionListByAccountResponseTransactionType = "withdrawal"
	TransactionListByAccountResponseTransactionTypeTransfer         TransactionListByAccountResponseTransactionType = "transfer"
	TransactionListByAccountResponseTransactionTypeRecurringPayment TransactionListByAccountResponseTransactionType = "recurring_payment"
	TransactionListByAccountResponseTransactionTypeSavingsWithdraw  TransactionListByAccountResponseTransactionType = "savings_withdraw"
)

type TransactionListByAccountParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListByAccountParams]'s query parameters as
// `url.Values`.
func (r TransactionListByAccountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
