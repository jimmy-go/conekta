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

func TestOrderCreate(t *testing.T) {

	c, err := NewClient("key", "secret", false)
	assert.Nil(t, err)
	assert.NotNil(t, c)

	or := &OrderRequest{
		Currency: MXN,
		CustomerInfo: Customer{
			CustomerID: "abc",
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

	or = &OrderRequest{
		Currency: MXN,
		CustomerInfo: Customer{
			CustomerID: "abc",
		},
		LineItems: []Product{
			Product{
				Name:      "Box of sls",
				UnitPrice: 150 * 100,
				Quantity:  1,
			},
		},
	}
	order, err := c.CreateOrder(or)
	assert.Nil(t, err)
	assert.NotNil(t, order)
}
