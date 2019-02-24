package stripe

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Charge struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Client struct {
	Key string
}

func (c *Client) Charge(amount int, source, desc string) (*Charge, error) {
	v := url.Values{}
	v.Set("amount", strconv.Itoa(amount))
	v.Set("currency", "usd")
	v.Set("source", source)
	v.Set("description", desc)
	req, err := http.NewRequest(http.MethodPost, "https://api.stripe.com/v1/charges", strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Key, "")

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var charge Charge
	err = json.Unmarshal(resBytes, &charge)
	if err != nil {
		return nil, err
	}
	return &charge, nil
}
