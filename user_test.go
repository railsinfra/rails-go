// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/rails-go"
	"github.com/stainless-sdks/rails-go/internal/testutil"
	"github.com/stainless-sdks/rails-go/option"
)

func TestUserNew(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := rails.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.New(context.TODO(), rails.UserNewParams{
		Email:        "dev@stainless.com",
		FirstName:    "first_name",
		LastName:     "last_name",
		Password:     "password",
		XEnvironment: rails.UserNewParamsXEnvironmentSandbox,
	})
	if err != nil {
		var apierr *rails.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
