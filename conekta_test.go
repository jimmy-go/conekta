// Package conekta contains tests for Conekta client.
/*
	Use of this source code is governed by a MIT
	license that can be found in the LICENSE file.
*/
package conekta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	key    = ""
	secret = "key_KjXFWbopqyujnagRLy4dtw"
)

func TestOrderCreate(t *testing.T) {

	c, err := NewClient(key, secret, false)
	c.debug = true
	assert.Nil(t, err)
	assert.NotNil(t, c)

	or := &OrderRequest{
		Currency: MXN,
		CustomerInfo: &Customer{
			CustomerID: "cus_2h3syNSMiZwfFM5XC",
		},
		LineItems: []Product{
			Product{
				Name:      "Box of sls",
				UnitPrice: 150 * 100,
				Quantity:  1,
			},
		},
		Charges: []Charge{
			Charge{
				PaymentMethod: PaymentMethod{
					Type: "default",
				},
			},
		},
	}

	order, err := c.CreateOrder(or)
	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestCustomerCreate(t *testing.T) {

	c, err := NewClient(key, secret, false)
	c.debug = true
	assert.Nil(t, err)
	assert.NotNil(t, c)

	cr := &CustomerRequest{
		Name:  "Mario Perez",
		Phone: "+5215555555555",
		Email: "usuario@example.com",
		PaymentSources: []PaymentSource{
			PaymentSource{
				TokenID: "tok_test_visa_4242",
				Type:    "card",
			},
		},
		Corporate: true,
	}
	cus, err := c.CreateCustomer(cr)
	assert.Nil(t, err)
	assert.NotNil(t, cus)
}
