package app_test

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	app "test-with-go/section-19-testing-with-http/app"
	"testing"

	"golang.org/x/net/publicsuffix"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	app.Home(w, r)

	resp := w.Result()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll() err : %s; want nil", err)
	}

	want := "<h1>Welcome!</h1>"
	got := string(body)

	if want != got {
		t.Errorf("Get / =%s; want %s", got, want)
	}

}

func signedInClient(t *testing.T, baseURL string) *http.Client {
	// our cookiejar will keep and set cookies for us between requests
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		t.Fatalf("cookejar.New() err = %s; wnat nil", err)
	}
	client := &http.Client{
		Jar: jar,
	}

	loginURL := baseURL + "/login"
	req, err := http.NewRequest(http.MethodPost, loginURL, nil)
	if err != nil {
		t.Fatalf("NewRequest() err = %s; want nil", err)
	}
	_, err = client.Do(req)
	if err != nil {
		t.Fatalf("POST /login err = %s; want nil", err)
	}
	t.Logf("Cookies : %v", client.Jar.Cookies(req.URL))
	return client
}
