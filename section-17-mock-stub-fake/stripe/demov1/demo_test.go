package demo

import (
	"fmt"
	"net/http"
	stripe "test-with-go/section-17-mock-stub-fake/stripe/v1"
	"testing"
)

type App struct {
	Stripe *stripe.Client
}

func (a *App) Run() {}

func TestApp(t *testing.T) {
	client, mux, teardown := stripe.TestClient(t)
	defer teardown()

	mux.HandleFunc("/v1/charges", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"id":"ch_1DEjEH2eZvKYlo2CxOmkZL4D","amount":2000,"description":"Charge for demo purposes.","status":"succeeded"}`)
	})

	app := App{
		Stripe: client,
	}
	app.Run()
}
