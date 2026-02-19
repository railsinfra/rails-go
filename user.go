// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/rails-go/internal/apijson"
	"github.com/stainless-sdks/rails-go/internal/requestconfig"
	"github.com/stainless-sdks/rails-go/option"
	"github.com/stainless-sdks/rails-go/packages/param"
	"github.com/stainless-sdks/rails-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the rails API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Create user
func (r *UserService) New(ctx context.Context, params UserNewParams, opts ...option.RequestOption) (res *UserNewResponse, err error) {
	if !param.IsOmitted(params.XEnvironment) {
		opts = append(opts, option.WithHeader("X-Environment", fmt.Sprintf("%v", params.XEnvironment)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type UserNewResponse struct {
	Status string `json:"status,required"`
	UserID string `json:"user_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserNewResponse) RawJSON() string { return r.JSON.raw }
func (r *UserNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	Email     string `json:"email,required" format:"email"`
	FirstName string `json:"first_name,required"`
	LastName  string `json:"last_name,required"`
	Password  string `json:"password,required" format:"password"`
	// Any of "sandbox", "production".
	XEnvironment UserNewParamsXEnvironment `header:"X-Environment,omitzero,required" json:"-"`
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParamsXEnvironment string

const (
	UserNewParamsXEnvironmentSandbox    UserNewParamsXEnvironment = "sandbox"
	UserNewParamsXEnvironmentProduction UserNewParamsXEnvironment = "production"
)
