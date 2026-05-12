// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/railsinfra/rails-go"
	"github.com/railsinfra/rails-go/internal/testutil"
	"github.com/railsinfra/rails-go/option"
)

func TestAuditEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.AuditEvents.List(context.TODO(), rails.AuditEventListParams{
		Action:      rails.String("action"),
		Environment: rails.AuditEventListParamsEnvironmentSandbox,
		From:        rails.Time(time.Now()),
		Outcome:     rails.AuditEventListParamsOutcomeSuccess,
		Page:        rails.Int(1),
		PerPage:     rails.Int(1),
		TargetID:    rails.String("target_id"),
		TargetType:  rails.String("target_type"),
		To:          rails.Time(time.Now()),
	})
	if err != nil {
		var apierr *rails.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
