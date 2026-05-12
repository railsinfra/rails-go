// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rails_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/rails-go"
	"github.com/stainless-sdks/rails-go/packages/param"
)

func TestAuditEventList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	client := rails.NewClient()
	_, err := client.Audit.Events.List(context.TODO(), rails.AuditEventListParams{
		Environment: param.NewOpt("sandbox"),
	})
	if err != nil {
		t.Error(err)
	}
}
