// Package conekta contains Conekta REST API wrappers.
/*
	Use of this source code is governed by a MIT
	license that can be found in the LICENSE file.
*/
package conekta

import "net/http"

const (
	// Host host URL.
	Host = "https://api.conekta.io"

	// OrderURI Order REST.
	OrderURI = "/orders"
)

// Client implements all Conekta REST API endpoints.
type Client struct {
	client      *http.Client
	key, secret string
	sandbox     bool
}

// NewClient returns a new Conekta client with http default settings. Use sandbox
// flag as false for production environments.
func NewClient(key, secret string, sandbox bool) (*Client, error) {
	c := &Client{
		key:     key,
		secret:  secret,
		sandbox: sandbox,
	}
	return c, nil
}

// CreateOrder _
func (c *Client) CreateOrder(or *OrderRequest) (*Order, error) {
	req, err := http.NewRequest("POST", Host+OrderURI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.conekta-v2.0.0+json")
	req.Header.Set("Content-Type", "application/json")

	order := &Order{}
	return order, nil
}

const (
	// MXN mexican pesos code ISO 4217.
	MXN = "MXN"
)

// OrderRequest type.
// see: https://developers.conekta.com/api?language=bash#order
type OrderRequest struct {
	Currency        string          `json:"currency"`
	LineItems       []Product       `json:"line_items"`
	ShippingLines   []Shipping      `json:"shipping_lines"`
	TaxLines        []Tax           `json:"tax_lines"`
	DiscountLines   []Discount      `json:"discount_lines"`
	PreAuthorize    bool            `json:"pre_authorize"`
	CustomerInfo    Customer        `json:"customer_info"`
	ShippingContact ShippingContact `json:"shipping_contact"`
	Charges         []Charge        `json:"charges"`
}

// Product type.
type Product struct {
	Name      string `json:"name"`
	UnitPrice int64  `json:"unit_price"`
	Quantity  int    `json:"quantity"`
}

// Shipping type.
type Shipping struct{}

// Tax type.
type Tax struct{}

// Discount type.
type Discount struct{}

// Customer type.
type Customer struct {
	CustomerID string `json:"customer_id"`
}

// ShippingContact type.
type ShippingContact struct{}

// Charge type.
type Charge struct {
	PaymentMethod PaymentMethod `json:"payment_method"`
}

type PaymentMethod struct {
	Type string `json:"type"`
}

// Order type.
type Order struct{}
