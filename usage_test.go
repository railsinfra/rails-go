// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails_test

import (
	"context"
	"os"
	"testing"

	"github.com/railsinfra/rails-go"
	"github.com/railsinfra/rails-go/internal/testutil"
	"github.com/railsinfra/rails-go/option"
)

func TestUsage(t *testing.T) {
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
	account, err := client.Accounts.New(context.TODO(), rails.AccountNewParams{
		AccountType: rails.AccountNewParamsAccountTypeChecking,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", account.ID)
}
