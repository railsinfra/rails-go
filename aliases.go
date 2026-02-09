// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails

import (
	"github.com/stainless-sdks/rails-go/internal/apierror"
	"github.com/stainless-sdks/rails-go/packages/param"
	"github.com/stainless-sdks/rails-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type Transaction = shared.Transaction

// This is an alias to an internal type.
type TransactionStatus = shared.TransactionStatus

// Equals "pending"
const TransactionStatusPending = shared.TransactionStatusPending

// Equals "completed"
const TransactionStatusCompleted = shared.TransactionStatusCompleted

// Equals "failed"
const TransactionStatusFailed = shared.TransactionStatusFailed

// Equals "cancelled"
const TransactionStatusCancelled = shared.TransactionStatusCancelled

// This is an alias to an internal type.
type TransactionTransactionType = shared.TransactionTransactionType

// Equals "deposit"
const TransactionTransactionTypeDeposit = shared.TransactionTransactionTypeDeposit

// Equals "withdrawal"
const TransactionTransactionTypeWithdrawal = shared.TransactionTransactionTypeWithdrawal

// Equals "transfer"
const TransactionTransactionTypeTransfer = shared.TransactionTransactionTypeTransfer

// Equals "recurring_payment"
const TransactionTransactionTypeRecurringPayment = shared.TransactionTransactionTypeRecurringPayment

// Equals "savings_withdraw"
const TransactionTransactionTypeSavingsWithdraw = shared.TransactionTransactionTypeSavingsWithdraw
