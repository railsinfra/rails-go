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

	"github.com/railsinfra/rails-go/internal/apijson"
	"github.com/railsinfra/rails-go/internal/apiquery"
	"github.com/railsinfra/rails-go/internal/requestconfig"
	"github.com/railsinfra/rails-go/option"
	"github.com/railsinfra/rails-go/packages/param"
	"github.com/railsinfra/rails-go/packages/respjson"
)

// Transactions
//
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
		return nil, err
	}
	path := fmt.Sprintf("api/v1/transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List transactions by organization
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *TransactionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List account transactions
func (r *TransactionService) ListByAccount(ctx context.Context, accountID string, query TransactionListByAccountParams, opts ...option.RequestOption) (res *[]TransactionListByAccountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s/transactions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type TransactionGetResponse struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AccountID    string    `json:"account_id" api:"required" format:"uuid"`
	Amount       string    `json:"amount" api:"required"`
	BalanceAfter string    `json:"balance_after" api:"required"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency     string    `json:"currency" api:"required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status TransactionGetResponseStatus `json:"status" api:"required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     TransactionGetResponseTransactionType `json:"transaction_type" api:"required"`
	UpdatedAt           time.Time                             `json:"updated_at" api:"required" format:"date-time"`
	Description         string                                `json:"description" api:"nullable"`
	ExternalRecipientID string                                `json:"external_recipient_id" api:"nullable"`
	RecipientAccountID  string                                `json:"recipient_account_id" api:"nullable" format:"uuid"`
	ReferenceID         string                                `json:"reference_id" api:"nullable" format:"uuid"`
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

type TransactionListResponse struct {
	Data       []TransactionListResponseData     `json:"data" api:"required"`
	Pagination TransactionListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transaction as returned by list-by-organization (organization_id,
// from/to_account_id, transaction_kind).
type TransactionListResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Amount in minor units
	Amount         int64     `json:"amount" api:"required"`
	CreatedAt      time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency       string    `json:"currency" api:"required"`
	FromAccountID  string    `json:"from_account_id" api:"required" format:"uuid"`
	OrganizationID string    `json:"organization_id" api:"required" format:"uuid"`
	// Any of "pending", "posted", "failed".
	Status      string `json:"status" api:"required"`
	ToAccountID string `json:"to_account_id" api:"required" format:"uuid"`
	// Any of "deposit", "withdraw", "transfer".
	TransactionKind string    `json:"transaction_kind" api:"required"`
	UpdatedAt       time.Time `json:"updated_at" api:"required" format:"date-time"`
	Environment     string    `json:"environment" api:"nullable"`
	FailureReason   string    `json:"failure_reason" api:"nullable"`
	IdempotencyKey  string    `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Amount          respjson.Field
		CreatedAt       respjson.Field
		Currency        respjson.Field
		FromAccountID   respjson.Field
		OrganizationID  respjson.Field
		Status          respjson.Field
		ToAccountID     respjson.Field
		TransactionKind respjson.Field
		UpdatedAt       respjson.Field
		Environment     respjson.Field
		FailureReason   respjson.Field
		IdempotencyKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListResponsePagination struct {
	Page       int64 `json:"page" api:"required"`
	PerPage    int64 `json:"per_page" api:"required"`
	TotalCount int64 `json:"total_count" api:"required"`
	TotalPages int64 `json:"total_pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		PerPage     respjson.Field
		TotalCount  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListByAccountResponse struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AccountID    string    `json:"account_id" api:"required" format:"uuid"`
	Amount       string    `json:"amount" api:"required"`
	BalanceAfter string    `json:"balance_after" api:"required"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency     string    `json:"currency" api:"required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status TransactionListByAccountResponseStatus `json:"status" api:"required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     TransactionListByAccountResponseTransactionType `json:"transaction_type" api:"required"`
	UpdatedAt           time.Time                                       `json:"updated_at" api:"required" format:"date-time"`
	Description         string                                          `json:"description" api:"nullable"`
	ExternalRecipientID string                                          `json:"external_recipient_id" api:"nullable"`
	RecipientAccountID  string                                          `json:"recipient_account_id" api:"nullable" format:"uuid"`
	ReferenceID         string                                          `json:"reference_id" api:"nullable" format:"uuid"`
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

type TransactionListParams struct {
	OrganizationID string           `query:"organization_id" api:"required" format:"uuid" json:"-"`
	Page           param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage        param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
