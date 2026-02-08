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

	"github.com/stainless-sdks/rails-go/internal/apijson"
	"github.com/stainless-sdks/rails-go/internal/apiquery"
	"github.com/stainless-sdks/rails-go/internal/requestconfig"
	"github.com/stainless-sdks/rails-go/option"
	"github.com/stainless-sdks/rails-go/packages/param"
	"github.com/stainless-sdks/rails-go/packages/respjson"
)

// AccountService contains methods and other services that help with interacting
// with the rails API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	return
}

// Create account
func (r *AccountService) New(ctx context.Context, body AccountNewParams, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve account
func (r *AccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List accounts
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *[]Account, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Close account
func (r *AccountService) Close(ctx context.Context, id string, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Deposit into account
func (r *AccountService) Deposit(ctx context.Context, id string, body AccountDepositParams, opts ...option.RequestOption) (res *AccountDepositResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/accounts/%s/deposit", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Transfer between accounts
func (r *AccountService) Transfer(ctx context.Context, id string, body AccountTransferParams, opts ...option.RequestOption) (res *AccountTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/accounts/%s/transfer", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update account status
func (r *AccountService) UpdateStatus(ctx context.Context, id string, body AccountUpdateStatusParams, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Withdraw from account
func (r *AccountService) Withdraw(ctx context.Context, id string, body AccountWithdrawParams, opts ...option.RequestOption) (res *AccountWithdrawResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/accounts/%s/withdraw", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Account struct {
	ID            string `json:"id,required" format:"uuid"`
	AccountNumber string `json:"account_number,required"`
	// Any of "checking", "saving".
	AccountType AccountAccountType `json:"account_type,required"`
	Balance     string             `json:"balance,required"`
	Currency    string             `json:"currency,required"`
	Environment string             `json:"environment,required"`
	// Any of "active", "suspended", "closed".
	Status         AccountStatus `json:"status,required"`
	UserID         string        `json:"user_id,required" format:"uuid"`
	AdminUserID    string        `json:"admin_user_id,nullable" format:"uuid"`
	CreatedAt      time.Time     `json:"created_at,nullable" format:"date-time"`
	OrganizationID string        `json:"organization_id,nullable" format:"uuid"`
	UpdatedAt      time.Time     `json:"updated_at,nullable" format:"date-time"`
	UserRole       string        `json:"user_role,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccountNumber  respjson.Field
		AccountType    respjson.Field
		Balance        respjson.Field
		Currency       respjson.Field
		Environment    respjson.Field
		Status         respjson.Field
		UserID         respjson.Field
		AdminUserID    respjson.Field
		CreatedAt      respjson.Field
		OrganizationID respjson.Field
		UpdatedAt      respjson.Field
		UserRole       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccountType string

const (
	AccountAccountTypeChecking AccountAccountType = "checking"
	AccountAccountTypeSaving   AccountAccountType = "saving"
)

type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusSuspended AccountStatus = "suspended"
	AccountStatusClosed    AccountStatus = "closed"
)

type Transaction struct {
	ID           string    `json:"id,required" format:"uuid"`
	AccountID    string    `json:"account_id,required" format:"uuid"`
	Amount       string    `json:"amount,required"`
	BalanceAfter string    `json:"balance_after,required"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	Currency     string    `json:"currency,required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status TransactionStatus `json:"status,required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     TransactionTransactionType `json:"transaction_type,required"`
	UpdatedAt           time.Time                  `json:"updated_at,required" format:"date-time"`
	Description         string                     `json:"description,nullable"`
	ExternalRecipientID string                     `json:"external_recipient_id,nullable"`
	RecipientAccountID  string                     `json:"recipient_account_id,nullable" format:"uuid"`
	ReferenceID         string                     `json:"reference_id,nullable" format:"uuid"`
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
func (r Transaction) RawJSON() string { return r.JSON.raw }
func (r *Transaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusCancelled TransactionStatus = "cancelled"
)

type TransactionTransactionType string

const (
	TransactionTransactionTypeDeposit          TransactionTransactionType = "deposit"
	TransactionTransactionTypeWithdrawal       TransactionTransactionType = "withdrawal"
	TransactionTransactionTypeTransfer         TransactionTransactionType = "transfer"
	TransactionTransactionTypeRecurringPayment TransactionTransactionType = "recurring_payment"
	TransactionTransactionTypeSavingsWithdraw  TransactionTransactionType = "savings_withdraw"
)

type AccountDepositResponse struct {
	Account     Account     `json:"account,required"`
	Transaction Transaction `json:"transaction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Transaction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountDepositResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountDepositResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTransferResponse struct {
	FromAccount Account     `json:"from_account,required"`
	ToAccount   Account     `json:"to_account,required"`
	Transaction Transaction `json:"transaction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromAccount respjson.Field
		ToAccount   respjson.Field
		Transaction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountTransferResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountTransferResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWithdrawResponse struct {
	Account     Account     `json:"account,required"`
	Transaction Transaction `json:"transaction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Transaction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountWithdrawResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountWithdrawResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewParams struct {
	// Any of "checking", "saving".
	AccountType    AccountNewParamsAccountType `json:"account_type,omitzero,required"`
	UserID         string                      `json:"user_id,required" format:"uuid"`
	Environment    param.Opt[string]           `json:"environment,omitzero"`
	OrganizationID param.Opt[string]           `json:"organization_id,omitzero" format:"uuid"`
	Currency       param.Opt[string]           `json:"currency,omitzero"`
	paramObj
}

func (r AccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewParamsAccountType string

const (
	AccountNewParamsAccountTypeChecking AccountNewParamsAccountType = "checking"
	AccountNewParamsAccountTypeSaving   AccountNewParamsAccountType = "saving"
)

type AccountListParams struct {
	UserID string `query:"user_id,required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountDepositParams struct {
	Amount      string            `json:"amount,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AccountDepositParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountDepositParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountDepositParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTransferParams struct {
	Amount      string            `json:"amount,required"`
	ToAccountID string            `json:"to_account_id,required" format:"uuid"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AccountTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateStatusParams struct {
	// Any of "active", "suspended", "closed".
	Status AccountUpdateStatusParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r AccountUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateStatusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountUpdateStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateStatusParamsStatus string

const (
	AccountUpdateStatusParamsStatusActive    AccountUpdateStatusParamsStatus = "active"
	AccountUpdateStatusParamsStatusSuspended AccountUpdateStatusParamsStatus = "suspended"
	AccountUpdateStatusParamsStatusClosed    AccountUpdateStatusParamsStatus = "closed"
)

type AccountWithdrawParams struct {
	Amount      string            `json:"amount,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AccountWithdrawParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountWithdrawParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountWithdrawParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
