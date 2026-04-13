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
func (r *AccountService) New(ctx context.Context, body AccountNewParams, opts ...option.RequestOption) (res *AccountNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve account
func (r *AccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
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
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *[]AccountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Close account
func (r *AccountService) Close(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountCloseResponse, err error) {
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
func (r *AccountService) Deposit(ctx context.Context, id string, body AccountDepositParams, opts ...option.RequestOption) (res *AccountDepositResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s/deposit", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transfer between accounts
func (r *AccountService) Transfer(ctx context.Context, id string, body AccountTransferParams, opts ...option.RequestOption) (res *AccountTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s/transfer", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update account status
func (r *AccountService) UpdateStatus(ctx context.Context, id string, body AccountUpdateStatusParams, opts ...option.RequestOption) (res *AccountUpdateStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Withdraw from account
func (r *AccountService) Withdraw(ctx context.Context, id string, body AccountWithdrawParams, opts ...option.RequestOption) (res *AccountWithdrawResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/accounts/%s/withdraw", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AccountNewResponse struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType AccountNewResponseAccountType `json:"account_type" api:"required"`
	Balance     string                        `json:"balance" api:"required"`
	Currency    string                        `json:"currency" api:"required"`
	Environment string                        `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         AccountNewResponseStatus `json:"status" api:"required"`
	UserID         string                   `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string                   `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time                `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string                   `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time                `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string                   `json:"user_role" api:"nullable"`
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
func (r AccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewResponseAccountType string

const (
	AccountNewResponseAccountTypeChecking AccountNewResponseAccountType = "checking"
	AccountNewResponseAccountTypeSaving   AccountNewResponseAccountType = "saving"
)

type AccountNewResponseStatus string

const (
	AccountNewResponseStatusActive    AccountNewResponseStatus = "active"
	AccountNewResponseStatusSuspended AccountNewResponseStatus = "suspended"
	AccountNewResponseStatusClosed    AccountNewResponseStatus = "closed"
)

type AccountGetResponse struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType AccountGetResponseAccountType `json:"account_type" api:"required"`
	Balance     string                        `json:"balance" api:"required"`
	Currency    string                        `json:"currency" api:"required"`
	Environment string                        `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         AccountGetResponseStatus `json:"status" api:"required"`
	UserID         string                   `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string                   `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time                `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string                   `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time                `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string                   `json:"user_role" api:"nullable"`
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
func (r AccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponseAccountType string

const (
	AccountGetResponseAccountTypeChecking AccountGetResponseAccountType = "checking"
	AccountGetResponseAccountTypeSaving   AccountGetResponseAccountType = "saving"
)

type AccountGetResponseStatus string

const (
	AccountGetResponseStatusActive    AccountGetResponseStatus = "active"
	AccountGetResponseStatusSuspended AccountGetResponseStatus = "suspended"
	AccountGetResponseStatusClosed    AccountGetResponseStatus = "closed"
)

type AccountListResponse struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType AccountListResponseAccountType `json:"account_type" api:"required"`
	Balance     string                         `json:"balance" api:"required"`
	Currency    string                         `json:"currency" api:"required"`
	Environment string                         `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         AccountListResponseStatus `json:"status" api:"required"`
	UserID         string                    `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string                    `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time                 `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string                    `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time                 `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string                    `json:"user_role" api:"nullable"`
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
func (r AccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseAccountType string

const (
	AccountListResponseAccountTypeChecking AccountListResponseAccountType = "checking"
	AccountListResponseAccountTypeSaving   AccountListResponseAccountType = "saving"
)

type AccountListResponseStatus string

const (
	AccountListResponseStatusActive    AccountListResponseStatus = "active"
	AccountListResponseStatusSuspended AccountListResponseStatus = "suspended"
	AccountListResponseStatusClosed    AccountListResponseStatus = "closed"
)

type AccountCloseResponse struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType AccountCloseResponseAccountType `json:"account_type" api:"required"`
	Balance     string                          `json:"balance" api:"required"`
	Currency    string                          `json:"currency" api:"required"`
	Environment string                          `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         AccountCloseResponseStatus `json:"status" api:"required"`
	UserID         string                     `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string                     `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time                  `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string                     `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time                  `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string                     `json:"user_role" api:"nullable"`
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
func (r AccountCloseResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountCloseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCloseResponseAccountType string

const (
	AccountCloseResponseAccountTypeChecking AccountCloseResponseAccountType = "checking"
	AccountCloseResponseAccountTypeSaving   AccountCloseResponseAccountType = "saving"
)

type AccountCloseResponseStatus string

const (
	AccountCloseResponseStatusActive    AccountCloseResponseStatus = "active"
	AccountCloseResponseStatusSuspended AccountCloseResponseStatus = "suspended"
	AccountCloseResponseStatusClosed    AccountCloseResponseStatus = "closed"
)

type AccountDepositResponse struct {
	Account     AccountDepositResponseAccount     `json:"account" api:"required"`
	Transaction AccountDepositResponseTransaction `json:"transaction" api:"required"`
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

type AccountDepositResponseAccount struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType string `json:"account_type" api:"required"`
	Balance     string `json:"balance" api:"required"`
	Currency    string `json:"currency" api:"required"`
	Environment string `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         string    `json:"status" api:"required"`
	UserID         string    `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string    `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string    `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string    `json:"user_role" api:"nullable"`
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
func (r AccountDepositResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *AccountDepositResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDepositResponseTransaction struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AccountID    string    `json:"account_id" api:"required" format:"uuid"`
	Amount       string    `json:"amount" api:"required"`
	BalanceAfter string    `json:"balance_after" api:"required"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency     string    `json:"currency" api:"required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     string    `json:"transaction_type" api:"required"`
	UpdatedAt           time.Time `json:"updated_at" api:"required" format:"date-time"`
	Description         string    `json:"description" api:"nullable"`
	ExternalRecipientID string    `json:"external_recipient_id" api:"nullable"`
	RecipientAccountID  string    `json:"recipient_account_id" api:"nullable" format:"uuid"`
	ReferenceID         string    `json:"reference_id" api:"nullable" format:"uuid"`
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
func (r AccountDepositResponseTransaction) RawJSON() string { return r.JSON.raw }
func (r *AccountDepositResponseTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTransferResponse struct {
	FromAccount AccountTransferResponseFromAccount `json:"from_account" api:"required"`
	ToAccount   AccountTransferResponseToAccount   `json:"to_account" api:"required"`
	Transaction AccountTransferResponseTransaction `json:"transaction" api:"required"`
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

type AccountTransferResponseFromAccount struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType string `json:"account_type" api:"required"`
	Balance     string `json:"balance" api:"required"`
	Currency    string `json:"currency" api:"required"`
	Environment string `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         string    `json:"status" api:"required"`
	UserID         string    `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string    `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string    `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string    `json:"user_role" api:"nullable"`
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
func (r AccountTransferResponseFromAccount) RawJSON() string { return r.JSON.raw }
func (r *AccountTransferResponseFromAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTransferResponseToAccount struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType string `json:"account_type" api:"required"`
	Balance     string `json:"balance" api:"required"`
	Currency    string `json:"currency" api:"required"`
	Environment string `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         string    `json:"status" api:"required"`
	UserID         string    `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string    `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string    `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string    `json:"user_role" api:"nullable"`
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
func (r AccountTransferResponseToAccount) RawJSON() string { return r.JSON.raw }
func (r *AccountTransferResponseToAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTransferResponseTransaction struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AccountID    string    `json:"account_id" api:"required" format:"uuid"`
	Amount       string    `json:"amount" api:"required"`
	BalanceAfter string    `json:"balance_after" api:"required"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency     string    `json:"currency" api:"required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     string    `json:"transaction_type" api:"required"`
	UpdatedAt           time.Time `json:"updated_at" api:"required" format:"date-time"`
	Description         string    `json:"description" api:"nullable"`
	ExternalRecipientID string    `json:"external_recipient_id" api:"nullable"`
	RecipientAccountID  string    `json:"recipient_account_id" api:"nullable" format:"uuid"`
	ReferenceID         string    `json:"reference_id" api:"nullable" format:"uuid"`
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
func (r AccountTransferResponseTransaction) RawJSON() string { return r.JSON.raw }
func (r *AccountTransferResponseTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateStatusResponse struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType AccountUpdateStatusResponseAccountType `json:"account_type" api:"required"`
	Balance     string                                 `json:"balance" api:"required"`
	Currency    string                                 `json:"currency" api:"required"`
	Environment string                                 `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         AccountUpdateStatusResponseStatus `json:"status" api:"required"`
	UserID         string                            `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string                            `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time                         `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string                            `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time                         `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string                            `json:"user_role" api:"nullable"`
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
func (r AccountUpdateStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountUpdateStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateStatusResponseAccountType string

const (
	AccountUpdateStatusResponseAccountTypeChecking AccountUpdateStatusResponseAccountType = "checking"
	AccountUpdateStatusResponseAccountTypeSaving   AccountUpdateStatusResponseAccountType = "saving"
)

type AccountUpdateStatusResponseStatus string

const (
	AccountUpdateStatusResponseStatusActive    AccountUpdateStatusResponseStatus = "active"
	AccountUpdateStatusResponseStatusSuspended AccountUpdateStatusResponseStatus = "suspended"
	AccountUpdateStatusResponseStatusClosed    AccountUpdateStatusResponseStatus = "closed"
)

type AccountWithdrawResponse struct {
	Account     AccountWithdrawResponseAccount     `json:"account" api:"required"`
	Transaction AccountWithdrawResponseTransaction `json:"transaction" api:"required"`
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

type AccountWithdrawResponseAccount struct {
	ID            string `json:"id" api:"required" format:"uuid"`
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "saving".
	AccountType string `json:"account_type" api:"required"`
	Balance     string `json:"balance" api:"required"`
	Currency    string `json:"currency" api:"required"`
	Environment string `json:"environment" api:"required"`
	// Any of "active", "suspended", "closed".
	Status         string    `json:"status" api:"required"`
	UserID         string    `json:"user_id" api:"required" format:"uuid"`
	AdminUserID    string    `json:"admin_user_id" api:"nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at" api:"nullable" format:"date-time"`
	OrganizationID string    `json:"organization_id" api:"nullable" format:"uuid"`
	UpdatedAt      time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	UserRole       string    `json:"user_role" api:"nullable"`
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
func (r AccountWithdrawResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *AccountWithdrawResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWithdrawResponseTransaction struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AccountID    string    `json:"account_id" api:"required" format:"uuid"`
	Amount       string    `json:"amount" api:"required"`
	BalanceAfter string    `json:"balance_after" api:"required"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency     string    `json:"currency" api:"required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     string    `json:"transaction_type" api:"required"`
	UpdatedAt           time.Time `json:"updated_at" api:"required" format:"date-time"`
	Description         string    `json:"description" api:"nullable"`
	ExternalRecipientID string    `json:"external_recipient_id" api:"nullable"`
	RecipientAccountID  string    `json:"recipient_account_id" api:"nullable" format:"uuid"`
	ReferenceID         string    `json:"reference_id" api:"nullable" format:"uuid"`
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
func (r AccountWithdrawResponseTransaction) RawJSON() string { return r.JSON.raw }
func (r *AccountWithdrawResponseTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewParams struct {
	// Any of "checking", "saving".
	AccountType AccountNewParamsAccountType `json:"account_type,omitzero" api:"required"`
	// Holder-based: unique per org+env. Requires X-API-Key.
	Email       param.Opt[string] `json:"email,omitzero" format:"email"`
	Environment param.Opt[string] `json:"environment,omitzero"`
	// Holder-based: holder first name.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Holder-based: holder last name.
	LastName       param.Opt[string] `json:"last_name,omitzero"`
	OrganizationID param.Opt[string] `json:"organization_id,omitzero" format:"uuid"`
	// Legacy: platform user ID. Omit when using holder (email + names).
	UserID   param.Opt[string] `json:"user_id,omitzero" format:"uuid"`
	Currency param.Opt[string] `json:"currency,omitzero"`
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
	UserID string `query:"user_id" api:"required" format:"uuid" json:"-"`
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
	Amount      string            `json:"amount" api:"required"`
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
	Amount      string            `json:"amount" api:"required"`
	ToAccountID string            `json:"to_account_id" api:"required" format:"uuid"`
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
	Amount      string            `json:"amount" api:"required"`
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
