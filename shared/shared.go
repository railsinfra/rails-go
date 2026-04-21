// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/railsinfra/rails-go/internal/apijson"
	"github.com/railsinfra/rails-go/packages/param"
	"github.com/railsinfra/rails-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Transaction struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AccountID    string    `json:"account_id" api:"required" format:"uuid"`
	Amount       string    `json:"amount" api:"required"`
	BalanceAfter string    `json:"balance_after" api:"required"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency     string    `json:"currency" api:"required"`
	// Any of "pending", "completed", "failed", "cancelled".
	Status TransactionStatus `json:"status" api:"required"`
	// Any of "deposit", "withdrawal", "transfer", "recurring_payment",
	// "savings_withdraw".
	TransactionType     TransactionTransactionType `json:"transaction_type" api:"required"`
	UpdatedAt           time.Time                  `json:"updated_at" api:"required" format:"date-time"`
	Description         string                     `json:"description" api:"nullable"`
	ExternalRecipientID string                     `json:"external_recipient_id" api:"nullable"`
	RecipientAccountID  string                     `json:"recipient_account_id" api:"nullable" format:"uuid"`
	ReferenceID         string                     `json:"reference_id" api:"nullable" format:"uuid"`
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
