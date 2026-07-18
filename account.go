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
	"github.com/railsinfra/rails-go/shared"
)

// Accounts
//
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
func (r *AccountService) New(ctx context.Context, params AccountNewParams, opts ...option.RequestOption) (res *Account, err error) {
	if !param.IsOmitted(params.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", params.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve account
func (r *AccountService) Get(ctx context.Context, id string, query AccountGetParams, opts ...option.RequestOption) (res *Account, err error) {
	if !param.IsOmitted(query.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", query.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List accounts
func (r *AccountService) List(ctx context.Context, params AccountListParams, opts ...option.RequestOption) (res *[]Account, err error) {
	if !param.IsOmitted(params.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", params.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Close account
func (r *AccountService) Close(ctx context.Context, id string, body AccountCloseParams, opts ...option.RequestOption) (res *Account, err error) {
	if !param.IsOmitted(body.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", body.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Deposit into account
func (r *AccountService) Deposit(ctx context.Context, id string, params AccountDepositParams, opts ...option.RequestOption) (res *AccountDepositResponse, err error) {
	if !param.IsOmitted(params.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", params.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s/deposit", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Transfer between accounts
func (r *AccountService) Transfer(ctx context.Context, id string, params AccountTransferParams, opts ...option.RequestOption) (res *AccountTransferResponse, err error) {
	if !param.IsOmitted(params.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", params.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s/transfer", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update account status
func (r *AccountService) UpdateStatus(ctx context.Context, id string, params AccountUpdateStatusParams, opts ...option.RequestOption) (res *Account, err error) {
	if !param.IsOmitted(params.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", params.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Withdraw from account
func (r *AccountService) Withdraw(ctx context.Context, id string, params AccountWithdrawParams, opts ...option.RequestOption) (res *AccountWithdrawResponse, err error) {
	if !param.IsOmitted(params.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", params.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s/withdraw", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type Account struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType AccountAccountType `json:"account_type" api:"required"`
	Balance     string             `json:"balance" api:"required"`
	Currency    string             `json:"currency" api:"required"`
	Environment string             `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         AccountStatus `json:"status" api:"required"`
	UserID         string        `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string        `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time     `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string        `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time     `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string        `json:"user_role" api:"nullable"`
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

type AccountDepositResponse struct {
	Account     Account            `json:"account" api:"required"`
	Transaction shared.Transaction `json:"transaction" api:"required"`
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
	FromAccount Account            `json:"from_account" api:"required"`
	ToAccount   Account            `json:"to_account" api:"required"`
	Transaction shared.Transaction `json:"transaction" api:"required"`
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
	Account     Account            `json:"account" api:"required"`
	Transaction shared.Transaction `json:"transaction" api:"required"`
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
	AccountType AccountNewParamsAccountType `json:"account_type,omitzero" api:"required"`
	// Three-letter uppercase ISO currency code, for example USD or ZAR.
	Currency       string            `json:"currency" api:"required"`
	UserID         string            `json:"user_id" api:"required" format:"uuid"`
	Environment    param.Opt[string] `json:"environment,omitzero"`
	OrganizationID param.Opt[string] `json:"organization_id,omitzero" format:"uuid"`
	// Any of "sandbox", "production".
	XEnvironment AccountNewParamsXEnvironment `header:"X-Environment,omitzero" json:"-"`
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

type AccountNewParamsXEnvironment string

const (
	AccountNewParamsXEnvironmentSandbox    AccountNewParamsXEnvironment = "sandbox"
	AccountNewParamsXEnvironmentProduction AccountNewParamsXEnvironment = "production"
)

type AccountGetParams struct {
	// Any of "sandbox", "production".
	XEnvironment AccountGetParamsXEnvironment `header:"X-Environment,omitzero" json:"-"`
	paramObj
}

type AccountGetParamsXEnvironment string

const (
	AccountGetParamsXEnvironmentSandbox    AccountGetParamsXEnvironment = "sandbox"
	AccountGetParamsXEnvironmentProduction AccountGetParamsXEnvironment = "production"
)

type AccountListParams struct {
	UserID string `query:"user_id" api:"required" format:"uuid" json:"-"`
	// Any of "sandbox", "production".
	XEnvironment AccountListParamsXEnvironment `header:"X-Environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountListParamsXEnvironment string

const (
	AccountListParamsXEnvironmentSandbox    AccountListParamsXEnvironment = "sandbox"
	AccountListParamsXEnvironmentProduction AccountListParamsXEnvironment = "production"
)

type AccountCloseParams struct {
	// Any of "sandbox", "production".
	XEnvironment AccountCloseParamsXEnvironment `header:"X-Environment,omitzero" json:"-"`
	paramObj
}

type AccountCloseParamsXEnvironment string

const (
	AccountCloseParamsXEnvironmentSandbox    AccountCloseParamsXEnvironment = "sandbox"
	AccountCloseParamsXEnvironmentProduction AccountCloseParamsXEnvironment = "production"
)

type AccountDepositParams struct {
	Amount      string            `json:"amount" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Any of "sandbox", "production".
	XEnvironment AccountDepositParamsXEnvironment `header:"X-Environment,omitzero" json:"-"`
	paramObj
}

func (r AccountDepositParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountDepositParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountDepositParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDepositParamsXEnvironment string

const (
	AccountDepositParamsXEnvironmentSandbox    AccountDepositParamsXEnvironment = "sandbox"
	AccountDepositParamsXEnvironmentProduction AccountDepositParamsXEnvironment = "production"
)

type AccountTransferParams struct {
	Amount      string            `json:"amount" api:"required"`
	ToAccountID string            `json:"to_account_id" api:"required" format:"uuid"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Any of "sandbox", "production".
	XEnvironment AccountTransferParamsXEnvironment `header:"X-Environment,omitzero" json:"-"`
	paramObj
}

func (r AccountTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTransferParamsXEnvironment string

const (
	AccountTransferParamsXEnvironmentSandbox    AccountTransferParamsXEnvironment = "sandbox"
	AccountTransferParamsXEnvironmentProduction AccountTransferParamsXEnvironment = "production"
)

type AccountUpdateStatusParams struct {
	// Any of "active", "suspended", "closed".
	Status AccountUpdateStatusParamsStatus `json:"status,omitzero"`
	// Any of "sandbox", "production".
	XEnvironment AccountUpdateStatusParamsXEnvironment `header:"X-Environment,omitzero" json:"-"`
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

type AccountUpdateStatusParamsXEnvironment string

const (
	AccountUpdateStatusParamsXEnvironmentSandbox    AccountUpdateStatusParamsXEnvironment = "sandbox"
	AccountUpdateStatusParamsXEnvironmentProduction AccountUpdateStatusParamsXEnvironment = "production"
)

type AccountWithdrawParams struct {
	Amount      string            `json:"amount" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Any of "sandbox", "production".
	XEnvironment AccountWithdrawParamsXEnvironment `header:"X-Environment,omitzero" json:"-"`
	paramObj
}

func (r AccountWithdrawParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountWithdrawParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountWithdrawParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWithdrawParamsXEnvironment string

const (
	AccountWithdrawParamsXEnvironmentSandbox    AccountWithdrawParamsXEnvironment = "sandbox"
	AccountWithdrawParamsXEnvironmentProduction AccountWithdrawParamsXEnvironment = "production"
)
