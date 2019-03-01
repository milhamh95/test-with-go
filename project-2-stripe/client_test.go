package stripe_test

import (
	"strings"
	stripe "test-with-go/project-2-stripe"
	"testing"
)

func TestClient_Customer(t *testing.T) {
	c := stripe.Client{
		Key: "sk_test_4eC39HqLyjWDarjtT1zdp7dc",
	}
	tok := "tok_amex"
	cus, err := c.Customer(tok)
	if err != nil {
		t.Errorf("Customer() err = %v; want %v", err, nil)
	}
	if cus == nil {
		t.Fatalf("Customer() = nil; want non-nil value")
	}
	if !strings.HasPrefix(cus.ID, "cus_") {
		t.Errorf("Customer() ID = %s; want prefix %q", cus.ID, "cus_")
	}
}
