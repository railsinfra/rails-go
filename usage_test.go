// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails_test

import (
	"context"
	"os"
	"testing"

	"github.com/sibabale/rails-go"
	"github.com/sibabale/rails-go/internal/testutil"
	"github.com/sibabale/rails-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	user, err := client.Users.New(context.TODO(), rails.UserNewParams{
		Email:        "dev@stainless.com",
		FirstName:    "first_name",
		LastName:     "last_name",
		Password:     "password",
		XEnvironment: rails.UserNewParamsXEnvironmentSandbox,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", user.UserID)
}
