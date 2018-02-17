package conekta

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	// Host host URL.
	Host = "https://api.conekta.io"

	// OrderURI Order REST.
	OrderURI = "/orders"

	// CustomerURI Customer REST.
	CustomerURI = "/customers"
)

// Client implements all Conekta REST API endpoints.
type Client struct {
	client      *http.Client
	key, secret string
	sandbox     bool
	debug       bool
}

// NewClient returns a new Conekta client with http default settings. Use sandbox
// flag as false for production environments.
func NewClient(key, secret string, sandbox bool) (*Client, error) {
	c := &Client{
		key:     key,
		secret:  secret,
		sandbox: sandbox,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
	return c, nil
}

func (c *Client) bind(dst, src interface{}, method, uri string) error {
	b, err := json.MarshalIndent(src, "", "	")
	if err != nil {
		return err
	}
	if c.debug {
		log.Printf("bind : method [%s] uri [%s] request body [%s]", method, uri, string(b))
	}
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.conekta-v2.0.0+json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.secret, "")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("bind : close body : err [%s]", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			log.Printf("bind : read error response : err [%s]", err)
		}
		if c.debug {
			log.Printf("bind : read error response [%s]", buf.String())
		}

		var f Fail
		if err := json.NewDecoder(buf).Decode(&f); err != nil {
			return err
		}
		return f
	}
	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return err
	}
	return nil
}

// CreateOrder _
func (c *Client) CreateOrder(or *OrderRequest) (*Order, error) {
	var order *Order
	err := c.bind(&order, or, "POST", Host+OrderURI)
	return order, err
}

// CreateCustomer _
func (c *Client) CreateCustomer(cr *CustomerRequest) (*Customer, error) {
	var cus *Customer
	err := c.bind(&cus, cr, "POST", Host+CustomerURI)
	return cus, err
}

// Fail  _
type Fail struct {
	Details []Detail `json:"details"`
	Object  string   `json:"object"`
	Type    string   `json:"type"`
	LogID   string   `json:"log_id"`
}

// Error implements error.
func (f Fail) Error() string {
	return f.Type
}

// Detail _
type Detail struct {
	DebugMessage string `json:"debug_message"`
	Message      string `json:"message"`
	Code         string `json:"code"`
	Param        string `json:"param"`
}

const (
	// MXN mexican pesos code ISO 4217.
	MXN = "MXN"
)

// OrderRequest type.
// see: https://developers.conekta.com/api?language=bash#order
type OrderRequest struct {
	Currency        string           `json:"currency,omitempty"`
	LineItems       []Product        `json:"line_items,omitempty"`
	ShippingLines   []Shipping       `json:"shipping_lines,omitempty"`
	TaxLines        []Tax            `json:"tax_lines,omitempty"`
	DiscountLines   []Discount       `json:"discount_lines,omitempty"`
	PreAuthorize    bool             `json:"pre_authorize,omitempty"`
	CustomerInfo    *Customer        `json:"customer_info,omitempty"`
	ShippingContact *ShippingContact `json:"shipping_contact,omitempty"`
	Charges         []Charge         `json:"charges,omitempty"`
}

// Product type.
type Product struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	UnitPrice   int64    `json:"unit_price,omitempty"`
	Quantity    int      `json:"quantity,omitempty"`
	SKU         string   `json:"sku,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Brand       string   `json:"brand,omitempty"`
}

// Shipping type.
type Shipping struct{}

// Tax type.
type Tax struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
}

// Discount type.
type Discount struct {
	ID     string `json:"id,omitempty"`
	Code   string `json:"code,omitempty"`
	Type   string `json:"type,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}

// CustomerRequest _
type CustomerRequest struct {
	Name             string            `json:"name,omitempty"`
	Phone            string            `json:"phone,omitempty"`
	Email            string            `json:"email,omitempty"`
	PlanID           string            `json:"plan_id,omitempty"`
	PaymentSources   []PaymentSource   `json:"payment_sources,omitempty"`
	Corporate        bool              `json:"corporate,omitempty"`
	ShippingContacts []ShippingContact `json:"shipping_contacts,omitempty"`
}

// PaymentSource _
type PaymentSource struct {
	TokenID string `json:"token_id,omitempty"`
	Type    string `json:"type,omitempty"`
}

// Customer type.
type Customer struct {
	CustomerID string `json:"customer_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email,omitempty"`
	Corporate  bool   `json:"corporate,omitempty"`
}

// ShippingContact type.
type ShippingContact struct {
	Phone          string  `json:"phone,omitempty"`
	Receiver       string  `json:"receiver,omitempty"`
	BetweenStreets string  `json:"between_streets,omitempty"`
	Address        Address `json:"address,omitempty"`
}

// Address _
type Address struct {
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	Residential bool   `json:"residential,omitempty"`
}

// Charge type.
type Charge struct {
	ID            string        `json:"id,omitempty"`
	CreatedAt     int64         `json:"created_at,omitempty"`
	Currency      string        `json:"currency,omitempty"`
	Amount        int64         `json:"amount,omitempty"`
	Livemode      bool          `json:"livemode,omitempty"`
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
	Fee           int64         `json:"fee,omitempty"`
}

// PaymentMethod _
type PaymentMethod struct {
	TokenID string `json:"token_id,omitempty"`
	Type    string `json:"type,omitempty"`
}

// Order type.
type Order struct {
	ID       string `json:"id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency,omitempty"`
	// TODO; add fields.
	// see: https://pastebin.com/nj8JnYYv
}
