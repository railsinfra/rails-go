// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails

import (
	"context"
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

// AuditService contains methods and other services that help with interacting
// with the rails API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditService] method instead.
type AuditService struct {
	Options []option.RequestOption
	Events  AuditEventService
}

// NewAuditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuditService(opts ...option.RequestOption) (r AuditService) {
	r = AuditService{}
	r.Options = opts
	r.Events = NewAuditEventService(opts...)
	return
}

// AuditEventService contains methods and other services that help with
// interacting with the rails API.
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

// List audit events for the authenticated business.
func (r *AuditEventService) List(ctx context.Context, query AuditEventListParams, opts ...option.RequestOption) (res *AuditEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/audit/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AuditEventListResponse struct {
	Data       []AuditEventListResponseData     `json:"data,required"`
	Pagination AuditEventListResponsePagination `json:"pagination,required"`
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
	ID             string                            `json:"id,required" format:"uuid"`
	Action         string                            `json:"action,required"`
	Actor          AuditEventListResponseDataActor   `json:"actor,required"`
	CreatedAt      time.Time                         `json:"created_at,required" format:"date-time"`
	Environment    string                            `json:"environment,required"`
	Metadata       map[string]any                    `json:"metadata,required"`
	OccurredAt     time.Time                         `json:"occurred_at,required" format:"date-time"`
	OrganizationID string                            `json:"organization_id,required" format:"uuid"`
	Outcome        string                            `json:"outcome,required"`
	Request        AuditEventListResponseDataRequest `json:"request,required"`
	SchemaVersion  int64                             `json:"schema_version,required"`
	SourceService  string                            `json:"source_service,required"`
	Target         AuditEventListResponseDataTarget  `json:"target,required"`
	CorrelationID  string                            `json:"correlation_id,nullable"`
	Reason         string                            `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Action         respjson.Field
		Actor          respjson.Field
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
		CorrelationID  respjson.Field
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
	ID    string   `json:"id,required"`
	Roles []string `json:"roles"`
	Type  string   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Roles       respjson.Field
		Type        respjson.Field
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
	ID        string `json:"id,required"`
	IP        string `json:"ip"`
	Method    string `json:"method,required"`
	Path      string `json:"path,required"`
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IP          respjson.Field
		Method      respjson.Field
		Path        respjson.Field
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
	ID   string `json:"id,required"`
	Type string `json:"type,required"`
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
	Page       int64 `json:"page,required"`
	PerPage    int64 `json:"per_page,required"`
	TotalCount int64 `json:"total_count,required"`
	TotalPages int64 `json:"total_pages,required"`
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
	Action      param.Opt[string]    `query:"action,omitzero" json:"-"`
	Environment param.Opt[string]    `query:"environment,omitzero" json:"-"`
	From        param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	Outcome     param.Opt[string]    `query:"outcome,omitzero" json:"-"`
	Page        param.Opt[int64]     `query:"page,omitzero" json:"-"`
	PerPage     param.Opt[int64]     `query:"per_page,omitzero" json:"-"`
	TargetID    param.Opt[string]    `query:"target_id,omitzero" json:"-"`
	TargetType  param.Opt[string]    `query:"target_type,omitzero" json:"-"`
	To          param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [AuditEventListParams]'s query parameters as `url.Values`.
func (r AuditEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
