// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails

import (
	"context"
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

// Audit events
//
// AuditEventService contains methods and other services that help with interacting
// with the rails API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditEventService] method instead.
type AuditEventService struct {
	Options []option.RequestOption
}

// NewAuditEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditEventService(opts ...option.RequestOption) (r AuditEventService) {
	r = AuditEventService{}
	r.Options = opts
	return
}

// List audit events
func (r *AuditEventService) List(ctx context.Context, query AuditEventListParams, opts ...option.RequestOption) (res *AuditEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/audit/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AuditEventListResponse struct {
	Data       []AuditEventListResponseData     `json:"data" api:"required"`
	Pagination AuditEventListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "users.business.register", "users.auth.login", "users.auth.refresh",
	// "users.auth.revoke", "users.password_reset.request",
	// "users.password_reset.complete", "users.beta.apply", "users.api_key.create",
	// "users.api_key.revoke", "accounts.account.create",
	// "accounts.account.update_status", "accounts.account.close",
	// "accounts.money.deposit", "accounts.money.withdraw", "accounts.money.transfer",
	// "ledger.transaction.post".
	Action        string                          `json:"action" api:"required"`
	Actor         AuditEventListResponseDataActor `json:"actor" api:"required"`
	CorrelationID string                          `json:"correlation_id" api:"required"`
	CreatedAt     time.Time                       `json:"created_at" api:"required" format:"date-time"`
	// Any of "sandbox", "production".
	Environment    string            `json:"environment" api:"required"`
	Metadata       map[string]string `json:"metadata" api:"required"`
	OccurredAt     time.Time         `json:"occurred_at" api:"required" format:"date-time"`
	OrganizationID string            `json:"organization_id" api:"required" format:"uuid"`
	// Any of "success", "client_error", "server_error".
	Outcome string                            `json:"outcome" api:"required"`
	Request AuditEventListResponseDataRequest `json:"request" api:"required"`
	// Any of 1.
	SchemaVersion int64 `json:"schema_version" api:"required"`
	// Any of "users", "accounts", "ledger".
	SourceService string                           `json:"source_service" api:"required"`
	Target        AuditEventListResponseDataTarget `json:"target" api:"required"`
	Reason        string                           `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Action         respjson.Field
		Actor          respjson.Field
		CorrelationID  respjson.Field
		CreatedAt      respjson.Field
		Environment    respjson.Field
		Metadata       respjson.Field
		OccurredAt     respjson.Field
		OrganizationID respjson.Field
		Outcome        respjson.Field
		Request        respjson.Field
		SchemaVersion  respjson.Field
		SourceService  respjson.Field
		Target         respjson.Field
		Reason         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListResponseDataActor struct {
	ID string `json:"id" api:"required"`
	// Any of "user", "api_key", "internal_service", "anonymous".
	Type  string   `json:"type" api:"required"`
	Roles []string `json:"roles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Roles       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponseDataActor) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponseDataActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListResponseDataRequest struct {
	ID        string `json:"id" api:"required"`
	Method    string `json:"method" api:"required"`
	Path      string `json:"path" api:"required"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Method      respjson.Field
		Path        respjson.Field
		IP          respjson.Field
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponseDataRequest) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponseDataRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListResponseDataTarget struct {
	ID   string `json:"id" api:"required"`
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponseDataTarget) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponseDataTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListResponsePagination struct {
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
func (r AuditEventListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListParams struct {
	Action     param.Opt[string]    `query:"action,omitzero" json:"-"`
	From       param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	Page       param.Opt[int64]     `query:"page,omitzero" json:"-"`
	PerPage    param.Opt[int64]     `query:"per_page,omitzero" json:"-"`
	TargetID   param.Opt[string]    `query:"target_id,omitzero" json:"-"`
	TargetType param.Opt[string]    `query:"target_type,omitzero" json:"-"`
	To         param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Environment to list audit events from. Defaults to sandbox when omitted.
	//
	// Any of "sandbox", "production".
	Environment AuditEventListParamsEnvironment `query:"environment,omitzero" json:"-"`
	// Any of "success", "client_error", "server_error".
	Outcome AuditEventListParamsOutcome `query:"outcome,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditEventListParams]'s query parameters as `url.Values`.
func (r AuditEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Environment to list audit events from. Defaults to sandbox when omitted.
type AuditEventListParamsEnvironment string

const (
	AuditEventListParamsEnvironmentSandbox    AuditEventListParamsEnvironment = "sandbox"
	AuditEventListParamsEnvironmentProduction AuditEventListParamsEnvironment = "production"
)

type AuditEventListParamsOutcome string

const (
	AuditEventListParamsOutcomeSuccess     AuditEventListParamsOutcome = "success"
	AuditEventListParamsOutcomeClientError AuditEventListParamsOutcome = "client_error"
	AuditEventListParamsOutcomeServerError AuditEventListParamsOutcome = "server_error"
)
